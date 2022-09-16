package acs

import (
	"fmt"
	"net/http"
	"net/url"
)

// Client is an http client with some extra parameters for interacting with ServiceNow specifically
type Client struct {
	*http.Client

	UserAgent string

	Username string
	Password string

	BaseURL *url.URL

	Retries int

	Sleep    string
	LogLevel string
}

func (c *Client) setAuth(req *http.Request) error {
	req.SetBasicAuth(c.Username, c.Password)

	if len(c.Username) > 0 && len(c.Password) > 0 {
		req.Header.Set("User-Agent", c.UserAgent)
		req.SetBasicAuth(c.Username, c.Password)
	}
	return fmt.Errorf("must provide either a username and password for ACS")
}
