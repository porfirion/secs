package mgame

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/porfirion/secs"
)

type MSystem interface {
	MDraw(*MGame, *ebiten.Image)
	MUpdate(*MGame)
}

type MGame struct {
	secs.Context
	bounds  image.Rectangle
	systems []MSystem
}

func (w *MGame) Update() error {
	for i := range w.systems {
		w.systems[i].MUpdate(w)
	}

	return nil
}

func (w *MGame) Draw(screen *ebiten.Image) {
	for i := range w.systems {
		w.systems[i].MDraw(w, screen)
	}
}

func (w *MGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	w.bounds = image.Rect(0, 0, outsideWidth, outsideHeight)
	return outsideWidth, outsideHeight
}

func (w *MGame) AddSystems(systems ...MSystem) {
	w.systems = append(w.systems, systems...)
}

func (w *MGame) Bounds() image.Rectangle {
	return w.bounds
}
