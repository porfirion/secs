package secs

type (
	Component interface{}

	EntityID uint64

	Entity struct {
		ID         EntityID
		Ctx        Context
		Components []any
	}

	Context interface {
		GetStores() *Stores
		NewEntity() *Entity
		GetEntity(EntityID) *Entity
		RemoveEntity(EntityID)
		AllEntities() map[EntityID]*Entity
	}
)

var (
	IDgen = func(start EntityID) func() EntityID {
		return func() EntityID {
			start++
			return start
		}
	}
)
