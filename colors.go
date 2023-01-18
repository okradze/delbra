package main

import (
	"github.com/muesli/termenv"
)

var term = termenv.EnvColorProfile()

// Color a string's foreground with the given value.
func ColorFg(val, color string) string {
	return termenv.String(val).Foreground(term.Color(color)).String()
}

// Return a function that will colorize the foreground of a given string.
func MakeFgStyle(color string) func(string) string {
	return termenv.Style{}.Foreground(term.Color(color)).Styled
}
