package secs

import (
	"fmt"
	"log"
)

type (
	storeChunk[KT comparable, VT any] struct {
		keys      []KT
		values    []VT
		capacity  int
		allocated int
	}

	Store[KT comparable, VT any] struct {
		index  map[KT]*VT
		chunks []*storeChunk[KT, VT]

		nextChunkNum int
		nextChunkInd int

		zeroKey   KT
		zeroValue VT
	}

	Stores = []any
)

func NewStore[EntityID comparable, CT any]() *Store[EntityID, CT] {
	return &Store[EntityID, CT]{
		index: make(map[EntityID]*CT, 1000),
	}
}

// alloc allocates space for value but doesn't register it in index
func (s *Store[KT, VT]) alloc(key KT) (chunkNum int, chunkInd int, ptr *VT) {
	defer func() {
		s.chunks[chunkNum].allocated++
		s.nextChunkNum = chunkNum
		s.nextChunkInd = chunkInd + 1
	}()

	if s.nextChunkNum < len(s.chunks) &&
		s.nextChunkInd < len(s.chunks[s.nextChunkNum].keys) &&
		s.chunks[s.nextChunkNum].keys[s.nextChunkInd] == s.zeroKey {
		chunkNum = s.nextChunkNum
		chunkInd = s.nextChunkInd

		s.chunks[chunkNum].keys[chunkInd] = key
		return chunkNum, chunkInd, &s.chunks[chunkNum].values[chunkInd]
	}

	for chunkNum, chunk := range s.chunks {
		if chunk.capacity == chunk.allocated {
			continue
		}
		for chunkInd := range chunk.keys {
			if chunk.keys[chunkInd] == s.zeroKey {
				chunk.keys[chunkInd] = key
				return chunkNum, chunkInd, &chunk.values[chunkInd]
			}
		}
		// we checked all elements of this chunk, but no free found
		if chunk.allocated < chunk.capacity {
			panic("allocated < capacity, but no free elements")
		}
	}

	// we checked all chunks - they all are full. Let's alloc new one.
	newChunk := &storeChunk[KT, VT]{
		keys:     make([]KT, 1000),
		values:   make([]VT, 1000),
		capacity: 1000,
	}
	s.chunks = append(s.chunks, newChunk)

	chunkNum = len(s.chunks) - 1
	chunkInd = 0

	newChunk.keys[0] = key
	return chunkNum, chunkInd, &newChunk.values[0]
}

func (s *Store[KT, VT]) free(chunkNum, chunkInd int) {
	chunk := s.chunks[chunkNum]
	chunk.keys[chunkInd] = s.zeroKey
	chunk.values[chunkInd] = s.zeroValue

	chunk.allocated--
	s.nextChunkNum = chunkNum
	s.nextChunkInd = chunkInd
}

func (s *Store[KT, VT]) New(key KT) *VT {
	if _, ok := s.index[key]; ok {
		panic(fmt.Sprintf("store %T already contains value for key %v", s, key))
	}

	_, _, ptr := s.alloc(key)
	s.index[key] = ptr
	return ptr
}

func (s *Store[KT, VT]) Remove(key KT) {
	for chunkNum, chunk := range s.chunks {
		for chunkInd := range chunk.keys {
			if chunk.keys[chunkInd] == key {
				s.free(chunkNum, chunkInd)
				delete(s.index, key)
				return
			}
		}
	}
	log.Printf("key %v not found in store %T\n", key, s)
}

func (s *Store[KT, VT]) Index() map[KT]*VT {
	return s.index
}
