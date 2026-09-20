package ui

import (
	"fmt"

	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
	"github.com/joshw34/navitui/internal/types"
)

type nowPlayingModel struct {
	song types.Song
}

func (n nowPlayingModel) Update(msg tea.Msg) (pageModel, tea.Cmd) {
	return n, nil
}

func (n nowPlayingModel) View() string {
	return fmt.Sprintf("Title: %s\tArtist: %s\tAlbum: %s", n.song.Title, n.song.ArtistID, n.song.AlbumID)
}

func (n nowPlayingModel) updateData(s types.Song) nowPlayingModel {
	n.song = s
	return n
}

func queueToListItems(songs []types.Song) []list.Item {
	items := make([]list.Item, len(songs))
	for i, s := range songs {
		items[i] = queueItem{songs: s}
	}
	return items
}
