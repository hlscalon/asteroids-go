package main

import (
	"asteroids-go/asteroids"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game, err := asteroids.NewGame()
	if err != nil {
		log.Fatal(err)
	}

	ebiten.SetWindowSize(asteroids.ScreenWidth, asteroids.ScreenHeight)
	ebiten.SetWindowTitle("Asteroids")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
