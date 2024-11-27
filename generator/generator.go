package generator

import (
	"math/rand"
	"time"
)

// GenerateHostname generates a random hostname by combining words from the lists
func GenerateHostname() string {
	rand.Seed(time.Now().UnixNano())

	// Shuffle slices for randomness
	shuffle(Adjectives)
	shuffle(Colors)
	shuffle(Animals)

	// Pick random elements
	adjective := capitalize(Adjectives[rand.Intn(len(Adjectives))])
	color := capitalize(Colors[rand.Intn(len(Colors))])
	animal := capitalize(Animals[rand.Intn(len(Animals))])

	// Combine into a hostname
	return adjective + "-" + color + "-" + animal
}
