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

func onKeypressSongs(key tea.KeyPressMsg, it list.Item) tea.Cmd {
	s := it.(songsItem)
	switch key.String() {
	case "enter":
		return func() tea.Msg { return playSongMsg{s.songs} }
	case "a":
		return func() tea.Msg { return addToQueueMsg{s.songs} }
	}
	return nil
}

func songsToListItems(songs []types.Song) []list.Item {
	items := make([]list.Item, len(songs))
	for i, s := range songs {
		items[i] = songsItem{songs: s}
	}
	return items
}
