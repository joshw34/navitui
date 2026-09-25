package navidrome

import (
	"net/url"
	"strconv"

	"github.com/joshw34/navitui/internal/types"
)

func (c *Navidrome) GetSongsByAlbum(searchId string) ([]types.Song, error) {
	c.logger.File("Server Request: GetSongsByAlbum, AlbumId: %s", searchId)
	v := url.Values{}
	v.Set("id", searchId)
	r, err := c.serverRequest("getAlbum", v)
	if err != nil {
		return nil, err
	}
	c.logger.File("Server Request: GetSongsByAlbum, AlbumId: %s -> success", searchId)
	return c.extractSongs(r.SubResp.AlbumSongs.Songs), nil
}

func (c *Navidrome) GetAllSongs() ([]types.Song, error) {
	c.logger.File("Server Request: GetAllSongs")
	offset := 0
	size := 500
	var result []types.Song
	for {
		v := url.Values{}
		v.Set("query", " ")
		v.Set("artistCount", "0")
		v.Set("albumCount", "0")
		v.Set("songCount", strconv.Itoa(size))
		v.Set("songOffset", strconv.Itoa(offset))
		r, err := c.serverRequest("search2", v)
		if err != nil {
			return nil, err
		}
		if r.SubResp.Search2.Songs == nil {
			c.logger.File("Server Request: GetAllSongs -> success")
			return result, nil
		}
		extracted := c.extractSongs(r.SubResp.Search2.Songs)
		result = append(result, extracted...)
		offset += size
	}
}

func (c *Navidrome) extractSongs(data []jsonSong) []types.Song {
	var result []types.Song

	for _, r := range data {
		var s types.Song
		s.ID = r.ID
		s.ArtistID = r.ArtistID
		s.AlbumID = r.AlbumID
		s.Artist = r.Artist
		s.Album = r.Album
		s.Title = r.Title
		s.FileType = r.FileType
		s.Track = r.Track
		s.Year = r.Year
		s.Duration = r.Duration
		s.Disc = r.Disc
		result = append(result, s)
	}

	return result
}
