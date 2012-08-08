package below

import (
	"./ui"
	"fmt"
	"math/rand"
)

type World struct {
	tiles  [][]Tile
	player Player
}

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

func RandomTile() Tile {
	tiles := []string{"floor", "wall"}
	return TILES[tiles[rand.Intn(len(tiles))]]
}

func RandomWorld() (world World) {
	tiles := make([][]Tile, WORLD_ROWS)
	for y := range tiles {
		tiles[y] = make([]Tile, WORLD_COLS)
		for x := range tiles[y] {
			tiles[y][x] = RandomTile()
		}
	}
	world.tiles = tiles

	for i := 0; i < 3; i++ {
		world.SmoothWorld()
	}

	world.player = NewPlayer(world)

	return world
}

func (world World) GetTile(coords Coords) Tile {
	x, y := coords.X(), coords.Y()
	if x >= 0 && x < WORLD_COLS && y >= 0 && y < WORLD_ROWS {
		return world.tiles[y][x]
	}
	return TILES["bound"]
}

func (world World) FindEmptyCoords() Coords {
	coords := RandomCoords()
	tile := world.GetTile(coords)
	if tile.kind != "floor" {
		return world.FindEmptyCoords()
	}
	return coords
}

func (world *World) SmoothWorld() {
	newTiles := make([][]Tile, WORLD_ROWS)

	for y, row := range world.tiles {
		newTiles[y] = world.smoothRow(row, y)
	}

	world.tiles = newTiles
}

func (world *World) smoothRow(row []Tile, y int) []Tile {
	var floorCount int
	newRow := make([]Tile, WORLD_COLS)

	// for each tile in row
	for x := range row {
		// if the 3x3 block centered on it contains 5 or more floors
		floorCount = 0
		for _, tile := range world.GetTileBlock(Coords{x, y}) {
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

func (world World) GetTileBlock(coords Coords) []Tile {
	x, y := coords.X(), coords.Y()
	tiles := make([]Tile, 9)
	for dx := x - 1; dx <= x+1; dx++ {
		for dy := y - 1; dy <= y+1; dy++ {
			tiles = append(tiles, world.GetTile(Coords{dx, dy}))
		}
	}
	return tiles
}

func (world World) MovePlayer(direction string) World {
	player := world.player
	return player.Move(world, DestinationCoords(player.location, direction))
}

func (world World) Draw(game *Game) {
	startX, startY, endX, endY := game.GetViewportCoords()

	cols := endX - startX
	rows := endY - startY

	var tile Tile

	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			tile = game.world.GetTile(Coords{x + startX, y + startY})
			ui.DrawWithColor(x, y, fmt.Sprintf("%c", tile.glyph), tile.color)
		}
	}

	world.player.Draw(game)
	ui.Draw(0, ui.Rows()-1, fmt.Sprint(world.player.location))
}
