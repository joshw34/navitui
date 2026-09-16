package ui

import (
	"fmt"

	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
	"github.com/joshw34/navitui/internal/types"
)

type artistItem struct {
	album types.Album
}

func (a artistItem) Title() string { return a.album.Name }

func (a artistItem) Description() string {
	return fmt.Sprintf("%s", yearLabel(a.album.Year))
}

func (a artistItem) FilterValue() string { return a.album.Name }

func (a artistModel) buildList(data []types.Album) artistModel {
	items := make([]list.Item, len(data))
	for i, album := range data {
		items[i] = artistItem{album: album}
	}
	a.list.SetItems(items)
	return a
}

type artistModel struct {
	list list.Model
}

func (a artistModel) Update(msg tea.Msg) (artistModel, tea.Cmd) {
	if key, ok := msg.(tea.KeyPressMsg); ok && key.String() == "enter" {
		if it, ok := a.list.SelectedItem().(artistItem); ok {
			return a, func() tea.Msg { return getAlbumMsg{albumID: it.album.ID} }
		}
	}

	var cmd tea.Cmd
	a.list, cmd = a.list.Update(msg)
	return a, cmd
}

func (a artistModel) View() string {
	return a.list.View()
}
