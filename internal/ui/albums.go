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

func onSelectedAlbums(it list.Item) tea.Cmd {
	return func() tea.Msg { return getSongsByAlbumMsg{it.(albumsItem).album.ID} }
}

func albumsToListItems(albums []types.Album) []list.Item {
	items := make([]list.Item, len(albums))
	for i, a := range albums {
		items[i] = albumsItem{album: a}
	}
	return items
}
