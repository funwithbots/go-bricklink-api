package util

import "math/rand"

func RandomString(length int, rnd *rand.Rand) string {
	const allowed = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = allowed[rnd.Intn(len(allowed))]
	}
	return string(b)
}
