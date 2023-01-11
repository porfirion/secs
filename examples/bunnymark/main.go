package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/porfirion/secs/examples/bunnymark/scene"
)

func main() {
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowSizeLimits(300, 200, -1, -1)
	ebiten.SetFPSMode(ebiten.FPSModeVsyncOffMaximum)
	ebiten.SetWindowResizable(true)
	rand.Seed(time.Now().UTC().UnixNano())
	//if err := ebiten.RunGame(engine.NewGame(&scene.Game{})); err != nil {
	//	log.Fatal(err)
	//}

	if err := ebiten.RunGame(scene.SetupMGame()); err != nil {
		log.Fatal(err)
	}
}
