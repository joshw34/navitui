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

func Init() (*Cache, error) {
	cacheDir := filepath.Join(xdg.CacheHome, "navitui")
	err := os.MkdirAll(cacheDir, 0o700)
	if err != nil {
		return nil, err
	}

	dbFile := filepath.Join(cacheDir, "cache.db")
	dbPragmas := "?_pragma=foreign_keys(1)"
	db, err := sql.Open("sqlite", filepath.Join(dbFile, dbPragmas))
	if err != nil {
		return nil, err
	}

	var result Cache
	result.Data = db
	_, err = result.createTables()
	if err != nil {
		db.Close()
		return nil, err
	}
	return &result, nil
}

func (c *Cache) createTables() (sql.Result, error) {
	sql := `CREATE TABLE IF NOT EXISTS artists (
		id STRING PRIMARY KEY,
		name TEXT NOT NULL
	);`
	return c.Data.Exec(sql)
}
