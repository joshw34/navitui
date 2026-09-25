package navidrome

import (
	"encoding/json"
	"fmt"
	"maps"
	"net/http"
	"net/url"

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
		c.logger.File("http request failed (req type: %v): %v", reqType, err)
		return nil, err
	}
	defer func() {
		_ = r.Body.Close()
	}()
	if r.StatusCode != http.StatusOK {
		c.logger.File("http error code: %v", r.StatusCode)
		return nil, fmt.Errorf("http error: %v/n", r.StatusCode)
	}

	var resp jsonResponse
	if err = json.NewDecoder(r.Body).Decode(&resp); err != nil {
		c.logger.File("Failed to decode json: %v", err)
		return nil, err
	}

	if resp.SubResp.Status == "failed" {
		c.logger.File("subsonic-response status failed, code: %v, message: %v", resp.SubResp.Error.Code, resp.SubResp.Error.Message)
		return nil, fmt.Errorf("subsonic error %v: %v",
			resp.SubResp.Error.Code, resp.SubResp.Error.Message)
	}

	return &resp, nil
}

func (c *Navidrome) GetStreamURL(songID string) (string, error) {
	v := url.Values{}
	v.Set("id", songID)
	return c.getURL("stream", v)
}

func (c *Navidrome) PingTest() types.PingResult {
	c.logger.File("Started PingTest")
	u, err := c.getURL("ping", nil)
	if err != nil {
		return types.PingURLBuildFailure
	}
	r, err := http.Get(u)
	if err != nil {
		c.logger.File("PingTest: http request failed: %v", err)
		return types.PingServerError
	}
	defer func() {
		_ = r.Body.Close()
	}()
	if r.StatusCode != http.StatusOK {
		c.logger.File("PingTest: http error code: %v", r.StatusCode)
		return types.PingServerError
	}
	var resp jsonResponse
	if err := json.NewDecoder(r.Body).Decode(&resp); err != nil {
		c.logger.File("PingTest: Failed to decode json: %v", err)
		return types.PingServerError
	}
	if resp.SubResp.Status == "failed" {
		c.logger.File("PingTest: subsonic-response status failed, code: %v, message: %v", resp.SubResp.Error.Code, resp.SubResp.Error.Message)
		return types.PingAuthFailure
	}
	return types.PingSuccess
}

func (c *Navidrome) getURL(req string, extraParams url.Values) (string, error) {
	uPath, err := url.JoinPath(c.baseURL, req)
	if err != nil {
		c.logger.File("Failed to join request %v to baseURL: %v, ", req, err)
		return "", err
	}
	u, err := url.Parse(uPath)
	if err != nil {
		c.logger.File("Failed to parse url %v: %v", uPath, err)
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
