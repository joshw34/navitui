package navidrome

import (
	"net/url"
	"strconv"

	"github.com/joshw34/navitui/internal/types"
)

func (c *Navidrome) GetAlbumsByArtist(searchId string) ([]types.Album, error) {
	c.logger.File("Server Request: GetAlbumsByArtist, ArtistId: %s", searchId)
	v := url.Values{}
	v.Set("id", searchId)
	r, err := c.serverRequest("getArtist", v)
	if err != nil {
		return nil, err
	}
	c.logger.File("Server Request: GetAlbumsByArtist, ArtistId: %s -> success", searchId)
	return c.extractAlbums(r.SubResp.ArtistAlbums.Albums), nil
}

func (c *Navidrome) GetAllAlbums() ([]types.Album, error) {
	c.logger.File("Server Request: GetAllAlbums")
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
			c.logger.File("Server Request: GetAllAlbums -> success")
			return result, nil
		}
		extracted := c.extractAlbums(r.SubResp.AllAlbums.Albums)
		result = append(result, extracted...)
		offset += size
	}
}

func (c *Navidrome) extractAlbums(data []jsonAlbum) []types.Album {
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
	return result
}
