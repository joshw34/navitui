package ui

import (
	"fmt"

	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
	"github.com/joshw34/navitui/internal/types"
)

type albumsItem struct {
	album types.Album
}

func (a albumsItem) Title() string { return a.album.Name }

func (a albumsItem) Description() string {
	return fmt.Sprintf("%s", yearLabel(a.album.Year))
}

func (a albumsItem) FilterValue() string { return a.album.Name }

func (a albumsModel) buildList(data []types.Album) albumsModel {
	items := make([]list.Item, len(data))
	for i, album := range data {
		items[i] = albumsItem{album: album}
	}
	a.list.SetItems(items)
	return a
}

type albumsModel struct {
	list list.Model
}

func (a albumsModel) Update(msg tea.Msg) (albumsModel, tea.Cmd) {
	if key, ok := msg.(tea.KeyPressMsg); ok && key.String() == "enter" {
		if it, ok := a.list.SelectedItem().(albumsItem); ok {
			return a, func() tea.Msg { return getSongsByAlbumMsg{albumID: it.album.ID} }
		}
	}

	var cmd tea.Cmd
	a.list, cmd = a.list.Update(msg)
	return a, cmd
}

func (a albumsModel) View() string {
	return a.list.View()
}
