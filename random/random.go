package random

import (
	"bytes"
	"math/rand"
	"time"
)

var defLetters = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
var defLettersRune = []rune(defLetters)

func String(n int, letters ...string) string {
	// Default letters if none are passed.
	var letterRunes []rune
	if len(letters) > 0 {
		letterRunes = []rune(letters[0])
	} else {
		letterRunes = defLettersRune
	}

	// Create a random source.
	randSource := rand.NewSource(time.Now().UnixNano())
	r := rand.New(randSource)

	var bb bytes.Buffer
	bb.Grow(n)
	l := len(letterRunes)

	// Generate random index and append.
	for i := 0; i < n; i++ {
		randomIndex := r.Intn(l)
		bb.WriteRune(letterRunes[randomIndex])
	}

	return bb.String()
}
