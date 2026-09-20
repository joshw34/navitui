package client

import (
	"net/url"
	"strconv"

	"github.com/joshw34/navitui/internal/types"
)

func (c *Client) GetAlbumsByArtist(searchId string) ([]types.Album, error) {
	v := url.Values{}
	v.Set("id", searchId)
	r, err := c.serverRequest("getArtist", v)
	if err != nil {
		return nil, err
	}
	return c.extractAlbums(r.SubResp.ArtistAlbums.Albums)
}

func (c *Client) GetAllAlbums() ([]types.Album, error) {
	offset := 0
	size := 500
	var result []types.Album
	for {
		v := url.Values{}
		v.Set("offset", strconv.Itoa(offset))
		v.Set("size", strconv.Itoa(size))
		v.Set("type", "alphabeticalByName")
		r, err := c.serverRequest("getAlbumList", v)
		if err != nil {
			return nil, err
		}
		if r.SubResp.AllAlbums.Albums == nil {
			return result, nil
		}
		extracted, err := c.extractAlbums(r.SubResp.AllAlbums.Albums)
		if err != nil {
			return nil, err
		}
		result = append(result, extracted...)
		offset += size
	}
}

func (c *Client) extractAlbums(data []jsonAlbum) ([]types.Album, error) {
	var result []types.Album

	for _, r := range data {
		var a types.Album
		a.ID = r.ID
		a.ArtistID = r.ArtistID
		a.Name = r.Name
		a.Artist = r.Artist
		for _, g := range r.Genres {
			a.Genres = append(a.Genres, g.Name)
		}
		a.Year = r.Year
		a.Duration = r.Duration
		a.SongCount = r.SongCount
		result = append(result, a)
	}
	return result, nil
}
