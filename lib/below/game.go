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
		game.Reset()
	case "play":
		switch input {
		case 'y':
			game.world = game.world.MovePlayer("nw")
		case 'k':
			game.world = game.world.MovePlayer("n")
		case 'u':
			game.world = game.world.MovePlayer("ne")
		case 'l':
			game.world = game.world.MovePlayer("e")
		case 'n':
			game.world = game.world.MovePlayer("se")
		case 'j':
			game.world = game.world.MovePlayer("s")
		case 'b':
			game.world = game.world.MovePlayer("sw")
		case 'h':
			game.world = game.world.MovePlayer("w")
		case ui.KEY_LF, ui.KEY_CR:
			game.screens = []Screen{"win"}
		default:
			game.screens = []Screen{"lose"}
		}
	default:
		if input == ui.KEY_BACKSPACE || input == ui.KEY_ALT_BACKSPACE || input == ui.KEY_DELETE {
			game.screens = []Screen{}
		} else {
			game.Reset()
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

func (game *Game) Reset() {
	game.world = RandomWorld()
	game.screens = []Screen{"play"}
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

func (game *Game) GetViewportCoords() (startX, startY, endX, endY int) {
	cols := ui.Cols()
	// Leave a row for status.
	rows := ui.Rows() - 1

	player := game.world.player
	centerX := player.location.X()
	centerY := player.location.Y()

	startX = Max(0, centerX-(cols/2))
	startY = Max(0, centerY-(rows/2))

	endX = Min(WORLD_COLS, startX+cols)
	endY = Min(WORLD_ROWS, startY+rows)

	// If I truncated the end coordinate I’ll have ended up with a
	// smaller-than-normal viewport. To fix that I’ll reset the start
	// coordinates one more time.
	startX = endX - cols
	startY = endY - rows

	return startX, startY, endX, endY
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
	}
}
