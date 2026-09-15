package ui

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
	"github.com/joshw34/navitui/internal/types"
)

type albumModel struct {
	cursor int
	data   []types.Song
}

func (a albumModel) Update(msg tea.Msg) (albumModel, tea.Cmd) {
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
			/*case "enter":
			return a, func() tea.Msg { return getArtistMsg{artistID: a.data[a.cursor].ID} }*/
		}
	}
	return a, nil
}

func (a albumModel) View() string {
	var s string
	for i, song := range a.data {
		cursor := " "
		if a.cursor == i {
			cursor = ">"
		}

		s += fmt.Sprintf("%s %s\n", cursor, song.Title)
	}
	s += "\nPress q to quit.\n"
	return s
}
