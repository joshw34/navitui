package ui

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
)

type mainModel struct {
	choices []string
	cursor  int
}

func (m mainModel) Update(msg tea.Msg) (mainModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter":
			return m, func() tea.Msg {
				return getArtistsMsg{}
			}
		}
	}
	return m, nil
}

func (m mainModel) View() string {
	var s string
	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		s += fmt.Sprintf("%s %s\n", cursor, choice)
		s += "\nPress q to quit.\n"
	}
	return s
}
