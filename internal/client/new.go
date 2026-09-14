// Package client: communicate with the server
package client

type Client struct {
	baseURL  string
	username string
	password string
	version  string
	client   string
	format   string
}

func New(baseURL, username, password string) *Client {
	var client Client
	client.baseURL = baseURL
	client.username = username
	client.password = password
	client.version = "1.16.1"
	client.client = "navitui"
	client.format = "json"
	return &client
}
