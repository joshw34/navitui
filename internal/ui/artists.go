package ui

import (
	"fmt"

	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
	"github.com/joshw34/navitui/internal/types"
)

type artistsItem struct {
	artist types.Artist
}

func (a artistsItem) Title() string { return a.artist.Name }

func (a artistsItem) Description() string {
	return fmt.Sprintf("%d %s", a.artist.AlbumCount, pluralize("album", a.artist.AlbumCount))
}

func (a artistsItem) FilterValue() string { return a.artist.Name }

func onKeypressArtists(key tea.KeyPressMsg, it list.Item) tea.Cmd {
	a := it.(artistsItem)
	switch key.String() {
	case "enter":
		return func() tea.Msg { return getAlbumsByArtistMsg{a.artist.ID} }
	}
	return nil
}

func artistsToListItems(artists []types.Artist) []list.Item {
	items := make([]list.Item, len(artists))
	for i, a := range artists {
		items[i] = artistsItem{artist: a}
	}
	return items
}
