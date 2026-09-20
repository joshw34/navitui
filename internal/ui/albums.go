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
	return fmt.Sprintf("%s - %s", yearLabel(a.album.Year), a.album.Artist)
}

func (a albumsItem) FilterValue() string { return a.album.Name }

func onKeypressAlbums(key tea.KeyPressMsg, it list.Item) tea.Cmd {
	a := it.(albumsItem)
	switch key.String() {
	case "enter":
		return func() tea.Msg { return getSongsByAlbumMsg{a.album.ID} }
	}
	return nil
}

func albumsToListItems(albums []types.Album) []list.Item {
	items := make([]list.Item, len(albums))
	for i, a := range albums {
		items[i] = albumsItem{album: a}
	}
	return items
}
