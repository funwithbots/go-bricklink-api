package util

import (
	"html"
	"math/rand"
)

func RandomString(length int, rnd *rand.Rand) string {
	const allowed = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = allowed[rnd.Intn(len(allowed))]
	}
	return string(b)
}

// NormalizeString processes a string to make it more readable.
// Currently, it only unescapes the string.
func NormalizeString(s string) string {
	return html.UnescapeString(html.UnescapeString(s))
}
