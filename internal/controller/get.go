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
		err = c.db.UpdateArtists(a, false)
		if err != nil {
			return nil, err
		}
		return c.db.GetArtists()
	}
	return a, nil
}

func (c *Controller) GetAllAlbums() ([]types.Album, error) {
	a, err := c.db.GetAllAlbums()
	if err != nil {
		return nil, err
	}
	if len(a) == 0 {
		a, err = c.srv.GetAllAlbums()
		if err != nil {
			return nil, err
		}
		err = c.db.UpdateAlbums(a, false)
		if err != nil {
			return nil, err
		}
		return c.db.GetAllAlbums()
	}
	return a, nil
}

func (c *Controller) GetAllSongs() ([]types.Song, error) {
	s, err := c.db.GetAllSongs()
	if err != nil {
		return nil, err
	}
	if len(s) == 0 {
		s, err = c.srv.GetAllSongs()
		if err != nil {
			return nil, err
		}
		err = c.db.UpdateSongs(s, false)
		if err != nil {
			return nil, err
		}
		return c.db.GetAllSongs()
	}
	return s, nil
}

func (c *Controller) ResyncLibrary() error {
	newArtists, err := c.srv.GetArtists()
	if err != nil {
		return err
	}
	err = c.db.UpdateArtists(newArtists, true)
	if err != nil {
		return err
	}
	newAlbums, err := c.srv.GetAllAlbums()
	if err != nil {
		return err
	}
	err = c.db.UpdateAlbums(newAlbums, true)
	if err != nil {
		return err
	}
	newSongs, err := c.srv.GetAllSongs()
	if err != nil {
		return err
	}
	err = c.db.UpdateSongs(newSongs, true)
	if err != nil {
		return err
	}
	return nil
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
		err = c.db.UpdateAlbums(a, false)
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
		err = c.db.UpdateSongs(s, false)
		if err != nil {
			return nil, err
		}
		return c.db.GetSongsByAlbum(searchId)
	}
	return s, nil
}
