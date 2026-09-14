package ui

import (
	tea "charm.land/bubbletea/v2"
	"github.com/joshw34/navitui/internal/controller"
	"github.com/joshw34/navitui/internal/types"
)

type getArtistsMsg struct{}
type artistsLoadedMsg struct {
	artists []types.Artist
	err     error
}

func loadArtists(ctrl *controller.Controller) tea.Cmd {
	return func() tea.Msg {
		artists, err := ctrl.GetArtists()
		return artistsLoadedMsg{artists: artists, err: err}
	}
}
