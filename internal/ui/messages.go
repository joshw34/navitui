package ui

import (
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

// PLAYER MESSAGES
type playSongMsg struct {
	s types.Song
}

type addToQueueMsg struct {
	s types.Song
}

type removeFromQueueMsg struct {
	index int
}

// UI UPDATE MESSAGES
type updateMsg struct {
	u controller.Update
}
