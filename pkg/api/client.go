package api

import (
	"net/http"
)

type Config struct {
	APIKey string
}

type Client struct {
	cli    *http.Client
	config *Config
}

func New(config *Config) *Client {
	return &Client{
		cli:    &http.Client{},
		config: config,
	}
}

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	if c.config.APIKey != "" {
		req.Header.Add("Authorization", "Bearer "+c.config.APIKey)
	}

	resp, err := c.cli.Do(req)
	return resp, err
}
