package ui

import (
	tea "charm.land/bubbletea/v2"
	"github.com/joshw34/navitui/internal/controller"
)

func loadAllArtists(ctrl *controller.Controller) tea.Cmd {
	return func() tea.Msg {
		data, err := ctrl.GetArtists()
		return loadedAllArtistsMsg{data: data, err: err}
	}
}

func loadAllAlbums(ctrl *controller.Controller) tea.Cmd {
	return func() tea.Msg {
		data, err := ctrl.GetAllAlbums()
		return loadedAlbumsMsg{data: data, err: err}
	}
}

func loadAllSongs(ctrl *controller.Controller) tea.Cmd {
	return func() tea.Msg {
		data, err := ctrl.GetAllSongs()
		return loadedSongsMsg{data: data, err: err}
	}
}

func loadAlbumsByArtist(ctrl *controller.Controller, artistID string) tea.Cmd {
	return func() tea.Msg {
		data, err := ctrl.GetAlbumsByArtist(artistID)
		return loadedAlbumsMsg{data: data, err: err}
	}
}

func loadSongsByAlbum(ctrl *controller.Controller, albumID string) tea.Cmd {
	return func() tea.Msg {
		data, err := ctrl.GetSongsByAlbum(albumID)
		return loadedSongsMsg{data: data, err: err}
	}
}
