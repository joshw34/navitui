package navidrome

import (
	"encoding/json"
	"fmt"
	"maps"
	"net"
	"net/http"
	"net/url"
	"strconv"

	"github.com/joshw34/navitui/internal/types"
)

// TODO: Add http return code check to avoid json parsing error html
func (c *Navidrome) serverRequest(reqType string, extraParams url.Values) (*jsonResponse, error) {
	u, err := c.getURL(reqType, extraParams)
	if err != nil {
		return nil, err
	}

	r, err := http.Get(u)
	if err != nil {

		return nil, err
	}
	defer func() {
		_ = r.Body.Close()
	}()
	if r.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http error: %d/n", r.StatusCode)
	}

	var resp jsonResponse
	if err := json.NewDecoder(r.Body).Decode(&resp); err != nil {
		return nil, err
	}

	if resp.SubResp.Status == "failed" {
		return nil, fmt.Errorf("subsonic error %d: %s",
			resp.SubResp.Error.Code, resp.SubResp.Error.Message)
	}

	return &resp, nil
}

func (c *Navidrome) GetStreamURL(songID string) (string, error) {
	v := url.Values{}
	v.Set("id", songID)
	//v.Set("format", "raw")
	return c.getURL("stream", v)
}

func (c *Navidrome) PingTest(host string, port int) types.PingResult {
	baseURL, err := url.JoinPath(host, "rest", "ping")
	if err != nil {
		return types.URLBuildFailure
	}
	u, err := url.Parse(baseURL)
	if err != nil {
		return types.URLBuildFailure
	}
	u.Host = net.JoinHostPort(u.Host, strconv.Itoa(port))
	return types.AuthFailure
}

func (c *Navidrome) getURL(req string, extraParams url.Values) (string, error) {
	uPath, err := url.JoinPath(c.baseURL, req)
	if err != nil {
		return "", err
	}
	u, err := url.Parse(uPath)
	if err != nil {
		return "", err
	}

	q := url.Values{}
	q.Set("u", c.username)
	q.Set("p", c.password)
	q.Set("v", c.version)
	q.Set("c", c.client)
	q.Set("f", c.format)
	maps.Copy(q, extraParams)
	u.RawQuery = q.Encode()

	return u.String(), nil
}
