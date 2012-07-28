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

type UI string
type Game struct {
	uis   []UI
	world World
}

func (game *Game) ProcessInput(input int) {
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
