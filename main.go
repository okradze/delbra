package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(InitialModel(), tea.WithAltScreen())

	m, err := p.Run()

	if err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}

	if m.(Model).deleted {
		fmt.Println(successFg("All selected branches deleted"))
	}
}
