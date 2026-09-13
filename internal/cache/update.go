package cache

import "github.com/joshw34/navitui/internal/types"

func (c *Cache) UpdateArtists(a []types.Artist) error {
	sql := `INSERT INTO artists (id, name)
								 VALUES (?, ?)
								 ON CONFLICT(id) DO UPDATE SET
								 	name = excluded.name;`

	for _, artist := range a {
		_, err := c.Data.Exec(sql, artist.ID, artist.Name)
		if err != nil {
			return err
		}
	}
	return nil
}
