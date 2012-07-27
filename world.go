package below

import (
	"code.google.com/p/termon"
	"fmt"
	"math/rand"
)

type World [][]Tile
type Tile struct {
	kind  string
	glyph rune
	color string
}

var (
	WORLD_COLS = 160
	WORLD_ROWS = 60

	TILES = map[string]Tile{
		"floor": Tile{"floor", '.', "white"},
		"wall":  Tile{"wall", '#', "white"},
		"bound": Tile{"bound", 'X', "black"},
	}
)

func (world World) Draw(game *Game) {
	cols := *term.Cols
	// Leave a row for status.
	rows := *term.Rows - 1
	startX := 0
	startY := 0
	endX := Min(WORLD_COLS, startX+cols)
	endY := Min(WORLD_ROWS, startY+rows)

	for y := startY; y < endY; y++ {
		for x := startX; x < endX; x++ {
			term.AddAt(x, y, fmt.Sprintf("%c", game.GetTile(x, y).glyph))
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
