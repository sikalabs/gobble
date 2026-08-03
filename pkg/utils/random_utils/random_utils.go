package random_utils

import (
	"math/rand"
	"time"
)

func RandomString(length int) string {
	const CHARS = "abcdefghijklmnopqrstuvwxyz0123456789"
	randSrc := rand.NewSource(time.Now().UnixNano())
	randGen := rand.New(randSrc)
	b := make([]byte, length)
	for i := range b {
		b[i] = CHARS[randGen.Intn(len(CHARS))]
	}
	return string(b)
}
