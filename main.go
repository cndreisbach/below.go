package main

import (
	"./lib/below"
	"code.google.com/p/termon"
)

func main() {
	// Must!
	term.Init()

	// Allows use of function keys and arrow keys.
	term.Keypad()

	// Suppress user input.
	term.Noecho()

	game := below.NewGame()
	game.Run()

	// Reset the terminal.
	// It will look as well as before term.Init() was called.
	term.End()
}
