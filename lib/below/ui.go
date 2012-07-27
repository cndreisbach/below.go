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
	color := Colors[colorVal]
	//color.On()
	term.AddAt(x, y, text)
	term.AddAt(*term.Cols-1, *term.Rows-1, " ")
	//color.Off()
}
