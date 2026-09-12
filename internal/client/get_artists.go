package client

type subRespArtist struct {
	Indexes []index `json:"index"`
}

type index struct {
	Artists []Artist `json:"artist"`
}

type Artist struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

// TODO: use getIndexes instead

func (c *Client) GetArtists() ([]Artist, error) {
	r, err := c.serverRequest("getArtists", nil)
	if err != nil {
		return nil, err
	}
	return c.extractArtists(r), nil
}

func (c *Client) extractArtists(r *response) []Artist {
	var result []Artist
	indexes := r.SubResp.SRArtist.Indexes

	for _, index := range indexes {
		result = append(result, index.Artists...)
	}

	return result
}
