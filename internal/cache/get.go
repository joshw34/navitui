package cache

import (
	"github.com/joshw34/navitui/internal/types"
)

func (c *Cache) GetArtists() ([]types.Artist, error) {
	query := `SELECT id, name
			  FROM artists
			  ORDER BY name;`
	rows, err := c.Data.Query(query)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = rows.Close()
	}()

	var result []types.Artist

	for rows.Next() {
		var a types.Artist
		err = rows.Scan(&a.ID, &a.Name)
		if err != nil {
			return nil, err
		}
		result = append(result, a)
	}

	return result, nil
}
