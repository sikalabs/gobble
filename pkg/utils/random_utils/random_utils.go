package random_utils

import (
	"math/rand"
	"time"
)

func RandomString(length int) string {
	const CHARS = "abcdefghijklmnopqrstuvwxyz0123456789"
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length)
	for i := range b {
		b[i] = CHARS[rand.Intn(len(CHARS))]
	}
	return string(b)
}
