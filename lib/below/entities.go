package below

import (
	"./ui"
	"fmt"
)

type Entity interface {
	Tick(world World) World
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

func (player Player) Draw(game *Game) {
	startX, startY, _, _ := game.GetViewportCoords()
	x := player.location.X() - startX
	y := player.location.Y() - startY

	ui.DrawWithColor(x, y, fmt.Sprintf("%c", player.glyph), "red")
}
