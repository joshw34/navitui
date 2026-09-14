package ui

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
	"github.com/joshw34/navitui/internal/types"
)

type artistsModel struct {
	cursor int
	data   []types.Artist
}

func (a artistsModel) Update(msg tea.Msg) (artistsModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "up", "k":
			if a.cursor > 0 {
				a.cursor--
			}
		case "down", "j":
			if a.cursor < len(a.data)-1 {
				a.cursor++
			}
		}
	}
	return a, nil
}

func (a artistsModel) View() string {
	var s string
	for i, artist := range a.data {
		cursor := " "
		if a.cursor == i {
			cursor = ">"
		}

		s += fmt.Sprintf("%s %s\n", cursor, artist.Name)
	}
	s += "\nPress q to quit.\n"
	return s
}
