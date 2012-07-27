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

func Min(a ...int) int {
	min := int(^uint(0) >> 1) // largest int
	for _, i := range a {
		if i < min {
			min = i
		}
	}
	return min
}

func (world World) GetTile(x int, y int) Tile {
	if x >= 0 && x < WORLD_COLS && y >= 0 && y < WORLD_ROWS {
		return world[y][x]
	}
	return TILES["bound"]
}

func (world World) SmoothWorld() World {
	newWorld := make(World, WORLD_ROWS)

	for y, row := range world {
		newWorld[y] = world.SmoothRow(row, y)
	}
	return newWorld
}

func (world World) SmoothRow(row []Tile, y int) []Tile {
	var floorCount int
	newRow := make([]Tile, WORLD_COLS)

	for x := range row {
		// for each tile
		// if the 3x3 block centered on it contains 5 or more floors
		floorCount = 0
		for _, tile := range world.GetTileBlock(x, y) {
			if tile == TILES["floor"] {
				floorCount += 1
			}
		}

		// then the tile is a floor
		// otherwise it is a wall
		if floorCount >= 5 {
			newRow[x] = TILES["floor"]
		} else {
			newRow[x] = TILES["wall"]
		}
	}

	return newRow
}

func (world World) GetTileBlock(x int, y int) []Tile {
	tiles := make([]Tile, 9)
	for dx := x - 1; dx <= x+1; dx++ {
		for dy := y - 1; dy <= y+1; dy++ {
			tiles = append(tiles, world.GetTile(dx, dy))
		}
	}
	return tiles
}

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
			term.AddAt(x, y, fmt.Sprintf("%c", game.world.GetTile(x, y).glyph))
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
