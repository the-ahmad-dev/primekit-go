package random

import (
	"math/rand"
	"time"
)

// RandomInt returns a random integer between min and max (inclusive)
func RandomInt(min, max int) int {
	return rand.Intn(max-min+1) + min
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// RandomString returns a random string of the specified length
func RandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

// SecureRandomInt returns a cryptographically secure random integer
func SecureRandomInt(min, max int) (int, error) {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	return rng.Intn(max-min+1) + min, nil
}
