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
		fmt.Println("===============")
		fmt.Println(string(bestOrganism.DNA))
		fmt.Println(bestOrganism.Fitness)
		fmt.Println(generation)

		if bytes.Compare(bestOrganism.DNA, target) == 0 {
			found = true
		} else {
			population = reproduce(population, target)
		}
	}

	return bestOrganism.DNA, generation
}

func reproduce(population []*Organism, target []byte) []*Organism {
	nextGeneration := make([]*Organism, len(population))

	rand.Seed(time.Now().UTC().UnixNano())

	for i := 0; i < len(population); i++ {
		r1 := rand.Intn(len(population))
		r2 := rand.Intn(len(population))

		parent1 := population[r1]
		parent2 := population[r2]

		child := parent1.crossOver(parent2, target)
		child.mutate()
		child.calculateFitness(target)

		nextGeneration[i] = child
	}

	return nextGeneration
}
