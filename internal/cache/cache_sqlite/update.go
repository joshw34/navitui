package cache_sqlite

import (
	"database/sql"

	"github.com/joshw34/navitui/internal/types"
)

func (c *CacheSQLite) UpdateArtists(a []types.Artist, resync bool) error {
	var tx *sql.Tx
	var stmt *sql.Stmt
	var err error
	tx, err = c.Data.Begin()
	if err != nil {
		return err
	}
	if resync {
		_, err = tx.Exec("DELETE FROM artists;")
		if err != nil {
			_ = tx.Rollback()
			return err
		}
	}
	stmt, err = tx.Prepare(`
		INSERT INTO artists (id, name, albumCount)
		VALUES (?, ?, ?)
		ON CONFLICT(id) DO UPDATE SET
			name = excluded.name,
			albumCount = excluded.albumCount;
	`)
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	defer func() { _ = stmt.Close() }()
	for _, artist := range a {
		_, err := stmt.Exec(artist.ID, artist.Name, artist.AlbumCount)
		if err != nil {
			_ = tx.Rollback()
			return err
		}
	}
	return tx.Commit()
}

func (c *CacheSQLite) UpdateAlbums(a []types.Album, resync bool) error {
	var tx *sql.Tx
	var stmt *sql.Stmt
	var err error
	tx, err = c.Data.Begin()
	if err != nil {
		return err
	}
	if resync {
		_, err = tx.Exec("DELETE FROM albums;")
		if err != nil {
			_ = tx.Rollback()
			return err
		}
	}
	stmt, err = tx.Prepare(`
		INSERT INTO albums (id, artistId, name, artist, genres, year, duration, songCount)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
		ON CONFLICT(id) DO UPDATE SET
			artistId = excluded.artistId,
			name = excluded.name,
			artist = excluded.artist,
			genres = excluded.genres,
			year = excluded.year,
			duration = excluded.duration,
			songCount = excluded.songCount;
	`)
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	defer func() { _ = stmt.Close() }()
	for _, album := range a {
		genresJSON, err := stringSliceToJSON(album.Genres)
		if err != nil {
			_ = tx.Rollback()
			return err
		}
		_, err = stmt.Exec(album.ID, album.ArtistID, album.Name, album.Artist, genresJSON, album.Year, album.Duration, album.SongCount)
		if err != nil {
			_ = tx.Rollback()
			return err
		}
	}
	return tx.Commit()
}

func (c *CacheSQLite) UpdateSongs(s []types.Song, resync bool) error {
	var tx *sql.Tx
	var stmt *sql.Stmt
	var err error
	tx, err = c.Data.Begin()
	if err != nil {
		return err
	}
	if resync {
		_, err = tx.Exec("DELETE FROM songs;")
		if err != nil {
			_ = tx.Rollback()
			return err
		}
	}
	stmt, err = tx.Prepare(`
		INSERT INTO songs (id, artistId, albumId, artist, album, title, filetype, track, year, duration, disc)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
		ON CONFLICT(id) DO UPDATE SET
			artistId = excluded.artistId,
			albumId = excluded.albumId,
			artist = excluded.artist,
			album = excluded.album,
			title = excluded.title,
			filetype = excluded.filetype,
			track = excluded.track,
			year = excluded.year,
			duration = excluded.duration,
			disc = excluded.disc;
	`)
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	defer func() { _ = stmt.Close() }()
	for _, song := range s {
		_, err := stmt.Exec(song.ID, song.ArtistID, song.AlbumID, song.Artist, song.Album, song.Title, song.FileType, song.Track, song.Year, song.Duration, song.Disc)
		if err != nil {
			_ = tx.Rollback()
			return err
		}
	}
	return tx.Commit()
}

func (c *CacheSQLite) ResyncLibrary(artists []types.Artist, albums []types.Album, songs []types.Song) error {
	var err error
	err = c.UpdateArtists(artists, true)
	if err != nil {
		return err
	}
	err = c.UpdateAlbums(albums, true)
	if err != nil {
		return err
	}
	err = c.UpdateSongs(songs, true)
	if err != nil {
		return err
	}
	return nil
}
