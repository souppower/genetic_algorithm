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

func GenerateNextGeneration(population []*Organism, target []byte) []*Organism {
	bestFitness := getBestFitness(population)
	nextGeneration := make([]*Organism, 0)

	for _, pop := range population {
		ratio := int(pop.Fitness / bestFitness * 100)

		for i := 0; i < ratio; i++ {
			nextGeneration = append(nextGeneration, pop)
		}
	}

	return nextGeneration
}

func getBestFitness(population []*Organism) float64 {
	var bestFitness float64
	for _, pop := range population {
		if bestFitness < pop.Fitness {
			bestFitness = pop.Fitness
		}
	}
	return bestFitness
}

