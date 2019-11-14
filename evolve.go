package genetic_algorithm

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
)

func Evolve(target []byte) ([]byte, int) {
	population := NewPopulation(target)
	var bestOrganism *Organism
	generation := 0
	found := false

	for !found {
		generation++

		bestOrganism = getOrganismWithBestFitness(population)
		fmt.Printf("Generation: %d, DNA: %s, Fitness: %.1f%%\n", generation, string(bestOrganism.DNA), bestOrganism.Fitness*100)

		if bytes.Compare(bestOrganism.DNA, target) == 0 {
			found = true
		} else {
			pool := createPool(population)
			population = reproduce(pool, target)
		}
	}

	return bestOrganism.DNA, generation
}

func reproduce(pool []*Organism, target []byte) []*Organism {
	nextGeneration := make([]*Organism, PopulationSize)

	rand.Seed(time.Now().UTC().UnixNano())

	for i := 0; i < PopulationSize; i++ {
		r1 := rand.Intn(len(pool))
		r2 := rand.Intn(len(pool))

		parent1 := pool[r1]
		parent2 := pool[r2]

		child := parent1.crossOver(parent2, target)
		child.mutate()
		child.calculateFitness(target)

		nextGeneration[i] = child
	}

	return nextGeneration
}
