package client

import "github.com/joshw34/navitui/internal/types"

type subRespArtist struct {
	Indexes []index `json:"index"`
}

type index struct {
	Artists []artist `json:"artist"`
}

type artist struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// TODO: use getIndexes instead

func (c *Client) GetArtists() ([]types.Artist, error) {
	r, err := c.serverRequest("getArtists", nil)
	if err != nil {
		return nil, err
	}
	return c.extractArtists(r), nil
}

func (c *Client) extractArtists(r *response) []types.Artist {
	var result []types.Artist
	indexes := r.SubResp.SRArtist.Indexes

	for _, index := range indexes {
		for _, artist := range index.Artists {
			var a types.Artist
			a.ID = artist.ID
			a.Name = artist.Name
			result = append(result, a)
		}
	}

	return result
}
