package internal

import (
	"net/http"
)

type Client interface {
	Do(req *http.Request) (*http.Response, error)
}

type client struct {
	HTTP Client
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

func WithHTTPClient(c Client) ClientOption {
	return func(opts *client) {
		opts.HTTP = c
	}
}

func NewClient(opts ...ClientOption) (Client, error) {
	c := &client{
		HTTP: http.DefaultClient,
	}

	c.withOpts(opts)

	return c, nil
}
