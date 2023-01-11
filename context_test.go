package secs

import (
	"fmt"
	"log"
	"testing"
)

type Tag string
type Tags []string
type Number int

func Test(t *testing.T) {
	stores := make(Stores, 0)
	ctx := NewContext(IDgen(0))

	e := ctx.NewEntity()
	tag := AddComponent(ctx, e, Tag("123"))
	tags := AddComponent(ctx, e, Tags{"456"})

	log.Println("Stores:")
	for i := range stores {
		log.Printf("\t%+v", stores[i])
	}

	log.Printf("%#v", tag)
	log.Printf("%#v", tags)

	Iter[*Tag](ctx, func(id EntityID, ct1 **Tag) {
		log.Println("found *Tag")
	})

	Iter[Tag](ctx, func(id EntityID, ct1 *Tag) {
		log.Println("found Tag")
	})

	log.Println("Iterating over Tag and Tags:")
	Iter2[Tag, Tags](ctx, func(id EntityID, ct1 *Tag, ct2 *Tags) {
		log.Printf("\t%v: %+v %+v (%[3]p)\n", id, *ct1, *ct2)
	})

	Iter2[Tag, Tags](ctx, func(id EntityID, ct1 *Tag, ct2 *Tags) {
		*ct1 = "789"
		*ct2 = append(*ct2, "1", "2", "3", "4", "5")
	})

	log.Println("Iterating over Tag and Tags:")
	Iter2[Tag, Tags](ctx, func(id EntityID, ct1 *Tag, ct2 *Tags) {
		log.Printf("\t%v: %+v %+v (%[3]p)\n", id, *ct1, *ct2)
	})
}

func BenchmarkAddComponent(b *testing.B) {
	b.ReportAllocs()
	ctx := NewContext(IDgen(0))

	e := ctx.NewEntity()
	for i := 0; i < b.N; i++ {
		AddComponent(ctx, e, Tag(""))
		RemoveComponent[Tag](ctx, e)
	}
}

func BenchmarkAddAndIter2(b *testing.B) {
	b.ReportAllocs()
	ctx := NewContext(IDgen(0))

	count := 0
	for i := 0; i < b.N; i++ {
		e := ctx.NewEntity()
		_ = e
		AddComponent(ctx, e, Tag(""))
		if i%2 == 0 {
			count++
			AddComponent(ctx, e, Number(i))
		}
	}

	n := Iter2[Number, Tag](ctx, func(id EntityID, ct2 *Number, ct1 *Tag) {})

	if n != count {
		b.Errorf("wrong number of entries: expected %d, got %d", count, n)
	}
}

func ExampleNewContext() {
	ctx := NewContext(IDgen(0))
	e := ctx.NewEntity()
	_ = AddComponent(ctx, e, Tag("my tag"))
	_ = AddComponent(ctx, e, int64(123))

	Iter2[Tag, int64](ctx, func(id EntityID, ct1 *Tag, ct2 *int64) {
		fmt.Printf("Here is entity #%d with tag=\"%s\" and int64=%d", id, *ct1, *ct2)
	})
	// Output: Here is entity #1 with tag="my tag" and int64=123
}
