// Package navidrome: implement types.MusicServer interface for navidrome servers
package navidrome

import (
	"net/url"

	"github.com/joshw34/navitui/internal/types"
)

var _ types.Server = (*Navidrome)(nil)

type Navidrome struct {
	baseURL  string
	username string
	password string
	version  string
	client   string
	format   string
}

func New(baseURL, username, password string) (*Navidrome, error) {
	var client Navidrome
	var err error
	client.baseURL, err = url.JoinPath(baseURL, "rest")
	if err != nil {
		return nil, err
	}
	client.username = username
	client.password = password
	client.version = "1.16.1"
	client.client = "navitui"
	client.format = "json"
	return &client, nil
}
