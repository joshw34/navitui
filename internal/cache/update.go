package cache

import "github.com/joshw34/navitui/internal/types"

func (c *Cache) UpdateArtists(a []types.Artist) error {
	query := `INSERT INTO artists (id, name)
			  VALUES (?, ?)
			  ON CONFLICT(id) DO UPDATE SET
			  	name = excluded.name;`

	for _, artist := range a {
		_, err := c.Data.Exec(query, artist.ID, artist.Name)
		if err != nil {
			return err
		}
	}
	return nil
}
