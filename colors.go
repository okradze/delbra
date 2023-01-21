package main

import (
	"github.com/muesli/termenv"
)

var term = termenv.EnvColorProfile()

var (
	titleFg     = MakeFgStyle("81")
	lightTextFg = MakeFgStyle("253")
	errorFg     = MakeFgStyle("160")
	successFg   = MakeFgStyle("155")
	selectedFg  = MakeFgStyle("123")
	hoveringFg  = MakeFgStyle("225")
	subtleFg    = MakeFgStyle("241")
	dot         = ColorFg(" • ", "236")
)

// Color a string's foreground with the given value.
func ColorFg(val, color string) string {
	return termenv.String(val).Foreground(term.Color(color)).String()
}

// Return a function that will colorize the foreground of a given string.
func MakeFgStyle(color string) func(string) string {
	return termenv.Style{}.Foreground(term.Color(color)).Styled
}
