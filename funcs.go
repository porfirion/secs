package secs

// AddComponent allocated new component with specified value and returns pointer to this component
func AddComponent[CT Component](ctx Context, e *Entity, cv CT) *CT {
	n := GetStore[CT](ctx).New(e.ID)
	*n = cv
	e.Components = append(e.Components, n)
	return n
}

func RemoveComponent[CT Component](ctx Context, e *Entity) {
	GetStore[CT](ctx).Remove(e.ID)
	for i := range e.Components {
		if _, ok := e.Components[i].(*CT); ok {
			e.Components = append(e.Components[:i], e.Components[i+1:]...)
		}
	}
}

func GetStore[CT Component](ctx Context) *Store[EntityID, CT] {
	var s *Store[EntityID, CT]
	var stores = ctx.GetStores()
	for _, st := range *stores {
		if st, ok := (st).(*Store[EntityID, CT]); ok {
			s = st
			break
		}
	}

	if s == nil {
		s = NewStore[EntityID, CT]()
		*stores = append(*stores, s)
	}

	return s
}

func Iter[CT1 Component](ctx Context, f func(id EntityID, ct1 *CT1)) int {
	all1 := GetStore[CT1](ctx).Index()
	count := 0
	for id, c1 := range all1 {
		f(id, c1)
		count++
	}
	return count
}
func Iter2[CT1 Component, CT2 Component](ctx Context, f func(id EntityID, ct1 *CT1, ct2 *CT2)) int {
	all1 := GetStore[CT1](ctx).Index()
	all2 := GetStore[CT2](ctx).Index()
	count := 0
	for id, c1 := range all1 {
		if c2, ok := all2[id]; ok {
			f(id, c1, c2)
			count++
		}
	}
	return count
}
func Iter3[CT1 Component, CT2 Component, CT3 Component](ctx Context, f func(id EntityID, ct1 *CT1, ct2 *CT2, ct3 *CT3)) int {
	all1 := GetStore[CT1](ctx).Index()
	all2 := GetStore[CT2](ctx).Index()
	all3 := GetStore[CT3](ctx).Index()
	count := 0
	for id, c1 := range all1 {
		if c2, ok := all2[id]; ok {
			if c3, ok := all3[id]; ok {
				f(id, c1, c2, c3)
				count++
			}
		}
	}
	return count
}
