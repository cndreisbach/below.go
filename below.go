package main

import (
	"code.google.com/p/termon"
	"fmt"
	"math/rand"
)

type UI string
type World [][]Tile
type Game struct {
	uis   []UI
	world World
}
type Tile struct {
	kind  string
	glyph rune
	color string
}

const (
	BACKSPACE = 8
	LF        = 10
	CR        = 13
	ESCAPE    = 27
	DELETE    = 127
)

var (
	WORLD_COLS = 160
	WORLD_ROWS = 60

	TILES = map[string]Tile{
		"floor": Tile{"floor", '.', "white"},
		"wall":  Tile{"wall", '#', "white"},
		"bound": Tile{"bound", 'X', "black"},
	}
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
		game.world.Draw()
	}
}

func (world World) Draw() {
	cols := *term.Cols
	// Leave a row for status.
	rows := *term.Rows - 1
	startX := 0
	startY := 0
	endX := Min(WORLD_COLS, startX+cols)
	endY := Min(WORLD_ROWS, startY+rows)

	for y := startY; y < endY; y++ {
		for x := startX; x < endX; x++ {
			term.AddAt(x, y, fmt.Sprintf("%c", world[y][x].glyph))
		}
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
		for y := 0; y < *term.Rows-1; y++ {
			term.AddAt(x, y, " ")
		}
	}
}

func (game *Game) ProcessInput(input int) {
	term.AddAt(0, *term.Rows-1, fmt.Sprint(input))
	ui := game.uis[len(game.uis)-1]
	switch ui {
	case "start":
		game.world = RandomWorld()
		game.uis = []UI{"play"}
	case "play":
		if input == LF || input == CR {
			game.uis = []UI{"win"}
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

func RandomTile() Tile {
	tiles := []string{"floor", "wall"}
	return TILES[tiles[rand.Intn(len(tiles))]]
}

func RandomWorld() World {
	world := make(World, WORLD_ROWS)
	for y := range world {
		world[y] = make([]Tile, WORLD_COLS)
		for x := range world[y] {
			world[y][x] = RandomTile()
		}
	}
	return world
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
