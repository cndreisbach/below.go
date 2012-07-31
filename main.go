package main

import (
	"./lib/below"
	"./lib/below/ui"
	"math/rand"
	"time"
)

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().Unix())

	ui.Init()
	defer ui.End()

	game := below.NewGame()
	game.Run()
}
