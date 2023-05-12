package go_bricklink_api

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/joho/godotenv"

	"github.com/funwithbots/go-bricklink-api/internal"
	"github.com/funwithbots/go-bricklink-api/internal/oauth"
)

const (
	oAuthSignatureMethod = "HMAC-SHA1"
	oAuthVersion         = "1.0"

	defaultContextTimeout = 3 * time.Second
	defaultBaseURL        = "https://api.bricklink.com/api/store/v1"
)

// Bricklink is the client for the Bricklink API.
type Bricklink struct {
	Client internal.HTTPClient

	baseURL          string
	consumerKey      string
	consumerSecret   string
	oAuthTokenKey    string
	oAuthTokenSecret string
	oAuth            oauth.OAuth1
	Timeout          time.Duration
	Rand             *rand.Rand
}

type BricklinkOption func(opts *Bricklink)

type bricklinkOptions struct {
	client           internal.HTTPClient
	baseURL          string
	consumerKey      string
	consumerSecret   string
	oAuthTokenKey    string
	oAuthTokenSecret string
	timeout          time.Duration
}

func (bl *Bricklink) withOpts(opts []BricklinkOption) {
	bl.Timeout = defaultContextTimeout
	bl.baseURL = defaultBaseURL
	bl.Client = http.DefaultClient

	for _, opt := range opts {
		opt(bl)
	}
}

func WithHTTPClient(client internal.HTTPClient) BricklinkOption {
	return func(opts *Bricklink) {
		opts.Client = client
	}
}

// WithEnv reads the environment variables from the specified files and sets them on the client.
// If the files are not specified, it will read from .env in the project root.
// If the expected file is not found, it will do nothing.
func WithEnv(files ...string) BricklinkOption {
	return func(opts *Bricklink) {
		if env, err := godotenv.Read(files...); err == nil {
			if v, ok := env["BL_API_CONSUMER_KEY"]; ok {
				opts.consumerKey = v
			}
			if v, ok := env["BL_API_CONSUMER_SECRET"]; ok {
				opts.consumerSecret = v
			}
			if v, ok := env["BL_API_TOKEN_KEY"]; ok {
				opts.oAuthTokenKey = v
			}
			if v, ok := env["BL_API_TOKEN_SECRET"]; ok {
				opts.oAuthTokenSecret = v
			}
		}
	}
}

func WithConsumerKey(consumerKey string) BricklinkOption {
	return func(opts *Bricklink) {
		opts.consumerKey = consumerKey
	}
}

func WithConsumerSecret(consumerKey string) BricklinkOption {
	return func(opts *Bricklink) {
		opts.consumerSecret = consumerKey
	}
}

func WithOAuthTokenKey(oAuthToken string) BricklinkOption {
	return func(opts *Bricklink) {
		opts.oAuthTokenKey = oAuthToken
	}
}

func WithOAuthTokenSecret(oAuthToken string) BricklinkOption {
	return func(opts *Bricklink) {
		opts.oAuthTokenSecret = oAuthToken
	}
}

func WithTimeout(timeout time.Duration) BricklinkOption {
	return func(opts *Bricklink) {
		opts.Timeout = timeout
	}
}

func WithBaseURL(baseURL string) BricklinkOption {
	return func(opts *Bricklink) {
		opts.baseURL = baseURL
	}
}

// New creates a new Bricklink API wrapper.
func New(opts ...BricklinkOption) (*Bricklink, error) {
	var bl Bricklink
	bl.withOpts(opts)

	if bl.consumerKey == "" {
		return nil, errors.New("consumer key is required")
	}
	if bl.consumerSecret == "" {
		return nil, errors.New("consumer secret is required")
	}
	if bl.oAuthTokenKey == "" {
		return nil, errors.New("oAuth token key is required")
	}
	if bl.oAuthTokenSecret == "" {
		return nil, errors.New("oAuth token secret is required")
	}

	bl.oAuth = oauth.OAuth1{
		ConsumerKey:     bl.consumerKey,
		ConsumerSecret:  bl.consumerSecret,
		AccessToken:     bl.oAuthTokenKey,
		AccessSecret:    bl.oAuthTokenSecret,
		SignatureMethod: oAuthSignatureMethod,
	}

	bl.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

	return &bl, nil
}

// NewRequest creates a new HTTP request with auth headers and the specified method and path.
// The body is optional. If present, it should be sent as an urlencoded strings.NewReader(body).
func (bl *Bricklink) NewRequest(method string, path string, params map[string]string, body []byte) (*http.Request, error) {
	url := bl.baseURL + path
	buf := bytes.NewReader(body)

	req, err := http.NewRequest(method, url, buf)
	if err != nil {
		return nil, err
	}
	if query := fromParams(params); query != "" {
		req.URL.RawQuery = query
	}

	authHeader := bl.auth(method, url, params)
	authHeader = authHeader[:5] + " realm=\"\"," + authHeader[5:]
	req.Header.Set("Authorization", authHeader)

	return req, nil
}

// NewRequestWithContext creates a new HTTP request with auth headers and the specified method and path.
// The body is optional. If present, it should be sent as an urlencoded strings.NewReader(body).
func (bl *Bricklink) NewRequestWithContext(ctx context.Context, method, path string, params map[string]string, body []byte) (*http.Request, error) {
	url := bl.baseURL + path
	buf := bytes.NewReader(body)

	req, err := http.NewRequestWithContext(ctx, method, url, buf)
	if err != nil {
		return nil, err
	}
	if query := fromParams(params); query != "" {
		req.URL.RawQuery = query
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Charset", "utf-8")

	authHeader := bl.auth(method, url, params)
	authHeader = authHeader[:5] + " realm=\"\"," + authHeader[5:]
	req.Header.Set("Authorization", authHeader)

	return req, nil
}

func (bl *Bricklink) auth(method, url string, params map[string]string) string {
	return bl.oAuth.BuildOAuth1Header(method, url, params)
}

func fromParams(params map[string]string) string {
	query := make([]string, 0, len(params))
	for k, v := range params {
		query = append(query, fmt.Sprintf("%s=%s", k, v))
	}

	return strings.Join(query, "&")
}
