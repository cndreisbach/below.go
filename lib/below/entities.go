package below

import (
	"./ui"
	"fmt"
)

type Entity interface {
	Tick(World) World
}

type Moveable interface {
	Move(World, Coords) World
	CanMove(World, Coords) bool
}

type Digger interface {
	Dig(World, Coords) World
	CanDig(World, Coords) bool
}

type Player struct {
	glyph    rune
	location Coords
}

func NewPlayer(world World) Player {
	return Player{'@', world.FindEmptyCoords()}
}

func (player Player) Tick(world World) World {
	return world
}

func (player Player) CanMove(world World, coords Coords) bool {
	tile := world.GetTile(coords)
	return tile.kind == "floor"
}

func (player Player) Move(world World, coords Coords) World {
	if player.CanMove(world, coords) {
		player.location = coords
		world.player = player
	}
	return world
}

func (player Player) CanDig(world World, coords Coords) bool {
	tile := world.GetTile(coords)
	return tile.kind == "wall"
}

func (player Player) Dig(world World, coords Coords) World {
	if player.CanDig(world, coords) {
		world.SetTile(coords, TILES["floor"])
	}
	return world
}

func (player Player) MoveOrDig(world World, coords Coords) (newWorld World) {
	if player.CanMove(world, coords) {
		newWorld = player.Move(world, coords)
	} else {
		newWorld = player.Dig(world, coords)
	}
	return newWorld
}

func (player Player) Draw(game *Game) {
	startX, startY, _, _ := game.GetViewportCoords()
	x := player.location.X() - startX
	y := player.location.Y() - startY

	ui.DrawWithColor(x, y, fmt.Sprintf("%c", player.glyph), "red")
}
