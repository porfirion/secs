package system

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/porfirion/secs"
	"github.com/porfirion/secs/examples/bunnymark/component"
	"github.com/porfirion/secs/examples/bunnymark/mgame"
	"github.com/sedyh/mizu/pkg/engine"
)

type Render struct {
	*component.Position
	*component.Sprite
	*component.Hue
}

func (r *Render) Draw(_ engine.World, screen *ebiten.Image) {
	sw, sh := float64(screen.Bounds().Dx()), float64(screen.Bounds().Dy())
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(r.Position.X*sw, r.Position.Y*sh)
	if *r.Hue.Colorful {
		op.ColorM.RotateHue(r.Hue.Value)
	}
	screen.DrawImage(r.Sprite.Image, op)
}

func (r *Render) MDraw(game *mgame.MGame, screen *ebiten.Image) {
	secs.Iter3[component.Position, component.Sprite, component.Hue](game, func(id secs.EntityID, Position *component.Position, Sprite *component.Sprite, Hue *component.Hue) {
		sw, sh := float64(screen.Bounds().Dx()), float64(screen.Bounds().Dy())
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(Position.X*sw, Position.Y*sh)
		if *Hue.Colorful {
			op.ColorM.RotateHue(Hue.Value)
		}
		screen.DrawImage(Sprite.Image, op)
	})
}

func (r *Render) MUpdate(game *mgame.MGame) {}
