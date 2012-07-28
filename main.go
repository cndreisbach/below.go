package main

import (
	"./lib/below"
	"code.google.com/p/termon"
	"math/rand"
	"time"
)

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().Unix())

	// Must!
	term.Init()

	// Allows use of function keys and arrow keys.
	term.Keypad()

	// Suppress user input.
	term.Noecho()

	below.SetupColors()

	game := below.NewGame()
	game.Run()

	// Reset the terminal.
	// It will look as well as before term.Init() was called.
	term.End()
}
