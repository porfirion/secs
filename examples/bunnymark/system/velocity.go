package system

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/porfirion/secs"
	"github.com/porfirion/secs/examples/bunnymark/component"
	"github.com/porfirion/secs/examples/bunnymark/mgame"
	"github.com/sedyh/mizu/pkg/engine"
)

type Velocity struct {
	*component.Position
	*component.Velocity
}

func (v *Velocity) Update(_ engine.World) {
	v.Position.X += v.Velocity.X
	v.Position.Y += v.Velocity.Y
}

func (v *Velocity) MDraw(game *mgame.MGame, image *ebiten.Image) {}

func (v *Velocity) MUpdate(game *mgame.MGame) {
	secs.Iter2[component.Position, component.Velocity](game, func(id secs.EntityID, Position *component.Position, Velocity *component.Velocity) {
		Position.X += Velocity.X
		Position.Y += Velocity.Y
	})
}
