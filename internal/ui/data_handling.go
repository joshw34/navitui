package ui

import (
	tea "charm.land/bubbletea/v2"
	"github.com/joshw34/navitui/internal/controller"
	"github.com/joshw34/navitui/internal/types"
)

// ARTIST MESSAGES
type getAllArtistsMsg struct{}

type loadedAllArtistsMsg struct {
	data []types.Artist
	err  error
}

// ALBUM MESSAGES
type getAlbumsByArtistMsg struct {
	artistID string
}

type getAllAlbumsMsg struct{}

type loadedAlbumsMsg struct {
	data []types.Album
	err  error
}

// SONG MESSAGES
type getSongsByAlbumMsg struct {
	albumID string
}

type getAllSongsMsg struct{}

type loadedSongsMsg struct {
	data []types.Song
	err  error
}

// MPV MESSAGES
type playSongMsg struct {
	songID string
}

// DATA RETRIEVAL
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
