package below

import (
	"./ui"
	"code.google.com/p/termon"
)

type Screen string
type Game struct {
	screens []Screen
	world   World
}

func (game *Game) ProcessInput(input int) {
	screen := game.screens[len(game.screens)-1]
	switch screen {
	case "start":
		game.world = RandomWorld()
		game.screens = []Screen{"play"}
	case "play":
		if input == ui.KEY_LF || input == ui.KEY_CR {
			game.screens = []Screen{"win"}
		} else if input == 's' {
			game.world = game.world.SmoothWorld()
			game.screens = []Screen{"play"}
		} else {
			game.screens = []Screen{"lose"}
		}
	default:
		if input == ui.KEY_BACKSPACE || input == ui.KEY_ALT_BACKSPACE || input == ui.KEY_DELETE {
			game.screens = []Screen{}
		} else {
			game.world = RandomWorld()
			game.screens = []Screen{"play"}
		}
	}
}

func (game *Game) Run() {
	for {
		if len(game.screens) > 0 {
			game.Draw()
			game.ProcessInput(term.GetChar())
		} else {
			break
		}
	}
}

func NewGame() *Game {
	return &Game{screens: []Screen{"start"}}
}

func (game *Game) Draw() {
	ui.Clear()
	for _, screen := range game.screens {
		screen.Draw(game)
	}
}

func (screen Screen) Draw(game *Game) {
	switch screen {
	case "start":
		ui.DrawWithColor(0, 0, "Welcome to Below!", "red")
		ui.Draw(0, 1, "Press any key to start.")
	case "win":
		ui.Draw(0, 0, "Congratulations, you win!")
		ui.Draw(0, 1, "Press Backspace to exit, anything else to play again.")
	case "lose":
		ui.Draw(0, 0, "Sorry, better luck next time.")
		ui.Draw(0, 1, "Press Backspace to exit, anything else to play again.")
	case "play":
		game.world.Draw(game)
		ui.DrawCrosshairs()
	}
}
