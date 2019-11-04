package genetic_algorithm

import (
	"fmt"
	"math/rand"
)

type Organism struct {
	DNA     []byte
	Fitness float64
}

func Evolve(target []byte) []byte {
	organizm := createOrganizm(target)
	fmt.Println(organizm)
	return target
}

func createOrganizm(target []byte) *Organism {
	text := generateRandomText(len(target))

	return &Organism{
		DNA:     text,
		Fitness: 0,
	}
}

func generateRandomText(length int) []byte {
	text := make([]byte, length)
	for i := range text {
		// see https://en.wikibooks.org/wiki/C%2B%2B_Programming/ASCII
		text[i] = byte(rand.Intn(95) + 32)
	}
	return text
}
