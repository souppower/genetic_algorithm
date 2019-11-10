package genetic_algorithm

import (
	"math/rand"
	"time"
)

func Evolve(target []byte) []byte {
	NewOrganism(target)
	return target
}

func reproduce(population []*Organism, target []byte) []*Organism {
	nextGeneration := make([]*Organism, len(population))

	rand.Seed(time.Now().UTC().UnixNano())

	for i := range population {
		r1 := rand.Intn(len(population))
		r2 := rand.Intn(len(population))

		parent1 := population[r1]
		parent2 := population[r2]

		child := parent1.crossOver(parent2, target)
		child.mutate()

		nextGeneration[i] = child
	}

	return nextGeneration
}
