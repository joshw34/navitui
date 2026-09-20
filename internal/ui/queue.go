package ui

import (
	"fmt"

	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
	"github.com/joshw34/navitui/internal/types"
)

type queueItem struct {
	songs types.Song
}

func (q queueItem) Title() string { return q.songs.Title }

func (q queueItem) Description() string {
	return fmt.Sprintf("%s", q.songs.Artist)
}

func (q queueItem) FilterValue() string { return q.songs.Title }

func onKeypressQueue(key tea.KeyPressMsg, it list.Item) tea.Cmd {
	_ = key
	_ = it
	return nil
}
