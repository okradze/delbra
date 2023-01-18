package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/muesli/termenv"
)

var (
	term     = termenv.EnvColorProfile()
	title    = MakeFgStyle("81")
	selected = MakeFgStyle("123")
	hovering = MakeFgStyle("225")
	subtle   = MakeFgStyle("241")
	dot      = ColorFg(" • ", "236")
)

func main() {
	p := tea.NewProgram(initialModel())

	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}

	// deleteBranches(branches)
}

type model struct {
	branches []string
	cursor   int
	selected map[int]struct{}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
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

		// The "enter" key and the spacebar (a literal space) toggle
		// the selected state for the item that the cursor is pointing at.
		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}

	return m, nil
}

func (m model) View() string {
	s := title("Select Git Branches") + "\n\n"

	for i, branch := range m.branches {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		checked := " "
		_, ok := m.selected[i]

		if ok {
			checked = "x"
		}

		line := fmt.Sprintf("%s [%s] %s", cursor, checked, branch)

		if ok {
			s += selected(line)
		} else if m.cursor == i {
			s += hovering(line)
		} else {
			s += line
		}

		s += "\n"
	}

	s += subtle("\nj/k, up/down: select") + dot + subtle("enter: choose") + dot + subtle("q, esc: quit")
	s += "\n"

	return s
}

func initialModel() model {
	branches := ParseBranches()

	return model{
		branches: branches,
		selected: make(map[int]struct{}),
		cursor:   0,
	}
}
