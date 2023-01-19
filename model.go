package main

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	branches          []string
	cursor            int
	selected          map[int]struct{}
	showConfirmation  bool
	confirmationInput textinput.Model
	deleted           bool
}

func (m Model) Init() tea.Cmd {
	if m.showConfirmation {
		return textinput.Blink
	}

	return nil
}

func InitialModel() Model {
	branches := ParseBranches()

	return Model{
		branches:          branches,
		selected:          make(map[int]struct{}),
		cursor:            0,
		deleted:           false,
		showConfirmation:  false,
		confirmationInput: InitialConfirmationInputModel(),
	}
}

func InitialConfirmationInputModel() textinput.Model {
	ti := textinput.NewModel()
	ti.Focus()
	ti.TextStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#f8f8f2"))
	ti.CharLimit = 1
	ti.Width = 1
	return ti
}
