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
	return fmt.Sprintf("%v   %s", d, a.songs.ID)
}

func (a songsItem) FilterValue() string { return a.songs.Title }

func (a songsModel) buildList(data []types.Song) songsModel {
	items := make([]list.Item, len(data))
	for i, song := range data {
		items[i] = songsItem{songs: song}
	}
	a.list.SetItems(items)
	return a
}

type songsModel struct {
	list list.Model
}

func (a songsModel) Update(msg tea.Msg) (songsModel, tea.Cmd) {
	if key, ok := msg.(tea.KeyPressMsg); ok && key.String() == "enter" {
		if it, ok := a.list.SelectedItem().(songsItem); ok {
			return a, func() tea.Msg { return playSongMsg{songID: it.songs.ID} }
		}
	}

	var cmd tea.Cmd
	a.list, cmd = a.list.Update(msg)
	return a, cmd
}

func (a songsModel) View() string {
	return a.list.View()
}
