package generator

import (
	"math/rand"
)

// GenerateHostname generates a random hostname by combining words from the lists
func GenerateHostname() string {
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
