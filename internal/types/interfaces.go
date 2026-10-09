package types

import (
	"github.com/joshw34/navitui/internal/player"
)

type Server interface {
	PingTest() PingResult
	GetStreamURL(songID string) (string, error)
	GetArtists() ([]Artist, error)
	GetAlbumsByArtist(searchId string) ([]Album, error)
	GetAllAlbums() ([]Album, error)
	GetSongsByAlbum(searchId string) ([]Song, error)
	GetAllSongs() ([]Song, error)
}

type Cache interface {
	SyncRequired() bool
	Close()
	GetArtists() ([]Artist, error)
	GetAlbumsByArtist(searchID string) ([]Album, error)
	GetAllAlbums() ([]Album, error)
	GetSongsByAlbum(searchID string) ([]Song, error)
	GetAllSongs() ([]Song, error)
	UpdateCache(artists []Artist, albums []Album, songs []Song, resync bool) error
}

type Player interface {
	SetEventHandler(fn func(player.Event))
	Close()
	PlaySong(reqId uint64, url string) error
	Stop() error
	TogglePause() error
	Quit() error
	Seek(t float64) error
}
