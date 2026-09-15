package client

import (
	"net/url"

	"github.com/joshw34/navitui/internal/types"
)

func (c *Client) GetSongsByAlbum(searchId string) ([]types.Song, error) {
	v := url.Values{}
	v.Set("id", searchId)
	r, err := c.serverRequest("getAlbum", v)
	if err != nil {
		return nil, err
	}
	return c.extractSongsFromAlbum(r)
}

func (c *Client) extractSongsFromAlbum(r *jsonResponse) ([]types.Song, error) {
	var result []types.Song
	data := r.SubResp.Album.Songs

	for _, r := range data {
		var s types.Song
		s.ID = r.ID
		s.ArtistID = r.ArtistID
		s.AlbumID = r.AlbumID
		s.Title = r.Title
		s.Track = r.Track
		s.Year = r.Year
		s.Duration = r.Duration
		s.Disc = r.Disc
		result = append(result, s)
	}
	return result, nil
}
