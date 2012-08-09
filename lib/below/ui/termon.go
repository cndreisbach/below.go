package ui

import (
	"code.google.com/p/termon"
)

const (
	BG_COLOR = term.COLOR_BLACK
	FG_COLOR = term.COLOR_WHITE

	KEY_BACKSPACE     = 8
	KEY_LF            = 10
	KEY_CR            = 13
	KEY_ESCAPE        = 27
	KEY_DELETE        = 127
	KEY_ALT_BACKSPACE = term.KEY_BACKSPACE
)

var (
	Colors = make(map[string]term.Color, 10)
)

func Init() {
	term.Init()

	// Allows use of function keys and arrow keys.
	term.Keypad()

	// Suppress user input.
	term.Noecho()

	SetupColors()
}

func End() {
	// Reset the terminal.
	// It will look as well as before term.Init() was called.
	term.End()
}

func SetupColors() {
	Colors["default"] = term.NewColor(FG_COLOR, BG_COLOR)
	Colors["black"] = term.NewColor(term.COLOR_BLACK, BG_COLOR)
	Colors["white"] = term.NewColor(term.COLOR_WHITE, BG_COLOR)
	Colors["red"] = term.NewColor(term.COLOR_RED, BG_COLOR)
}

func Cols() int {
	return *term.Cols
}

func Rows() int {
	return *term.Rows
}

func Draw(x, y int, text string) {
	DrawWithColor(x, y, text, "default")
}

func DrawWithColor(x, y int, text string, colorVal string) {
	color := Colors[colorVal]
	color.On()
	term.AddAt(x, y, text)
	term.AddAt(*term.Cols-1, *term.Rows-1, " ")
	color.Off()
}

func Clear() {
	for x := 0; x < *term.Cols; x++ {
		for y := 0; y < *term.Rows; y++ {
			Draw(x, y, " ")
		}
	}
}

func DrawCrosshairs() {
	x := *term.Cols / 2
	y := *term.Rows / 2
	DrawWithColor(x, y, "X", "red")
}
