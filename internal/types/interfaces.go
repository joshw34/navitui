package types

import (
	"github.com/joshw34/navitui/internal/player"
)

type Server interface {
	PingTest(host string, port int) PingResult
	GetStreamURL(songID string) (string, error)
	GetArtists() ([]Artist, error)
	GetAlbumsByArtist(searchId string) ([]Album, error)
	GetAllAlbums() ([]Album, error)
	GetSongsByAlbum(searchId string) ([]Song, error)
	GetAllSongs() ([]Song, error)
}

type Cache interface {
	GetArtists() ([]Artist, error)
	GetAlbumsByArtist(searchID string) ([]Album, error)
	GetAllAlbums() ([]Album, error)
	GetSongsByAlbum(searchID string) ([]Song, error)
	GetAllSongs() ([]Song, error)
	UpdateArtists(a []Artist, resync bool) error
	UpdateAlbums(a []Album, resync bool) error
	UpdateSongs(s []Song, resync bool) error
	ResyncLibrary(artists []Artist, albums []Album, songs []Song) error
}

type Player interface {
	SetEventHandler(fn func(player.Event))
	Close() error
	PlaySong(reqId uint64, url string) error
	Stop() error
	TogglePause() error
	Quit() error
	Seek(t float64) error
}
