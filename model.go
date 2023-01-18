package main

import tea "github.com/charmbracelet/bubbletea"

type Model struct {
	branches []string
	cursor   int
	selected map[int]struct{}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func InitialModel() Model {
	branches := ParseBranches()

	return Model{
		branches: branches,
		selected: make(map[int]struct{}),
		cursor:   0,
	}
}
