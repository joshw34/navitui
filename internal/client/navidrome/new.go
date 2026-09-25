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
	logger   *types.NavituiLogger
}

func New(baseURL, username, password string, logger *types.NavituiLogger) (*Navidrome, error) {
	var client Navidrome
	var err error
	client.logger = logger
	client.baseURL, err = url.JoinPath(baseURL, "rest")
	if err != nil {
		logger.File("Failed to build baseURL: %v", err)
		return nil, err
	}
	client.username = username
	client.password = password
	client.version = "1.16.1"
	client.client = "navitui"
	client.format = "json"
	return &client, nil
}
