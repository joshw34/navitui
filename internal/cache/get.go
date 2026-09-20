package cache

import (
	"github.com/joshw34/navitui/internal/types"
)

func (c *Cache) GetArtists() ([]types.Artist, error) {
	query := `SELECT id, name, albumCount
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
		err = rows.Scan(&a.ID, &a.Name, &a.AlbumCount)
		if err != nil {
			return nil, err
		}
		result = append(result, a)
	}

	return result, nil
}

func (c *Cache) GetAlbumsByArtist(searchID string) ([]types.Album, error) {
	query := `
		SELECT id, artistId, name, artist, genres, year, duration, songCount
		FROM albums
		WHERE artistId = ?
		ORDER BY
		    CASE WHEN year = 0 THEN 1 ELSE 0 END,
		    year,
		    name COLLATE NOCASE;`

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
		var genresJSON []byte
		if err := rows.Scan(&a.ID, &a.ArtistID, &a.Name, &a.Artist, &genresJSON, &a.Year, &a.Duration, &a.SongCount); err != nil {
			return nil, err
		}
		genresSlice, err := jsonToStringSlice(genresJSON)
		if err != nil {
			return nil, err
		}
		a.Genres = genresSlice
		result = append(result, a)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Cache) GetAllAlbums() ([]types.Album, error) {
	query := `
		SELECT id, artistId, name, artist, genres, year, duration, songCount
		FROM albums
		ORDER BY name;`

	rows, err := c.Data.Query(query)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()

	var result []types.Album

	for rows.Next() {
		var a types.Album
		var genresJSON []byte
		if err := rows.Scan(&a.ID, &a.ArtistID, &a.Name, &a.Artist, &genresJSON, &a.Year, &a.Duration, &a.SongCount); err != nil {
			return nil, err
		}
		genresSlice, err := jsonToStringSlice(genresJSON)
		if err != nil {
			return nil, err
		}
		a.Genres = genresSlice
		result = append(result, a)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Cache) GetSongsByAlbum(searchID string) ([]types.Song, error) {
	query := `
		SELECT id, artistId, albumId, artist, album, title, filetype, track, year, duration, disc
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
		err = rows.Scan(&s.ID, &s.ArtistID, &s.AlbumID, &s.Artist, &s.Album, &s.Title, &s.FileType, &s.Track, &s.Year, &s.Duration, &s.Disc)
		if err != nil {
			return nil, err
		}
		result = append(result, s)
	}

	return result, nil
}

func (c *Cache) GetAllSongs() ([]types.Song, error) {
	query := `
		SELECT id, artistId, albumId, artist, album, title, filetype, track, year, duration, disc
		FROM songs
		ORDER BY title;`

	rows, err := c.Data.Query(query)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = rows.Close()
	}()

	var result []types.Song

	for rows.Next() {
		var s types.Song
		err = rows.Scan(&s.ID, &s.ArtistID, &s.AlbumID, &s.Artist, &s.Album, &s.Title, &s.FileType, &s.Track, &s.Year, &s.Duration, &s.Disc)
		if err != nil {
			return nil, err
		}
		result = append(result, s)
	}

	return result, nil
}
