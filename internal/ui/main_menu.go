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
func (a mainItem) Description() string { return "" }
func (a mainItem) FilterValue() string { return a.option }

type mainModel struct {
	list list.Model
}

func onKeypressMain(key tea.KeyPressMsg, it list.Item, index int) tea.Cmd {
	_ = index
	if key.String() == "enter" {
		switch it.(mainItem).option {
		case "Artists":
			return func() tea.Msg { return getAllArtistsMsg{} }
		case "Albums":
			return func() tea.Msg { return getAllAlbumsMsg{} }
		case "Songs":
			return func() tea.Msg { return getAllSongsMsg{} }
		}
	}
	return nil
}

func mainOptionsToListItem() []list.Item {
	options := []mainItem{
		{option: "Artists", action: artists},
		{option: "Albums", action: albums},
		{option: "Songs", action: songs},
	}
	items := make([]list.Item, len(options))
	for i, opt := range options {
		items[i] = opt
	}
	return items
}
