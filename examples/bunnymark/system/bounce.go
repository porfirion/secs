package system

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/porfirion/secs"
	"github.com/porfirion/secs/examples/bunnymark/component"
	"github.com/porfirion/secs/examples/bunnymark/helper"
	"github.com/porfirion/secs/examples/bunnymark/mgame"
	"github.com/sedyh/mizu/pkg/engine"
)

type Bounce struct {
	*component.Position
	*component.Velocity
	*component.Sprite
}

func (b *Bounce) Update(w engine.World) {
	sw, sh := float64(w.Bounds().Dx()), float64(w.Bounds().Dy())
	iw, ih := float64(b.Sprite.Image.Bounds().Dx()), float64(b.Sprite.Image.Bounds().Dy())
	relW, relH := iw/sw, ih/sh
	if b.Position.X+relW > 1 {
		b.Velocity.X *= -1
		b.Position.X = 1 - relW
	}
	if b.Position.X < 0 {
		b.Velocity.X *= -1
		b.Position.X = 0
	}
	if b.Position.Y+relH > 1 {
		b.Velocity.Y *= -0.85
		b.Position.Y = 1 - relH
		if helper.Chance(0.5) {
			b.Velocity.Y -= helper.RangeFloat(0, 0.009)
		}
	}
	if b.Position.Y < 0 {
		b.Velocity.Y = 0
		b.Position.Y = 0
	}
}

func (b *Bounce) MDraw(game *mgame.MGame, image *ebiten.Image) {}

func (b *Bounce) MUpdate(game *mgame.MGame) {
	secs.Iter3[component.Position, component.Velocity, component.Sprite](game, func(id secs.EntityID, Position *component.Position, Velocity *component.Velocity, Sprite *component.Sprite) {
		sw, sh := float64(game.Bounds().Dx()), float64(game.Bounds().Dy())
		iw, ih := float64(Sprite.Image.Bounds().Dx()), float64(Sprite.Image.Bounds().Dy())
		relW, relH := iw/sw, ih/sh
		if Position.X+relW > 1 {
			Velocity.X *= -1
			Position.X = 1 - relW
		}
		if Position.X < 0 {
			Velocity.X *= -1
			Position.X = 0
		}
		if Position.Y+relH > 1 {
			Velocity.Y *= -0.85
			Position.Y = 1 - relH
			if helper.Chance(0.5) {
				Velocity.Y -= helper.RangeFloat(0, 0.009)
			}
		}
		if Position.Y < 0 {
			Velocity.Y = 0
			Position.Y = 0
		}
	})
}
