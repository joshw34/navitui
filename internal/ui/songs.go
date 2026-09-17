package ui

import (
	"fmt"
	"strconv"
	"time"

	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
	"github.com/joshw34/navitui/internal/types"
)

type songsItem struct {
	songs types.Song
}

func (a songsItem) Title() string { return a.songs.Title }

func (a songsItem) Description() string {
	d, _ := time.ParseDuration(strconv.Itoa(a.songs.Duration) + "s")
	return fmt.Sprintf("%v", d)
}

func (a songsItem) FilterValue() string { return a.songs.Title }

func onSelectedSongs(it list.Item) tea.Cmd {
	return func() tea.Msg { return playSongMsg{it.(songsItem).songs.ID} }
}

func songsToListItems(songs []types.Song) []list.Item {
	items := make([]list.Item, len(songs))
	for i, s := range songs {
		items[i] = songsItem{songs: s}
	}
	return items
}
