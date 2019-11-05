package genetic_algorithm

import (
	"math/rand"
)

const PopulationSize = 300

func Evolve(target []byte) []byte {
	NewOrganism(target)
	return target
}

func generateRandomText(length int) []byte {
	text := make([]byte, length)
	for i := range text {
		// see https://en.wikibooks.org/wiki/C%2B%2B_Programming/ASCII
		text[i] = byte(rand.Intn(95) + 32)
	}
	return text
}

func NewPopulation(target []byte) []*Organism {
	population := make([]*Organism, PopulationSize)
	for i := range population {
		population[i] = NewOrganism(target)
	}
	return population
}
