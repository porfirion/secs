package system

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/porfirion/secs"
	"github.com/porfirion/secs/examples/bunnymark/component"
	"github.com/porfirion/secs/examples/bunnymark/entity"
	"github.com/porfirion/secs/examples/bunnymark/helper"
	"github.com/porfirion/secs/examples/bunnymark/mgame"
	"github.com/sedyh/mizu/pkg/engine"
)

type Spawn struct {
	*component.Settings
}

func (s *Spawn) Update(w engine.World) {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		s.addBunnies(w)
	}

	if ids := ebiten.AppendTouchIDs(nil); len(ids) > 0 {
		s.addBunnies(w) // not accurate, cause no input manager for this
	}

	if _, offset := ebiten.Wheel(); offset != 0 {
		s.Amount += int(offset * 10)
		if s.Amount < 0 {
			s.Amount = 0
		}
	}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) {
		s.Colorful = !s.Colorful
	}
}

func (s *Spawn) addBunnies(w engine.World) {
	// Spawns specific amount of bunnies at the edges of the screen
	// It will alternately add bunnies to the left and right corners of the screen
	for i := 0; i < s.Amount; i++ {
		w.AddEntities(&entity.Bunny{
			Position: component.Position{
				X: float64(w.Entities() % 2), // Alternate screen edges
			},
			Velocity: component.Velocity{
				X: helper.RangeFloat(0, 0.005),
				Y: helper.RangeFloat(0.0025, 0.005)},
			Hue: component.Hue{
				Colorful: &s.Colorful,
				Value:    helper.RangeFloat(0, 2*math.Pi),
			},
			Gravity: component.Gravity{Value: 0.00095},
			Sprite:  component.Sprite{Image: s.Sprite},
		})
	}
}

func (s *Spawn) MDraw(game *mgame.MGame, image *ebiten.Image) {}

func (s *Spawn) MUpdate(game *mgame.MGame) {
	secs.Iter[component.Settings](game, func(id secs.EntityID, settings *component.Settings) {
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			s.maddBunnies(game, settings)
		}

		if ids := ebiten.AppendTouchIDs(nil); len(ids) > 0 {
			s.maddBunnies(game, settings) // not accurate, cause no input manager for this
		}

		if _, offset := ebiten.Wheel(); offset != 0 {
			settings.Amount += int(offset * 10)
			if settings.Amount < 0 {
				settings.Amount = 0
			}
		}

		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) {
			settings.Colorful = !settings.Colorful
		}
	})
}

func (s *Spawn) maddBunnies(w *mgame.MGame, settings *component.Settings) {
	// Spawns specific amount of bunnies at the edges of the screen
	// It will alternately add bunnies to the left and right corners of the screen
	for i := 0; i < settings.Amount; i++ {
		e := w.NewEntity()
		secs.AddComponent(w, e, component.Position{
			X: float64(len(w.AllEntities()) % 2), // Alternate screen edges
		})
		secs.AddComponent(w, e, component.Velocity{
			X: helper.RangeFloat(0, 0.005),
			Y: helper.RangeFloat(0.0025, 0.005),
		})
		secs.AddComponent(w, e, component.Hue{
			Colorful: &settings.Colorful,
			Value:    helper.RangeFloat(0, 2*math.Pi),
		})
		secs.AddComponent(w, e, component.Gravity{Value: 0.00095})
		secs.AddComponent(w, e, component.Sprite{Image: settings.Sprite})
	}
}
