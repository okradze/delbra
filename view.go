package main

import (
	"fmt"
)

func (m Model) View() string {
	s := formatTitle()
	s += formatBranchList(m)
	s += formatConfirmation(m)
	s += formatHelp()

	return s
}

func formatTitle() string {
	return "\n" + titleFg("Which branches do you want to delete?") + "\n\n"
}

func formatBranchList(m Model) string {
	s := ""

	for i, branch := range m.branches {
		hovering := m.cursor == i

		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		checkbox := "[ ]"
		_, checked := m.selected[i]

		if checked {
			checkbox = "[x]"
		}

		merged := ""

		if !branch.merged {
			merged = errorFg("(not merged)")
		}

		line := fmt.Sprintf("%s %s %s %s", cursor, checkbox, branch.name, merged)

		if checked {
			s += selectedFg(line)
		} else if hovering {
			s += hoveringFg(line)
		} else {
			s += line
		}

		s += "\n"
	}

	return s
}

func formatConfirmation(m Model) string {
	s := ""

	if m.showConfirmation {
		s += "\n"
		s += lightTextFg("Are you sure you want to delete the selected branches? (y/n) ")
		s += m.confirmationInput.View() + "\n"
	}

	return s
}

func formatHelp() string {
	return subtleFg("\nj/k, up/down: select") + dot + subtleFg("space: choose") + dot + subtleFg("enter: delete") + dot + subtleFg("q, esc: quit") + "\n"
}
