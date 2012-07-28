package below

import (
	"code.google.com/p/goncurses"
	// "fmt"
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
	uis    []UI
	world  World
	window goncurses.Window
}

func NewGame(window goncurses.Window) *Game {
	return &Game{uis: []UI{"start"}, window: window}
}

func (game *Game) Run() {
	for {
		if len(game.uis) > 0 {
			game.DrawUIs()
			input := game.window.GetChar()
			// game.Draw(0, 0, fmt.Sprintf("key: %d", input))
			// game.Draw(0, 1, fmt.Sprintf("enter: %d", goncurses.KEY_ENTER))
			// game.window.GetChar()
			game.ProcessInput(input)
		} else {
			break
		}
	}
}

func (game *Game) ProcessInput(input goncurses.Key) {
	ui := game.uis[len(game.uis)-1]
	switch ui {
	case "start":
		game.world = RandomWorld()
		game.uis = []UI{"play"}
	case "play":
		if input == 10 { // goncurses.KEY_ENTER {
			game.uis = []UI{"win"}
		} else if input == 's' {
			game.world = game.world.SmoothWorld()
			game.uis = []UI{"play"}
		} else {
			game.uis = []UI{"lose"}
		}
	default:
		if input == 127 { // goncurses.KEY_BACKSPACE || input == goncurses.KEY_DC {
			game.uis = []UI{}
		} else {
			game.world = RandomWorld()
			game.uis = []UI{"play"}
		}
	}
}
