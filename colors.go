package main

import (
	"github.com/muesli/termenv"
)

var term = termenv.EnvColorProfile()

var (
	titleFg     = makeFgStyle("81")
	lightTextFg = makeFgStyle("253")
	errorFg     = makeFgStyle("160")
	successFg   = makeFgStyle("155")
	selectedFg  = makeFgStyle("123")
	hoveringFg  = makeFgStyle("225")
	subtleFg    = makeFgStyle("241")
	dot         = colorFg(" • ", "236")
)

// Color a string's foreground with the given value.
func colorFg(val, color string) string {
	return termenv.String(val).Foreground(term.Color(color)).String()
}

// Return a function that will colorize the foreground of a given string.
func makeFgStyle(color string) func(string) string {
	return termenv.Style{}.Foreground(term.Color(color)).Styled
}
