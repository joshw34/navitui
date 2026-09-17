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

func (a artistsModel) buildList(data []types.Artist) artistsModel {
	items := make([]list.Item, len(data))
	for i, artist := range data {
		items[i] = artistsItem{artist: artist}
	}
	a.list.SetItems(items)
	return a
}

type artistsModel struct {
	list list.Model
}

func (a artistsModel) Update(msg tea.Msg) (artistsModel, tea.Cmd) {
	if key, ok := msg.(tea.KeyPressMsg); ok && key.String() == "enter" {
		if it, ok := a.list.SelectedItem().(artistsItem); ok {
			return a, func() tea.Msg { return getAlbumsByArtistMsg{artistID: it.artist.ID} }
		}
	}

	var cmd tea.Cmd
	a.list, cmd = a.list.Update(msg)
	return a, cmd
}

func (a artistsModel) View() string {
	return a.list.View()
}
