package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if len(m.branches) == 0 {
		return m, tea.Quit
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		// These keys should exit the program.
		case "ctrl+c", "q", "esc":
			return m, tea.Quit

		// The "up" and "k" keys move the cursor up
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if m.cursor < len(m.branches)-1 {
				m.cursor++
			}

		// the spacebar (a literal space) toggle
		case " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}

		// Delete branches on enter
		case "enter":
			if len(m.selected) == 0 {
				return m, nil
			}

			branches := []string{}

			for i := range m.selected {
				branches = append(branches, m.branches[i])
			}

			DeleteBranches(branches)

			m.deleted = true
			return m, tea.Quit
		}

	}

	return m, nil
}
