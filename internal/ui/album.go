package ui

import (
	"fmt"
	"strconv"
	"time"

	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
	"github.com/joshw34/navitui/internal/types"
)

type albumItem struct {
	songs types.Song
}

func (a albumItem) Title() string { return a.songs.Title }

func (a albumItem) Description() string {
	d, _ := time.ParseDuration(strconv.Itoa(a.songs.Duration) + "s")
	return fmt.Sprintf("%v", d)
}

func (a albumItem) FilterValue() string { return a.songs.Title }

func (a albumModel) buildList(data []types.Song) albumModel {
	items := make([]list.Item, len(data))
	for i, song := range data {
		items[i] = albumItem{songs: song}
	}
	a.list.SetItems(items)
	return a
}

type albumModel struct {
	list list.Model
}

func (a albumModel) Update(msg tea.Msg) (albumModel, tea.Cmd) {
	/*if key, ok := msg.(tea.KeyPressMsg); ok && key.String() == "enter" {
		if it, ok := a.list.SelectedItem().(albumItem); ok {
			return a, func() tea.Msg { return getAlbumMsg{albumID: it.songs.ID} }
		}
	}*/

	var cmd tea.Cmd
	a.list, cmd = a.list.Update(msg)
	return a, cmd
}

func (a albumModel) View() string {
	return a.list.View()
}
