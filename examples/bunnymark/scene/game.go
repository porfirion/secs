package scene

import (
	"time"

	ecs "github.com/porfirion/secs"
	"github.com/porfirion/secs/examples/bunnymark/assets"
	"github.com/porfirion/secs/examples/bunnymark/component"
	"github.com/porfirion/secs/examples/bunnymark/entity"
	"github.com/porfirion/secs/examples/bunnymark/helper"
	"github.com/porfirion/secs/examples/bunnymark/mgame"
	"github.com/porfirion/secs/examples/bunnymark/system"
	"github.com/sedyh/mizu/pkg/engine"
)

type Game struct{}

func (g *Game) Setup(w engine.World) {
	w.AddComponents(
		component.Position{}, component.Velocity{}, component.Gravity{},
		component.Sprite{}, component.Hue{}, component.Settings{},
	)
	w.AddSystems(
		&system.Background{}, &system.Velocity{}, &system.Gravity{},
		&system.Bounce{}, &system.Render{},
		&system.Metrics{},
		&system.Spawn{},
	)
	w.AddEntities(
		&entity.Settings{
			Settings: component.Settings{
				Ticker:   time.NewTicker(500 * time.Millisecond),
				Gpu:      helper.GpuInfo(),
				Tps:      helper.NewPlot(20, 60),
				Fps:      helper.NewPlot(20, 60),
				Objects:  helper.NewPlot(20, 60000),
				Sprite:   assets.Bunny,
				Colorful: false,
				Amount:   100,
			},
		},
	)
}

func SetupMGame() *mgame.MGame {
	var id ecs.EntityID = 0

	g := &mgame.MGame{
		Context: ecs.NewContext(func() ecs.EntityID {
			id++
			return id
		}),
	}
	g.AddSystems(
		&system.Background{},
		&system.Velocity{},
		&system.Gravity{},
		&system.Bounce{},
		&system.Render{},
		&system.Metrics{},
		&system.Spawn{},
	)

	e := g.NewEntity()
	_ = ecs.AddComponent(g, e, component.Settings{
		Ticker:   time.NewTicker(500 * time.Millisecond),
		Gpu:      helper.GpuInfo(),
		Tps:      helper.NewPlot(20, 60),
		Fps:      helper.NewPlot(20, 60),
		Objects:  helper.NewPlot(20, 60000),
		Sprite:   assets.Bunny,
		Colorful: false,
		Amount:   100,
	})

	return g
}
