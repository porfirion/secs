secs (simple ecs) is a minimalistic implementation of ECS pattern in go with type arguments support.

Idea is very simple:
- components are identified by their type.
- all components are stored inside Context in chunks (Context supposed to embed anywhere).
- when you add component - you pass initial value for this component (it is copied into context)
- entity stores pointers to these values.
- you can iterate over entities by specifying required types of components.

Too much words, just look at the code:
```go
type Tag string                             // component can be of any type 
type Coords struct {X, Y int}               // 

ctx := secs.NewContext(ecs.IDgen(0))        // (aka world/scene/whatever)
entity := ctx.NewEntity()                   // entity already has ID

AddComponent(ctx, e, Tag("my tag"))         // it is Tag component, not string
AddComponent(ctx, e, Coords{X: 10, Y: 15})
AddComponent(ctx, e, int64(0))              // int64 component %)
AddComponent(ctx, e, new(bool))             // *bool component (you'll get **bool while iteration)

secs.Iter2[Tag, int64](ctx, func(id secs.EntityID, ct1 *Tag, ct2 *int64) {
    fmt.Printf("Here is entity #%d with tag=\"%s\" and int64=%d", id, *ct1, *ct2)
})
```

There is no type called "System" in this package, but it should be very easy to make it, 
since there are functions to iterate over components by their types (see example).
Yeah, max for 3 types, but you can implement for any amount - they very simple.

Also, "Component" is just an empty interface to avoid any limitations - it can be 
structure, primitive type or even pointer. 

## Example
See examples/bunnymark - it's [mizu's Bunnymark](https://github.com/sedyh/mizu/tree/master/examples/bunnymark) ported to secs

Also examples/bunnymark/mgame - contains example implementation for ebiten.Game
