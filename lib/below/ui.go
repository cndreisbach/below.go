package below

import (
	"code.google.com/p/termon"
	"fmt"
)

const (
	BG_COLOR = term.COLOR_BLUE
	FG_COLOR = term.COLOR_WHITE
)

var (
	Colors = map[string]term.Color{
		"black":   term.NewColor(term.COLOR_BLACK, BG_COLOR),
		"white":   term.NewColor(term.COLOR_WHITE, BG_COLOR),
		"red":     term.NewColor(term.COLOR_RED, BG_COLOR),
		"default": term.NewColor(FG_COLOR, BG_COLOR),
	}
)

func Draw(x, y int, text string) {
	DrawWithColor(x, y, text, "default")
}

func DrawWithColor(x, y int, text string, colorVal string) {
	// color := Colors[colorVal]
	// color.On()
	term.AddAt(x, y, text)
	term.AddAt(*term.Cols-1, *term.Rows-1, " ")
	// color.Off()
}

func Clear() {
	for x := 0; x < *term.Cols; x++ {
		// Do not clear status line for now.
		for y := 0; y < *term.Rows-1; y++ {
			Draw(x, y, " ")
		}
	}
}

func DrawCrosshairs() {
	x := *term.Cols / 2
	y := *term.Rows / 2
	DrawWithColor(x, y, "X", "red")
}

func (game *Game) Draw() {
	Clear()
	for _, ui := range game.uis {
		ui.Draw(game)
	}
}

func (world World) Draw(game *Game) {
	cols := *term.Cols
	// Leave a row for status.
	rows := *term.Rows - 1
	startX := 0
	startY := 0
	endX := Min(WORLD_COLS, startX+cols)
	endY := Min(WORLD_ROWS, startY+rows)

	var tile Tile

	for y := startY; y < endY; y++ {
		for x := startX; x < endX; x++ {
			tile = game.world.GetTile(x, y)
			DrawWithColor(x, y, fmt.Sprintf("%c", tile.glyph), tile.color)
		}
	}
}

func (ui UI) Draw(game *Game) {
	switch ui {
	case "start":
		DrawWithColor(0, 0, "Welcome to Below!", "red")
		Draw(0, 1, "Press any key to start.")
	case "win":
		Draw(0, 0, "Congratulations, you win!")
		Draw(0, 1, "Press Backspace to exit, anything else to play again.")
	case "lose":
		Draw(0, 0, "Sorry, better luck next time.")
		Draw(0, 1, "Press Backspace to exit, anything else to play again.")
	case "play":
		game.world.Draw(game)
		DrawCrosshairs()
	}
}
