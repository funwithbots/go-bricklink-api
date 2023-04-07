package go_bricklink_api

import (
	"context"
	"crypto/hmac"
	"crypto/sha1"
	"errors"
	"net/http"

	"github.com/joho/godotenv"

	"github.com/funwithbots/go-bricklink-api/util"
)

const (
	apiURL               = "https://api.bricklink.com/api/store/v1"
	oauthSignatureMethod = "HMAC-SHA1"
	oauthVersion         = "1.0"
)

type Client interface {
	NewRequest(method string, baseUrl string) (*http.Request, error)
	Do(ctx context.Context, req *http.Request) (*http.Response, error)
}

type client struct {
	HTTP        *http.Client
	ConsumerKey string
	OathToken   string
}

func (c *client) NewRequest(method string, baseUrl string) (*http.Request, error) {
	req, err := http.NewRequest(method, baseUrl, nil)
	if err != nil {
		return nil, util.ErrInvalidArgument
	}
	return req, util.ErrNotImplemented
}

func (c *client) Do(ctx context.Context, req *http.Request) (*http.Response, error) {
	c.addHash(req)
	return c.HTTP.Do(req)
}

func (c *client) addHash(req *http.Request) {
	h := hmac.New(sha1.New, []byte(c.ConsumerKey+c.OathToken+req.URL.String()))
	req.Header.Add("oath_signature", string(h.Sum(nil)))
}

type ClientOption func(opts *client)

type clientOptions struct {
	HTTP        *http.Client
	ConsumerKey string
	OathToken   string
}

func (c *client) withOpts(opts []ClientOption) {
	for _, opt := range opts {
		opt(c)
	}
}

func WithHTTPClient(http *http.Client) ClientOption {
	return func(opts *client) {
		opts.HTTP = http
	}
}

// WithEnv reads the environment variables from the specified files and sets them on the client.
// If the files are not specified, it will read from .env in the project root.
// If the expected file is not found, it will do nothing.
func WithEnv(files ...string) ClientOption {
	return func(opts *client) {
		if env, err := godotenv.Read(files...); err == nil {
			if key, ok := env["BL_API_KEY"]; ok {
				opts.ConsumerKey = key
			}
			if token, ok := env["BL_API_TOKEN"]; ok {
				opts.OathToken = token
			}
		}

	}
}

func WithConsumerKey(consumerKey string) ClientOption {
	return func(opts *client) {
		opts.ConsumerKey = consumerKey
	}
}

func WithOathToken(oathToken string) ClientOption {
	return func(opts *client) {
		opts.OathToken = oathToken
	}
}

func NewClient(opts ...ClientOption) (Client, error) {
	c := &client{
		HTTP: http.DefaultClient,
	}
	c.withOpts(opts)

	if c.ConsumerKey == "" {
		return nil, errors.New("consumer key is required")
	}

	if c.OathToken == "" {
		return nil, errors.New("oath token is required")
	}

	return c, nil
}
