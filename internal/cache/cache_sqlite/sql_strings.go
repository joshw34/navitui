package cache_sqlite

const getArtistsQuery = `
						SELECT id, name, albumCount
						FROM artists
						ORDER BY name;`

const getAlbumsByArtistQuery = `
						SELECT id, artistId, name, artist, genres, year, duration, songCount
						FROM albums
						WHERE artistId = ?
						ORDER BY
		    				CASE WHEN year = 0 THEN 1 ELSE 0 END,
		    				year,
		    				name COLLATE NOCASE;`

const getAllAlbumsQuery = `
						SELECT id, artistId, name, artist, genres, year, duration, songCount
						FROM albums
						ORDER BY name;`

const getSongsByAlbumQuery = `
						SELECT id, artistId, albumId, artist, album, title, filetype, track, year, duration, disc
						FROM songs
						WHERE albumId = ?
						ORDER BY
							CASE WHEN track = 0 THEN 1 ELSE 0 END,
							track,
							title COLLATE NOCASE;`

const getAllSongsQuery = `
						SELECT id, artistId, albumId, artist, album, title, filetype, track, year, duration, disc
						FROM songs
						ORDER BY title;`

const insertArtistsQuery = `
						INSERT INTO artists (id, name, albumCount)
						VALUES (?, ?, ?)
						ON CONFLICT(id) DO UPDATE SET
							name = excluded.name,
							albumCount = excluded.albumCount;`

const insertAlbumsQuery = `
						INSERT INTO albums (id, artistId, name, artist, genres, year, duration, songCount)
						VALUES (?, ?, ?, ?, ?, ?, ?, ?)
						ON CONFLICT(id) DO UPDATE SET
							artistId = excluded.artistId,
							name = excluded.name,
							artist = excluded.artist,
							genres = excluded.genres,
							year = excluded.year,
							duration = excluded.duration,
							songCount = excluded.songCount;`

const insertSongsQuery = `
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
							disc = excluded.disc;`

const deleteAllArtistsQuery = `DELETE FROM artists`

const deleteAllAlbumsQuery = `DELETE FROM albums`

const deleteAllSongsQuery = `DELETE FROM songs`
