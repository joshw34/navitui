package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func (c *Client) serverRequest(reqType string, extraParams url.Values) (*response, error) {
	u, err := c.getURL(reqType, extraParams)
	if err != nil {
		return nil, err
	}

	r, err := http.Get(u)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	var resp response
	if err := json.NewDecoder(r.Body).Decode(&resp); err != nil {
		return nil, err
	}

	if resp.SubResp.Status == "failed" {
		return nil, fmt.Errorf("subsonic error %d: %s",
			resp.SubResp.Error.Code, resp.SubResp.Error.Message)
	}

	return &resp, nil
}

func (c *Client) getURL(req string, extraParams url.Values) (string, error) {
	u, err := url.Parse(c.baseURL + req)
	if err != nil {
		return "", err
	}

	q := url.Values{}
	q.Set("u", c.username)
	q.Set("p", c.password)
	q.Set("v", c.version)
	q.Set("c", c.client)
	q.Set("f", c.format)

	for key, value := range extraParams {
		q[key] = value
	}
	u.RawQuery = q.Encode()
	return u.String(), nil
}
