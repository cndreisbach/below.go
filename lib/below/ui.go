package below

import (
	"code.google.com/p/goncurses"
	"fmt"
)

const (
	BG_COLOR = goncurses.C_BLACK
	FG_COLOR = goncurses.C_WHITE
)

var (
	currentColorPair = byte(0)
	Colors           = make(map[string]byte, 10)
)

func NewColor(fg, bg int) byte {
	currentColorPair++
	goncurses.InitPair(currentColorPair, fg, bg)
	return currentColorPair
}

func SetupColors() {
	Colors["default"] = NewColor(FG_COLOR, BG_COLOR)
	Colors["black"] = NewColor(goncurses.C_BLACK, goncurses.C_WHITE)
	Colors["white"] = NewColor(goncurses.C_WHITE, BG_COLOR)
	Colors["red"] = NewColor(goncurses.C_RED, BG_COLOR)
}

func (game *Game) Draw(x, y int, text string) {
	game.window.Print(y, x, text)
}

func (game *Game) DrawWithColor(x, y int, text string, colorVal string) {
	color := Colors[colorVal]
	game.window.ColorOn(color)
	game.window.Print(y, x, text)
	game.window.ColorOff(color)
}

func (game *Game) Clear() {
	game.window.Clear()
}

func (game *Game) DrawCrosshairs() {
	x := game.Cols() / 2
	y := game.Rows() / 2
	game.DrawWithColor(x, y, "X", "red")
}

func (game *Game) DrawUIs() {
	game.Clear()
	for _, ui := range game.uis {
		ui.Draw(game)
	}
}

func (game *Game) Cols() int {
	_, cols := game.window.Maxyx()
	return cols - 1
}

func (game *Game) Rows() int {
	rows, _ := game.window.Maxyx()
	return rows - 1
}

func (world World) Draw(game *Game) {
	// Leave a row for status.
	rows := game.Rows() - 1
	cols := game.Cols()
	startX := 0
	startY := 0
	endX := Min(WORLD_COLS, startX+cols)
	endY := Min(WORLD_ROWS, startY+rows)

	var tile Tile

	for y := startY; y < endY; y++ {
		for x := startX; x < endX; x++ {
			tile = game.world.GetTile(x, y)
			game.DrawWithColor(x, y, fmt.Sprintf("%c", tile.glyph), tile.color)
		}
	}
}

func (ui UI) Draw(game *Game) {
	switch ui {
	case "start":
		game.DrawWithColor(0, 0, "Welcome to Below!", "red")
		game.Draw(0, 1, "Press any key to start.")
		game.Draw(0, 3, fmt.Sprintf("HasColors: %v", goncurses.HasColors()))
		game.Draw(0, 4, fmt.Sprintf("Colors: %v", Colors))
	case "win":
		game.Draw(0, 0, "Congratulations, you win!")
		game.Draw(0, 1, "Press Backspace to exit, anything else to play again.")
	case "lose":
		game.Draw(0, 0, "Sorry, better luck next time.")
		game.Draw(0, 1, "Press Backspace to exit, anything else to play again.")
	case "play":
		game.world.Draw(game)
		game.DrawCrosshairs()
	}
}
