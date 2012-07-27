package main

import (
	"code.google.com/p/termon"
	"fmt"
)

type UI string
type Game struct {
	uis   []UI
	world World
}

func (ui UI) Draw(game *Game) {
	switch ui {
	case "start":
		term.AddAt(0, 0, "Welcome to Below!")
		term.AddAt(0, 1, "Press any key to start.")
	case "win":
		term.AddAt(0, 0, "Congratulations, you win!")
		term.AddAt(0, 1, "Press Backspace to exit, anything else to play again.")
	case "lose":
		term.AddAt(0, 0, "Sorry, better luck next time.")
		term.AddAt(0, 1, "Press Backspace to exit, anything else to play again.")
	case "play":
		game.world.Draw(game)
	}
}

func (game *Game) Draw() {
	game.Clear()
	for _, ui := range game.uis {
		ui.Draw(game)
	}
}

func (game *Game) Clear() {
	for x := 0; x < *term.Cols; x++ {
		// Do not clear status line for now.
		for y := 0; y < *term.Rows-1; y++ {
			term.AddAt(x, y, " ")
		}
	}
}

func (game *Game) ProcessInput(input int) {
	// Temporarily print input on status line
	term.AddAt(0, *term.Rows-1, "               ")
	term.AddAt(0, *term.Rows-1, fmt.Sprintf("%c", input))

	ui := game.uis[len(game.uis)-1]
	switch ui {
	case "start":
		game.world = RandomWorld()
		game.uis = []UI{"play"}
	case "play":
		if input == LF || input == CR {
			game.uis = []UI{"win"}
		} else if input == 's' {
			game.world = game.world.SmoothWorld()
			game.uis = []UI{"play"}
		} else {
			game.uis = []UI{"lose"}
		}
	default:
		if input == term.KEY_BACKSPACE || input == BACKSPACE || input == DELETE {
			game.uis = []UI{}
		} else {
			game.world = RandomWorld()
			game.uis = []UI{"play"}
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
