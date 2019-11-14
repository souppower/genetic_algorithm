package genetic_algorithm

const PopulationSize = 500

func NewPopulation(target []byte) []*Organism {
	population := make([]*Organism, PopulationSize)
	for i := range population {
		population[i] = NewOrganism(target)
	}
	return population
}

// createPool returns a pool of organisms for the reproduction phase
// More organisms with high fitness and less organisms with less fitness are contained in the pool
func createPool(population []*Organism) []*Organism {
	bestFitness := getOrganismWithBestFitness(population).Fitness
	pool := make([]*Organism, 0)

	for _, pop := range population {
		num := int(pop.Fitness / bestFitness * 100)

		for i := 0; i < num; i++ {
			pool = append(pool, pop)
		}
	}

	return pool
}

func getOrganismWithBestFitness(population []*Organism) *Organism {
	var bestFitness float64
	index := 0
	for i, pop := range population {
		if bestFitness < pop.Fitness {
			index = i
			bestFitness = pop.Fitness
		}
	}
	return population[index]
}
