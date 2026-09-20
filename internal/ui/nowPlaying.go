package ui

import (
	"fmt"
	"log"

	"charm.land/bubbles/v2/list"
	"charm.land/bubbles/v2/progress"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/joshw34/navitui/internal/types"
)

var (
	helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#626262")).Render
	yellow    = lipgloss.Color("#FDFF8C")
	pink      = lipgloss.Color("#FF7CCB")
)

type nowPlayingModel struct {
	song       types.Song
	timePos    float64
	percentage float64
	prog       progress.Model
}

func (n nowPlayingModel) Update(msg tea.Msg) (pageModel, tea.Cmd) {
	return n, nil
}

func (n nowPlayingModel) View() string {
	//t, _ := time.ParseDuration(strconv.Itoa(n.timePos) + "s")
	//return fmt.Sprintf("Title: %s\tArtist: %s\tAlbum: %s\nTime: %v", n.song.Title, n.song.Artist, n.song.Album, t)
	return fmt.Sprintf("Track: %s\nAlbum: %s\nArtist: %s\n%s\n", n.song.Title, n.song.Album, n.song.Artist, n.prog.ViewAs(n.percentage))
}

func (n nowPlayingModel) updateSong(s types.Song) nowPlayingModel {
	n.song = s
	return n
}

func (n nowPlayingModel) updateTP(tp float64) nowPlayingModel {
	n.timePos = tp
	n.percentage = tp / float64(n.song.Duration)
	log.Printf("PERCENT: %f", n.percentage)
	return n
}

func queueToListItems(songs []types.Song) []list.Item {
	items := make([]list.Item, len(songs))
	for i, s := range songs {
		items[i] = queueItem{songs: s}
	}
	return items
}
