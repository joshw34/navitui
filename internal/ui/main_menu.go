package ui

import (
	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
)

// in artists_menu.go
type mainItem struct {
	option string
	action page
}

func (a mainItem) Title() string       { return a.option }
func (a mainItem) Description() string { return "" }       // or genre, etc.
func (a mainItem) FilterValue() string { return a.option } // what filtering matches on

type mainModel struct {
	list list.Model
}

func (m mainModel) Update(msg tea.Msg) (mainModel, tea.Cmd) {
	if key, ok := msg.(tea.KeyPressMsg); ok && key.String() == "enter" {
		if it, ok := m.list.SelectedItem().(mainItem); ok {
			switch it.action {
			case artists:
				return m, func() tea.Msg { return getAllArtistsMsg{} }
			case albums:
				return m, func() tea.Msg { return getAllAlbumsMsg{} }
			case songs:
				return m, func() tea.Msg { return getAllSongsMsg{} }
			}
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m mainModel) View() string {
	return m.list.View()
}
