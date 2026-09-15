package client

import (
	"net/url"

	"github.com/joshw34/navitui/internal/types"
)

func (c *Client) GetAlbumsByArtist(searchId string) ([]types.Album, error) {
	v := url.Values{}
	v.Set("id", searchId)
	r, err := c.serverRequest("getArtist", v)
	if err != nil {
		return nil, err
	}
	return c.extractAlbumsFromArtist(r)
}

func (c *Client) extractAlbumsFromArtist(r *jsonResponse) ([]types.Album, error) {
	var result []types.Album
	data := r.SubResp.Artist.Albums

	for _, r := range data {
		var a types.Album
		a.ID = r.ID
		a.ArtistID = r.ArtistID
		a.Name = r.Name
		a.Year = r.Year
		a.Duration = r.Duration
		a.SongCount = r.SongCount
		result = append(result, a)
	}
	return result, nil
}
