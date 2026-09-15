package cache

import (
	"github.com/joshw34/navitui/internal/types"
)

func (c *Cache) GetArtists() ([]types.Artist, error) {
	query := `SELECT id, name
			  FROM artists
			  ORDER BY name;`
	rows, err := c.Data.Query(query)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = rows.Close()
	}()

	var result []types.Artist

	for rows.Next() {
		var a types.Artist
		err = rows.Scan(&a.ID, &a.Name)
		if err != nil {
			return nil, err
		}
		result = append(result, a)
	}

	return result, nil
}

func (c *Cache) GetAlbumsByArtist(searchID string) ([]types.Album, error) {
	query := `
		SELECT id, artistId, name, year, duration, songCount
		FROM albums
		WHERE artistId = ?
		ORDER BY name;`

	rows, err := c.Data.Query(query, searchID)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = rows.Close()
	}()

	var result []types.Album

	for rows.Next() {
		var a types.Album
		err = rows.Scan(&a.ID, &a.ArtistID, &a.Name, &a.Year, &a.Duration, &a.SongCount)
		if err != nil {
			return nil, err
		}
		result = append(result, a)
	}

	return result, nil
}

func (c *Cache) GetSongsByAlbum(searchID string) ([]types.Song, error) {
	query := `
		SELECT id, artistId, albumId, title, track, year, duration, disc
		FROM songs
		WHERE albumId = ?
		ORDER BY
    		CASE WHEN track = 0 THEN 1 ELSE 0 END,
    		track,
    		title COLLATE NOCASE;`

	rows, err := c.Data.Query(query, searchID)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = rows.Close()
	}()

	var result []types.Song

	for rows.Next() {
		var s types.Song
		err = rows.Scan(&s.ID, &s.ArtistID, &s.AlbumID, &s.Title, &s.Track, &s.Year, &s.Duration, &s.Disc)
		if err != nil {
			return nil, err
		}
		result = append(result, s)
	}

	return result, nil
}
