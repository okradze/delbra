package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		return handleKey(m, msg)
	}

	return m, nil
}

func handleKey(m Model, msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c", "q", "esc":
		return m, tea.Quit

	case "up", "k":
		if m.cursor > 0 {
			m.cursor--
		}

	case "down", "j":
		if m.cursor < len(m.branches)-1 {
			m.cursor++
		}

	case " ":
		return handleSpaceKey(m)
	case "enter":
		return handleEnterKey(m)
	default:
		if m.showConfirmation {
			m.confirmationInput, _ = m.confirmationInput.Update(msg)
		}

	}

	return m, nil
}

func handleSpaceKey(m Model) (tea.Model, tea.Cmd) {
	if m.showConfirmation {
		return m, nil
	}

	_, ok := m.selected[m.cursor]
	if ok {
		delete(m.selected, m.cursor)
	} else {
		m.selected[m.cursor] = struct{}{}
	}

	return m, nil
}

func handleEnterKey(m Model) (tea.Model, tea.Cmd) {
	if len(m.selected) == 0 {
		return m, nil
	}

	if m.showConfirmation {
		val := m.confirmationInput.Value()

		if val == "y" {
			return handleConfirmation(m)
		}

		return m, tea.Quit
	} else {
		m.showConfirmation = true
	}

	return m, nil
}

func handleConfirmation(m Model) (tea.Model, tea.Cmd) {
	branches := []string{}

	for i := range m.selected {
		branches = append(branches, m.branches[i].name)
	}

	DeleteBranches(branches)
	m.deleted = true
	return m, tea.Quit
}
