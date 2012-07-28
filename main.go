package main

import (
	"./lib/below"
	"code.google.com/p/goncurses"
	"math/rand"
	"os"
	"time"
)

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().Unix())

	stdscr, err := goncurses.Init()
	defer goncurses.End()

	goncurses.StartColor()

	goncurses.Raw(true)
	goncurses.Echo(false)
	goncurses.CBreak(true)
	goncurses.Cursor(0)

	below.SetupColors()

	if err != nil {
		os.Exit(1)
	}

	game := below.NewGame(stdscr)
	game.Run()
}
