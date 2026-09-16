package ui

import (
	tea "charm.land/bubbletea/v2"
	"github.com/joshw34/navitui/internal/controller"
	"github.com/joshw34/navitui/internal/types"
)

type getArtistsListMsg struct{}
type artistsListLoadedMsg struct {
	data []types.Artist
	err  error
}

type getArtistMsg struct {
	artistID string
}

type artistLoadedMsg struct {
	data []types.Album
	err  error
}

type getAlbumMsg struct {
	albumID string
}

type albumLoadedMsg struct {
	data []types.Song
	err  error
}

type playSongMsg struct {
	songID string
}

func loadArtistsList(ctrl *controller.Controller) tea.Cmd {
	return func() tea.Msg {
		data, err := ctrl.GetArtists()
		return artistsListLoadedMsg{data: data, err: err}
	}
}

func loadArtist(ctrl *controller.Controller, artistID string) tea.Cmd {
	return func() tea.Msg {
		data, err := ctrl.GetAlbumsByArtist(artistID)
		return artistLoadedMsg{data: data, err: err}
	}
}

func loadAlbum(ctrl *controller.Controller, albumID string) tea.Cmd {
	return func() tea.Msg {
		data, err := ctrl.GetSongsByAlbum(albumID)
		return albumLoadedMsg{data: data, err: err}
	}
}
