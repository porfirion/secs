package system

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/porfirion/secs"
	"github.com/porfirion/secs/examples/bunnymark/component"
	"github.com/porfirion/secs/examples/bunnymark/mgame"
	"github.com/sedyh/mizu/pkg/engine"
)

type Gravity struct {
	*component.Velocity
	*component.Gravity
}

func (g *Gravity) Update(_ engine.World) {
	g.Velocity.Y += g.Gravity.Value
}

func (g *Gravity) MDraw(game *mgame.MGame, image *ebiten.Image) {}

func (g *Gravity) MUpdate(game *mgame.MGame) {
	secs.Iter2[component.Velocity, component.Gravity](game, func(id secs.EntityID, Velocity *component.Velocity, Gravity *component.Gravity) {
		Velocity.Y += Gravity.Value
	})
}
