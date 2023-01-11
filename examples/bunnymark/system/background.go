package system

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/porfirion/secs/examples/bunnymark/assets"
	"github.com/porfirion/secs/examples/bunnymark/mgame"
	"github.com/sedyh/mizu/pkg/engine"
)

type Background struct{}

func (b *Background) Draw(_ engine.World, screen *ebiten.Image) {
	screen.Fill(assets.Background)
}

func (b *Background) MDraw(game *mgame.MGame, screen *ebiten.Image) {
	screen.Fill(assets.Background)
}

func (b *Background) MUpdate(game *mgame.MGame) {}
