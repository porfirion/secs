package secs

type context struct {
	*Stores
	idGenerator func() EntityID

	// Special store for storing entities itself %)
	Entities *Store[EntityID, Entity]
}

func NewContext(idGenerator func() EntityID) Context {
	return &context{
		Stores:      new(Stores),
		idGenerator: idGenerator,
		Entities:    NewStore[EntityID, Entity](),
	}
}

func (c *context) NewEntity() *Entity {
	eid := c.idGenerator()
	e := c.Entities.New(eid)
	*e = Entity{
		ID:  eid,
		Ctx: c,
	}
	return e
}

func (c *context) GetEntity(id EntityID) *Entity {
	return c.Entities.Index()[id]
}

func (c *context) RemoveEntity(id EntityID) {
	c.Entities.Remove(id)
}

func (c *context) GetStores() *Stores {
	return c.Stores
}

func (c *context) AllEntities() map[EntityID]*Entity {
	return c.Entities.Index()
}
