package main

import (
	"code.google.com/p/termon"
)

type UI string
type World [][]int
type Game struct {
	uis   []UI
	world World
}
type Tile struct {
	kind  string
	glyph string
	color string
}

var (
	WORLD_SIZE = []int{160, 60}

	TILES = map[string]Tile{
		"floor": Tile{"floor", ".", "white"},
		"wall":  Tile{"wall", "#", "white"},
		"bound": Tile{"bound", "X", "black"},
	}
)

func (ui UI) Draw() {
	switch ui {
	case "start":
		term.AddAt(0, 0, "Welcome to Below!")
		term.AddAt(0, 1, "Press Enter to win, anything else to lose.")
	case "win":
		term.AddAt(0, 0, "Congratulations, you win!")
		term.AddAt(0, 1, "Press Escape to exit, anything else to play again.")
	case "lose":
		term.AddAt(0, 0, "Sorry, better luck next time.")
		term.AddAt(0, 1, "Press Escape to exit, anything else to play again.")
	}
}

func (game *Game) Draw() {
	game.Clear()
	for _, ui := range game.uis {
		ui.Draw()
	}
}

func (game *Game) Clear() {
	for x := 0; x < *term.Cols; x++ {
		for y := 0; y < *term.Rows; y++ {
			term.AddAt(x, y, " ")
		}
	}
}

func (game *Game) ProcessInput(input int) {
	ui := game.uis[len(game.uis)-1]
	switch ui {
	case "start":
		if input == term.KEY_ENTER || input == 0xa {
			game.uis = []UI{"win"}
		} else {
			game.uis = []UI{"lose"}
		}
	default:
		if input == 0x1b {
			game.uis = []UI{}
		} else {
			game.uis = []UI{"start"}
		}
	}
}

func (game *Game) Run() {
	for {
		if len(game.uis) > 0 {
			game.Draw()
			game.ProcessInput(term.GetChar())
		} else {
			break
		}
	}
}

func NewGame() *Game {
	return &Game{uis: []UI{"start"}}
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
