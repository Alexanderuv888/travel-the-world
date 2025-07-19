package main

import (
	"log"

	"travel-the-world/game"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowTitle("Isometric (Ebitengine Demo)")
	ebiten.SetWindowSize(1728, 1117)
	ebiten.SetWindowResizable(true)

	g, err := game.NewGame()
	if err != nil {
		log.Fatal(err)
	}

	if err = ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
