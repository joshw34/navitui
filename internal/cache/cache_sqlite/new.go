// Package cache_sqlite: implement the types.Cache interface for sqlite
package cache_sqlite

import (
	"database/sql"
	"os"
	"path/filepath"

	"github.com/adrg/xdg"
	"github.com/joshw34/navitui/internal/types"
	_ "modernc.org/sqlite"
)

var _ types.Cache = (*CacheSQLite)(nil)

type CacheSQLite struct {
	data          *sql.DB
	syncOnStartup bool
	logger        *types.NavituiLogger
}

func (c *CacheSQLite) Close() {
	_ = c.data.Close()
}

func New(logger *types.NavituiLogger) (*CacheSQLite, error) {
	var c CacheSQLite
	c.logger = logger

	cacheDir := filepath.Join(xdg.CacheHome, "navitui")
	err := os.MkdirAll(cacheDir, 0o700)
	if err != nil {
		c.logger.File("failed to create the cache directory: %v", err)
		return nil, err
	}

	dbFile := filepath.Join(cacheDir, "cache.db")
	_, err = os.Stat(dbFile)
	newDB := os.IsNotExist(err)

	dbPragmas := "?_pragma=foreign_keys=1"
	db, err := sql.Open("sqlite", dbFile+dbPragmas)
	if err != nil {
		c.logger.File("failed to open the database: %v", err)
		return nil, err
	}
	if err := db.Ping(); err != nil {
		c.logger.File("failed to ping database: %v", err)
		return nil, err
	}

	c.data = db
	c.syncOnStartup = newDB
	_, err = c.createTables()
	if err != nil {
		_ = db.Close()
		c.logger.File("failed to create the database tables: %v", err)
		return nil, err
	}
	return &c, nil
}

func (c *CacheSQLite) createTables() (sql.Result, error) {
	schema := `
		CREATE TABLE IF NOT EXISTS artists (
			id TEXT PRIMARY KEY,
			name TEXT NOT NULL,
			albumCount INTEGER NOT NULL
		);

		CREATE TABLE IF NOT EXISTS albums (
    		 id TEXT PRIMARY KEY,
    		 artistId TEXT NOT NULL,
    		 name TEXT NOT NULL,
    		 artist TEXT NOT NULL,
		     genres TEXT,
    		 year INTEGER NOT NULL DEFAULT 0,
    		 duration INTEGER NOT NULL,
    		 songCount INTEGER NOT NULL
		);

		CREATE INDEX IF NOT EXISTS idx_albums_artistId ON albums(artistId);
		CREATE INDEX IF NOT EXISTS idx_albums_name ON albums(name COLLATE NOCASE);
		
		CREATE TABLE IF NOT EXISTS songs (
		    id TEXT PRIMARY KEY,
		    artistId TEXT NOT NULL,
		    albumId TEXT NOT NULL,
		    artist TEXT NOT NULL,
		    album TEXT NOT NULL,
			title TEXT NOT NULL,
		    filetype TEXT NOT NULL,
		    track INTEGER NOT NULL DEFAULT 0,
		    year INTEGER NOT NULL DEFAULT 0,
		    duration INTEGER NOT NULL,
		    disc INTEGER NOT NULL DEFAULT 1
		);

		CREATE INDEX IF NOT EXISTS idx_songs_artistId ON songs(artistId);
		CREATE INDEX IF NOT EXISTS idx_songs_title ON songs(title COLLATE NOCASE);
		CREATE INDEX IF NOT EXISTS idx_songs_album_track ON songs(albumId, disc, track);`

	return c.data.Exec(schema)
}

func (c *CacheSQLite) SyncRequired() bool {
	return c.syncOnStartup
}
