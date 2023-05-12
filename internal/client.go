package internal

import (
	"net/http"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type client struct {
	HTTP HTTPClient
}

func (c *client) Do(req *http.Request) (*http.Response, error) {
	return c.HTTP.Do(req)
}

type ClientOption func(opts *client)

type clientOptions struct {
	HTTP *http.Client
}

func (c *client) withOpts(opts []ClientOption) {
	// set default values
	c.HTTP = &http.Client{}

	for _, opt := range opts {
		opt(c)
	}
}

func WithHTTPClient(c HTTPClient) ClientOption {
	return func(opts *client) {
		opts.HTTP = c
	}
}

func NewClient(opts ...ClientOption) (HTTPClient, error) {
	c := &client{
		HTTP: http.DefaultClient,
	}

	c.withOpts(opts)

	return c, nil
}
