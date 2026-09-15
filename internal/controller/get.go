package controller

import (
	"github.com/joshw34/navitui/internal/types"
)

func (c *Controller) GetArtists() ([]types.Artist, error) {
	// pull from db
	a, err := c.db.GetArtists()
	if err != nil {
		return nil, err
	}
	if len(a) == 0 {
		// pull from server
		a, err = c.srv.GetArtists()
		if err != nil {
			return nil, err
		}
		// update db
		err = c.db.UpdateArtists(a)
		if err != nil {
			return nil, err
		}
		return c.db.GetArtists()
	}
	return a, nil
}

func (c *Controller) GetAlbumsByArtist(searchId string) ([]types.Album, error) {
	a, err := c.db.GetAlbumsByArtist(searchId)
	if err != nil {
		return nil, err
	}
	if len(a) == 0 {
		a, err = c.srv.GetAlbumsByArtist(searchId)
		if err != nil {
			return nil, err
		}
		err = c.db.UpdateAlbums(a)
		if err != nil {
			return nil, err
		}
		return c.db.GetAlbumsByArtist(searchId)
	}
	return a, nil
}

func (c *Controller) GetSongsByAlbum(searchId string) ([]types.Song, error) {
	s, err := c.db.GetSongsByAlbum(searchId)
	if err != nil {
		return nil, err
	}
	if len(s) == 0 {
		s, err = c.srv.GetSongsByAlbum(searchId)
		if err != nil {
			return nil, err
		}
		err = c.db.UpdateSongs(s)
		if err != nil {
			return nil, err
		}
		return c.db.GetSongsByAlbum(searchId)
	}
	return s, nil
}
