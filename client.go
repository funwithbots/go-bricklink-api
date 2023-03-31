package go_bricklink_api

import (
	"context"
	"net/http"

	"github.com/funwithbots/go-bricklink-api/util"
)

const (
	apiURL               = "https://api.bricklink.com/api/store/v1"
	oauthSignatureMethod = "HMAC-SHA1"
	oauthVersion         = "1.0"
)

type Client interface {
	Orders() OrderClient
}

type client struct {
	HTTP        *http.Client
	ConsumerKey string
	OathToken   string
}

func (c *client) makeRequest(ctx context.Context, method string, opts ...options) (*http.Request, error) {
	return nil, util.ErrNotImplemented
}

func (c *client) sendRequest(ctx context.Context, req *http.Request) (*http.Response, error) {
	return nil, util.ErrNotImplemented
}

type options interface {
	withOpts(opts []func(opts any) any)
}

type ClientOption func(opts clientOptions) clientOptions

type clientOptions struct {
	HTTP        *http.Client
	ConsumerKey string
	OathToken   string
}

func (co *clientOptions) withOpts(opts []func(opts clientOptions) clientOptions) {
	for _, opt := range opts {
		*co = opt(*co)
	}
}

func WithHTTPClient(http *http.Client) ClientOption {
	return func(opts clientOptions) clientOptions {
		opts.HTTP = http
		return opts
	}
}

func WithConsumerKey(consumerKey string) ClientOption {
	return func(opts clientOptions) clientOptions {
		opts.ConsumerKey = consumerKey
		return opts
	}
}

func WithOathToken(oathToken string) ClientOption {
	return func(opts clientOptions) clientOptions {
		opts.OathToken = oathToken
		return opts
	}
}

func NewClient(opts ...ClientOption) Client {
	return nil
}
