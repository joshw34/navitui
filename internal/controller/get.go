package controller

import (
	"github.com/joshw34/navitui/internal/types"
)

func (c *Controller) GetArtists() ([]types.Artist, error) {
	// pull from db
	a, err := c.db.GetArtists()
	if err != nil {
		return nil, err
	}
	if len(a) == 0 {
		// pull from server
		a, err = c.srv.GetArtists()
		if err != nil {
			return nil, err
		}
		// update db
		err = c.db.UpdateArtists(a)
		if err != nil {
			return nil, err
		}
		return c.db.GetArtists()
	}
	return a, nil
}
