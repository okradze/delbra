package main

import (
	"fmt"
)

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

func (m Model) View() string {
	if len(m.branches) == 0 {
		return errorFg("No branches to delete") + "\n"
	}

	s := formatTitle()
	s += formatBranchList(m)
	s += formatConfirmation(m)
	s += formatHelp()

	if m.deleted {
		s += successFg("\nAll selected branches deleted") + "\n"
	}

	return s
}

func formatTitle() string {
	return titleFg("\nWhich branches do you want to delete?") + "\n\n"
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

		line := fmt.Sprintf("%s %s %s", cursor, checkbox, branch)

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
