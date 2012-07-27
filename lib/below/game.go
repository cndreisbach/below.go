package below

import (
	"code.google.com/p/termon"
	"fmt"
)

const (
	BACKSPACE = 8
	LF        = 10
	CR        = 13
	ESCAPE    = 27
	DELETE    = 127
)

type UI string
type Game struct {
	uis   []UI
	world World
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
			Draw(x, y, " ")
		}
	}
}

func (game *Game) DrawCrosshairs() {
	x := *term.Cols / 2
	y := *term.Rows / 2
	DrawWithColor(x, y, "X", "red")
}

func (ui UI) Draw(game *Game) {
	switch ui {
	case "start":
		DrawWithColor(0, 0, "Welcome to Below!", "red")
		Draw(0, 1, "Press any key to start.")
	case "win":
		Draw(0, 0, "Congratulations, you win!")
		Draw(0, 1, "Press Backspace to exit, anything else to play again.")
	case "lose":
		Draw(0, 0, "Sorry, better luck next time.")
		Draw(0, 1, "Press Backspace to exit, anything else to play again.")
	case "play":
		game.world.Draw(game)
		game.DrawCrosshairs()
	}
}

func (game *Game) ProcessInput(input int) {
	// Temporarily print input on status line
	Draw(0, *term.Rows-1, "               ")
	Draw(0, *term.Rows-1, fmt.Sprint(input))

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
