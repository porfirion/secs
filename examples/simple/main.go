package main

import (
	"fmt"

	"github.com/porfirion/secs"
)

type Tag string

func main() {
	ctx := secs.NewContext(secs.IDgen(0))
	e := ctx.NewEntity()
	_ = secs.AddComponent(ctx, e, Tag("my tag"))
	_ = secs.AddComponent(ctx, e, int64(123))
	_ = secs.AddComponent(ctx, e, new(bool))

	secs.Iter3[Tag, int64, *bool](ctx, func(id secs.EntityID, ct1 *Tag, ct2 *int64, ct3 **bool) {
		fmt.Printf("Here is entity #%d with tag=\"%s\" and int64=%d and *bool=&%t", id, *ct1, *ct2, **ct3)
	})
}
