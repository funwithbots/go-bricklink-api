package oauth

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"math/rand"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type OAuth1 struct {
	ConsumerKey     string
	ConsumerSecret  string
	AccessToken     string
	AccessSecret    string
	SignatureMethod string
}

const headerPlaceholder = "OAuth oauth_consumer_key=\"%s\", oauth_nonce=\"%s\", oauth_signature=\"%s\", oauth_signature_method=\"%s\", oauth_timestamp=\"%s\", oauth_token=\"%s\", oauth_version=\"%s\""

// Params being any key-value url query parameter pairs
func (auth OAuth1) BuildOAuth1Header(method, path string, params map[string]string) string {
	vals := auth.getURLVals(params)
	parameterString := strings.ReplaceAll(vals.Encode(), "+", "%20")

	// Calculating Signature Base String and Signing Key
	signatureBase := buildQueryString(
		strings.ToUpper(method),
		url.QueryEscape(strings.Split(path, "?")[0]),
		url.QueryEscape(parameterString),
	)
	signingKey := buildQueryString(
		url.QueryEscape(auth.ConsumerSecret),
		url.QueryEscape(auth.AccessSecret),
	)
	signature := calculateSignature(signatureBase, signingKey)
	return auth.headerString(vals, signature)
}

func (auth OAuth1) getURLVals(params map[string]string) url.Values {
	vals := url.Values{}
	// vals.Add("realm", "")
	vals.Add("oauth_consumer_key", auth.ConsumerKey)
	vals.Add("oauth_nonce", generateNonce())
	vals.Add("oauth_signature_method", auth.SignatureMethod)
	vals.Add("oauth_timestamp", strconv.Itoa(int(time.Now().Unix())))
	vals.Add("oauth_token", auth.AccessToken)
	vals.Add("oauth_version", "1.0")

	for k, v := range params {
		vals.Add(k, v)
	}
	return vals
}

func (auth OAuth1) headerString(vals url.Values, signature string) string {
	return fmt.Sprintf(
		headerPlaceholder,
		url.QueryEscape(vals.Get("oauth_consumer_key")),
		url.QueryEscape(vals.Get("oauth_nonce")),
		url.QueryEscape(signature),
		url.QueryEscape(vals.Get("oauth_signature_method")),
		url.QueryEscape(vals.Get("oauth_timestamp")),
		url.QueryEscape(vals.Get("oauth_token")),
		url.QueryEscape(vals.Get("oauth_version")),
	)
}

func calculateSignature(base, key string) string {
	hash := hmac.New(sha1.New, []byte(key))
	hash.Write([]byte(base))
	signature := hash.Sum(nil)
	return base64.StdEncoding.EncodeToString(signature)
}

func generateNonce() string {
	const allowed = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 16)
	for i := range b {
		b[i] = allowed[rand.Intn(len(allowed))]
	}
	return string(b)
}

func buildQueryString(inputs ...string) string {
	return strings.Join(inputs, "&")
}
