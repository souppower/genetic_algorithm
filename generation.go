package genetic_algorithm

const PopulationSize = 300

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
