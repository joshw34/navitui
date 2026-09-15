package ui

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
	"github.com/joshw34/navitui/internal/types"
)

type artistModel struct {
	cursor int
	data   []types.Album
}

func (a artistModel) Update(msg tea.Msg) (artistModel, tea.Cmd) {
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
		case "enter":
			return a, func() tea.Msg { return getAlbumMsg{albumID: a.data[a.cursor].ID} }
		}
	}
	return a, nil
}

func (a artistModel) View() string {
	var s string
	for i, album := range a.data {
		cursor := " "
		if a.cursor == i {
			cursor = ">"
		}

		s += fmt.Sprintf("%s %s\n", cursor, album.Name)
	}
	s += "\nPress q to quit.\n"
	return s
}
