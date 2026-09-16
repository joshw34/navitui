package cache

import (
	"github.com/joshw34/navitui/internal/types"
)

// TODO: Add upsert behaviour

func (c *Cache) UpdateArtists(a []types.Artist) error {
	query := `INSERT INTO artists (id, name, albumCount)
			  VALUES (?, ?, ?)
			  ON CONFLICT(id) DO UPDATE SET
			  	name = excluded.name,
				albumCount = excluded.albumCount;`

	for _, artist := range a {
		_, err := c.Data.Exec(query, artist.ID, artist.Name, artist.AlbumCount)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Cache) UpdateAlbums(a []types.Album) error {
	query := `
		INSERT INTO albums (id, artistId, name, genres, year, duration, songCount)
		VALUES (?, ?, ?, ?, ?, ?, ?)
		ON CONFLICT(id) DO UPDATE SET
			artistId = excluded.artistId,
			name = excluded.name,
			genres = excluded.genres,
			year = excluded.year,
			duration = excluded.duration,
			songCount = excluded.songCount;`

	for _, album := range a {
		genresJSON, err := stringSliceToJSON(album.Genres)
		if err != nil {
			return err
		}
		_, err = c.Data.Exec(query, album.ID, album.ArtistID, album.Name, genresJSON, album.Year, album.Duration, album.SongCount)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Cache) UpdateSongs(s []types.Song) error {
	query := `
		INSERT INTO songs (id, artistId, albumId, title, filetype, track, year, duration, disc)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
		ON CONFLICT(id) DO UPDATE SET
			artistId = excluded.artistId,
			albumId = excluded.albumId,
			title = excluded.title,
			filetype = excluded.filetype,
			track = excluded.track,
			year = excluded.year,
			duration = excluded.duration,
			disc = excluded.disc;`
	for _, song := range s {
		_, err := c.Data.Exec(query, song.ID, song.ArtistID, song.AlbumID, song.Title, song.FileType, song.Track, song.Year, song.Duration, song.Disc)
		if err != nil {
			return err
		}
	}
	return nil
}
