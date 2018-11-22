package generator

import (
	"math/rand"
)

// ID returns new 12 len id for smth
func ID() string {
	const (
		length = 12
	)

	var (
		letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	)

	b := make([]rune, length)

	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}

	return string(b)
}
