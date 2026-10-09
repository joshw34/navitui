package cache_sqlite

import (
	"database/sql"

	"github.com/joshw34/navitui/internal/types"
)

func (c *CacheSQLite) UpdateCache(artists []types.Artist, albums []types.Album, songs []types.Song, resync bool) error {
	c.logger.File("Cache Update: Resync: %v", resync)
	var tx *sql.Tx
	var err error

	if tx, err = c.data.Begin(); err != nil {
		c.logger.File("Error: Failed to begin transaction: %v", err)
		return err
	}

	defer func() { c.rollbackOnError(tx, err) }()

	if len(artists) != 0 {
		if err = c.updateTable(tx, resync, deleteAllArtistsQuery, insertArtistsQuery,
			func(stmt *sql.Stmt) error {
				return c.artistRows(stmt, artists)
			}); err != nil {
			return err
		}
	}

	if len(albums) != 0 {
		if err = c.updateTable(tx, resync, deleteAllAlbumsQuery, insertAlbumsQuery,
			func(stmt *sql.Stmt) error {
				return c.albumRows(stmt, albums)
			}); err != nil {
			return err
		}
	}

	if len(songs) != 0 {
		if err = c.updateTable(tx, resync, deleteAllSongsQuery, insertSongsQuery,
			func(stmt *sql.Stmt) error {
				return c.songRows(stmt, songs)
			}); err != nil {
			return err
		}
	}

	if err = tx.Commit(); err != nil {
		c.logger.File("Error: Failed to commit transaction: %v", err)
		return err
	}

	c.logger.File("Cache Update: Resync: %v -> success", resync)
	return nil
}

func (c *CacheSQLite) updateTable(tx *sql.Tx, resync bool, del string, insert string, createRows func(*sql.Stmt) error) error {
	if resync {
		if _, err := tx.Exec(del); err != nil {
			c.logger.File("Error: Failed to delete existing rows: %v", err)
			return err
		}
	}

	var stmt *sql.Stmt
	var err error
	if stmt, err = tx.Prepare(insert); err != nil {
		c.logger.File("Error: Failed to prepare statement: %v", err)
		return err
	}
	defer c.closeStmt(stmt)

	return createRows(stmt)
}

func (c *CacheSQLite) artistRows(stmt *sql.Stmt, artists []types.Artist) error {
	for _, artist := range artists {
		if _, err := stmt.Exec(artist.ID, artist.Name, artist.AlbumCount); err != nil {
			c.logger.File("Error: Failed to insert row: %v", err)
			return err
		}
	}
	return nil
}

func (c *CacheSQLite) albumRows(stmt *sql.Stmt, albums []types.Album) error {
	for _, album := range albums {
		var genresJSON []byte
		var err error
		if genresJSON, err = stringSliceToJSON(album.Genres); err != nil {
			c.logger.File("Error: Failed to parse genres: %v", err)
			return err
		}
		if _, err = stmt.Exec(album.ID, album.ArtistID, album.Name, album.Artist, genresJSON, album.Year, album.Duration, album.SongCount); err != nil {
			c.logger.File("Error: Failed to insert row: %v", err)
			return err
		}
	}
	return nil
}

func (c *CacheSQLite) songRows(stmt *sql.Stmt, songs []types.Song) error {
	for _, song := range songs {
		if _, err := stmt.Exec(song.ID, song.ArtistID, song.AlbumID, song.Artist, song.Album, song.Title, song.FileType, song.Track, song.Year, song.Duration, song.Disc); err != nil {
			c.logger.File("Error: Failed to insert row: %v", err)
			return err
		}
	}
	return nil
}

func (c *CacheSQLite) rollbackOnError(tx *sql.Tx, err error) {
	if tx != nil && err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			c.logger.File("Error: Failed to rollback transaction: %v", rbErr)
		}
	}
}

func (c *CacheSQLite) closeStmt(stmt *sql.Stmt) {
	if stmt != nil {
		if err := stmt.Close(); err != nil {
			c.logger.File("Error: Failed to close statement: %v", err)
		}
	}
}
