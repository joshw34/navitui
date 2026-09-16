package ui

import (
	"fmt"

	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
	"github.com/joshw34/navitui/internal/types"
)

type artistsListItem struct {
	artist types.Artist
}

func (a artistsListItem) Title() string { return a.artist.Name }

func (a artistsListItem) Description() string {
	return fmt.Sprintf("%d %s", a.artist.AlbumCount, pluralize("album", a.artist.AlbumCount))
}

func (a artistsListItem) FilterValue() string { return a.artist.Name }

func (a artistsListModel) buildList(data []types.Artist) artistsListModel {
	items := make([]list.Item, len(data))
	for i, artist := range data {
		items[i] = artistsListItem{artist: artist}
	}
	a.list.SetItems(items)
	return a
}

type artistsListModel struct {
	list list.Model
}

func (a artistsListModel) Update(msg tea.Msg) (artistsListModel, tea.Cmd) {
	if key, ok := msg.(tea.KeyPressMsg); ok && key.String() == "enter" {
		if it, ok := a.list.SelectedItem().(artistsListItem); ok {
			return a, func() tea.Msg { return getArtistMsg{artistID: it.artist.ID} }
		}
	}

	var cmd tea.Cmd
	a.list, cmd = a.list.Update(msg)
	return a, cmd
}

func (a artistsListModel) View() string {
	return a.list.View()
}
