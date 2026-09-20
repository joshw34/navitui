package ui

import (
	"fmt"

	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
	"github.com/joshw34/navitui/internal/types"
)

type nowPlayingModel struct {
	song    types.Song
	timePos float64
}

func (n nowPlayingModel) Update(msg tea.Msg) (pageModel, tea.Cmd) {
	return n, nil
}

func (n nowPlayingModel) View() string {
	return fmt.Sprintf("Title: %s\tArtist: %s\tAlbum: %s\nTime: %f", n.song.Title, n.song.Artist, n.song.Album, n.timePos)
}

func (n nowPlayingModel) updateSong(s types.Song) nowPlayingModel {
	n.song = s
	return n
}

func (n nowPlayingModel) updateTP(tp float64) nowPlayingModel {
	n.timePos = tp
	return n
}

func queueToListItems(songs []types.Song) []list.Item {
	items := make([]list.Item, len(songs))
	for i, s := range songs {
		items[i] = queueItem{songs: s}
	}
	return items
}
