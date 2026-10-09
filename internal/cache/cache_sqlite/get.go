package cache_sqlite

import (
	"database/sql"

	"github.com/joshw34/navitui/internal/types"
)

func (c *CacheSQLite) GetArtists() ([]types.Artist, error) {
	c.logger.File("Cache Query: GetArtists")
	var rows *sql.Rows
	var err error

	if rows, err = c.data.Query(getArtistsQuery); err != nil {
		c.logger.File("Error: Query Failed: %v", err)
		return nil, err
	}
	defer c.closeRows(rows)

	var result []types.Artist
	if result, err = c.extractArtists(rows); err != nil {
		return nil, err
	}

	c.logger.File("Cache Query: GetArtists --> Success")
	return result, nil
}

func (c *CacheSQLite) GetAlbumsByArtist(searchID string) ([]types.Album, error) {
	c.logger.File("Cache Query: GetAlbumsByArtist: ArtistID: %v", searchID)

	var rows *sql.Rows
	var err error

	if rows, err = c.data.Query(getAlbumsByArtistQuery, searchID); err != nil {
		c.logger.File("Error: Query Failed: %v", err)
		return nil, err
	}
	defer c.closeRows(rows)

	var result []types.Album
	if result, err = c.extractAlbums(rows); err != nil {
		return nil, err
	}

	c.logger.File("Cache Query: GetAlbumsByArtist: ArtistID: %v --> Success", searchID)
	return result, nil
}

func (c *CacheSQLite) GetAllAlbums() ([]types.Album, error) {
	c.logger.File("Cache Query: GetAllAlbums")
	var rows *sql.Rows
	var err error

	if rows, err = c.data.Query(getAllAlbumsQuery); err != nil {
		c.logger.File("Error: Query Failed: %v", err)
		return nil, err
	}
	defer c.closeRows(rows)

	var result []types.Album
	if result, err = c.extractAlbums(rows); err != nil {
		return nil, err
	}

	c.logger.File("Cache Query: GetAllAlbums --> Success")
	return result, nil
}

func (c *CacheSQLite) GetSongsByAlbum(searchID string) ([]types.Song, error) {
	c.logger.File("Cache Query: GetSongsByAlbum: AlbumID: %v", searchID)
	var rows *sql.Rows
	var err error
	if rows, err = c.data.Query(getSongsByAlbumQuery, searchID); err != nil {
		c.logger.File("Error: Query Failed: %v", err)
		return nil, err
	}
	defer c.closeRows(rows)

	var result []types.Song
	if result, err = c.extractSongs(rows); err != nil {
		return nil, err
	}

	c.logger.File("Cache Query: GetSongsByAlbum: AlbumID: %v --> Success", searchID)
	return result, nil
}

func (c *CacheSQLite) GetAllSongs() ([]types.Song, error) {
	c.logger.File("Cache Query: GetAllSongs")
	var rows *sql.Rows
	var err error
	if rows, err = c.data.Query(getAllSongsQuery); err != nil {
		return nil, err
	}
	defer c.closeRows(rows)

	var result []types.Song
	if result, err = c.extractSongs(rows); err != nil {
		return nil, err
	}

	c.logger.File("Cache Query: GetAllSongs --> Success")
	return result, nil
}

func (c *CacheSQLite) extractArtists(rows *sql.Rows) ([]types.Artist, error) {
	var result []types.Artist

	for rows.Next() {
		var a types.Artist
		if err := rows.Scan(&a.ID, &a.Name, &a.AlbumCount); err != nil {
			c.logger.File("Error: Scan Failed: %v", err)
			return nil, err
		}
		result = append(result, a)
	}

	if err := rows.Err(); err != nil {
		c.logger.File("Error: Row Iteration Failed: %v", err)
		return nil, err
	}

	return result, nil
}

func (c *CacheSQLite) extractAlbums(rows *sql.Rows) ([]types.Album, error) {
	var result []types.Album
	var err error

	for rows.Next() {
		var a types.Album
		var genresJSON []byte
		if err = rows.Scan(&a.ID, &a.ArtistID, &a.Name, &a.Artist, &genresJSON, &a.Year, &a.Duration, &a.SongCount); err != nil {
			c.logger.File("Error: Scan Failed: %v", err)
			return nil, err
		}
		var genresSlice []string
		if genresSlice, err = jsonToStringSlice(genresJSON); err != nil {
			c.logger.File("Error: Failed to parse genres: %v", err)
			return nil, err
		}
		a.Genres = genresSlice
		result = append(result, a)
	}

	if err = rows.Err(); err != nil {
		c.logger.File("Error: Row Iteration Failed: %v", err)
		return nil, err
	}

	return result, nil
}

func (c *CacheSQLite) extractSongs(rows *sql.Rows) ([]types.Song, error) {
	var result []types.Song
	var err error

	for rows.Next() {
		var s types.Song
		err = rows.Scan(&s.ID, &s.ArtistID, &s.AlbumID, &s.Artist, &s.Album, &s.Title, &s.FileType, &s.Track, &s.Year, &s.Duration, &s.Disc)
		if err != nil {
			c.logger.File("Error: Scan Failed: %v", err)
			return nil, err
		}
		result = append(result, s)
	}

	if err = rows.Err(); err != nil {
		c.logger.File("Error: Row Iteration Failed: %v", err)
		return nil, err
	}

	return result, nil
}

func (c *CacheSQLite) closeRows(rows *sql.Rows) {
	err := rows.Close()
	if err != nil {
		c.logger.File("Error: Failed to close rows: %v", err)
	}
}
