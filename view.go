package main

import "fmt"

var (
	titleFg    = MakeFgStyle("81")
	selectedFg = MakeFgStyle("123")
	hoveringFg = MakeFgStyle("225")
	subtleFg   = MakeFgStyle("241")
	dot        = ColorFg(" • ", "236")
)

func (m Model) View() string {
	s := titleFg("Select Git Branches") + "\n\n"
	s += formatBranchList(m)
	s += formatHelp()

	return s
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

func formatHelp() string {
	s := subtleFg("\nj/k, up/down: select") + dot + subtleFg("enter: choose") + dot + subtleFg("q, esc: quit")
	s += "\n"

	return s
}
