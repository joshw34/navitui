// Package cache: Database handler
package cache

import (
	"database/sql"
	"os"
	"path/filepath"

	"github.com/adrg/xdg"
	_ "modernc.org/sqlite"
)

type Cache struct {
	Data *sql.DB
}

func New() (*Cache, error) {
	cacheDir := filepath.Join(xdg.CacheHome, "navitui")
	err := os.MkdirAll(cacheDir, 0o700)
	if err != nil {
		return nil, err
	}

	dbFile := filepath.Join(cacheDir, "cache.db")
	dbPragmas := "?_pragma=foreign_keys=1"
	db, err := sql.Open("sqlite", dbFile+dbPragmas)
	if err != nil {
		return nil, err
	}

	var result Cache
	result.Data = db
	_, err = result.createTables()
	if err != nil {
		_ = db.Close()
		return nil, err
	}
	return &result, nil
}

func (c *Cache) createTables() (sql.Result, error) {
	schema := `
		CREATE TABLE IF NOT EXISTS artists (
			id TEXT PRIMARY KEY,
			name TEXT NOT NULL,
			albumCount INTEGER NOT NULL
		);

		CREATE TABLE IF NOT EXISTS albums (
    		 id TEXT PRIMARY KEY,
    		 artistId TEXT NOT NULL REFERENCES artists(id) ON DELETE CASCADE,
    		 name TEXT NOT NULL,
		     genres TEXT,
    		 year INTEGER NOT NULL DEFAULT 0,
    		 duration INTEGER NOT NULL,
    		 songCount INTEGER NOT NULL
		);

		CREATE INDEX IF NOT EXISTS idx_albums_artistId ON albums(artistId);
		
		CREATE TABLE IF NOT EXISTS songs (
		    id TEXT PRIMARY KEY,
		    artistId TEXT NOT NULL REFERENCES artists(id) ON DELETE CASCADE,
		    albumId TEXT NOT NULL REFERENCES albums(id) ON DELETE CASCADE,
			title TEXT NOT NULL,
		    filetype TEXT NOT NULL,
		    track INTEGER NOT NULL DEFAULT 0,
		    year INTEGER NOT NULL DEFAULT 0,
		    duration INTEGER NOT NULL,
		    disc INTEGER NOT NULL DEFAULT 1
		);

		CREATE INDEX IF NOT EXISTS idx_songs_artistId ON songs(artistId);`

	return c.Data.Exec(schema)
}
