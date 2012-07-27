package below

import (
	"code.google.com/p/termon"
)

const (
	BACKSPACE = 8
	LF        = 10
	CR        = 13
	ESCAPE    = 27
	DELETE    = 127
)

func Min(a ...int) int {
	min := int(^uint(0) >> 1) // largest int
	for _, i := range a {
		if i < min {
			min = i
		}
	}
	return min
}

func main() {
	// Must!
	term.Init()

	// Allows use of function keys and arrow keys.
	term.Keypad()

	// Suppress user input.
	term.Noecho()

	game := NewGame()
	game.Run()

	// Reset the terminal.
	// It will look as well as before term.Init() was called.
	term.End()
}
