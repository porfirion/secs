package scene

import (
	"fmt"
	"math"
	"testing"

	"github.com/porfirion/secs"
	"github.com/porfirion/secs/examples/bunnymark/assets"
	"github.com/porfirion/secs/examples/bunnymark/component"
	"github.com/porfirion/secs/examples/bunnymark/entity"
	"github.com/porfirion/secs/examples/bunnymark/helper"
	"github.com/porfirion/secs/examples/bunnymark/mgame"
	"github.com/sedyh/mizu/pkg/engine"
)

var amounts = []int{0, 100, 1_000, 10_000, 25_000, 50_000, 100_000, 200_000, 500_000}

func BenchmarkAdd(b *testing.B) {
	b.ReportAllocs()
	g := engine.NewGame(&Game{})
	_ = g.Update()
	addBunnies(g.(engine.World), b.N)
}

func BenchmarkMAdd(b *testing.B) {
	b.ReportAllocs()
	g := SetupMGame()
	addMBunnies(g, b.N)
}

func BenchmarkGame(b *testing.B) {
	for i := range amounts {
		amount := amounts[i]
		b.Run(fmt.Sprintf("%d", amount), func(b *testing.B) {
			runGame(b, amount)
		})
	}
}

func BenchmarkMGame(b *testing.B) {
	for i := range amounts {
		amount := amounts[i]
		b.Run(fmt.Sprintf("%d", amount), func(b *testing.B) {
			runMGame(b, amount)
		})
	}
}

func runGame(b *testing.B, amount int) {
	g := engine.NewGame(&Game{})
	_ = g.Update()
	addBunnies(g.(engine.World), amount)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := g.Update(); err != nil {
			b.Fatal(err)
		}
	}
}
func runMGame(b *testing.B, amount int) {
	g := SetupMGame()

	addMBunnies(g, amount)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := g.Update(); err != nil {
			b.Fatal(err)
		}
	}
}

func addBunnies(w engine.World, amount int) {
	for i := 0; i < amount; i++ {
		w.AddEntities(&entity.Bunny{
			Position: component.Position{
				X: float64(w.Entities() % 2), // Alternate screen edges
			},
			Velocity: component.Velocity{
				X: helper.RangeFloat(0, 0.005),
				Y: helper.RangeFloat(0.0025, 0.005)},
			Hue: component.Hue{
				Colorful: new(bool),
				Value:    helper.RangeFloat(0, 2*math.Pi),
			},
			Gravity: component.Gravity{Value: 0.00095},
			Sprite:  component.Sprite{Image: assets.Bunny},
		})
	}
}
func addMBunnies(w *mgame.MGame, amount int) {
	for i := 0; i < amount; i++ {
		e := w.NewEntity()
		secs.AddComponent(w, e, component.Position{
			X: float64(len(w.AllEntities()) % 2), // Alternate screen edges
		})
		secs.AddComponent(w, e, component.Velocity{
			X: helper.RangeFloat(0, 0.005),
			Y: helper.RangeFloat(0.0025, 0.005),
		})
		secs.AddComponent(w, e, component.Hue{
			Colorful: new(bool),
			Value:    helper.RangeFloat(0, 2*math.Pi),
		})
		secs.AddComponent(w, e, component.Gravity{Value: 0.00095})
		secs.AddComponent(w, e, component.Sprite{Image: assets.Bunny})
	}
}
