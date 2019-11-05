package genetic_algorithm

import "math/rand"

type Organism struct {
	DNA     []byte
	Fitness float64
}

func NewOrganism(target []byte) *Organism {
	text := generateRandomText(len(target))

	o := &Organism{
		DNA:     text,
		Fitness: 0,
	}
	o.calculateFitness(target)

	return o
}

func (o *Organism) calculateFitness(target []byte) {
	var match int
	for i := range o.DNA {
		if o.DNA[i] == target[i] {
			match++
		}
	}

	o.Fitness = float64(match) / float64(len(o.DNA))
}

func generateRandomText(length int) []byte {
	text := make([]byte, length)
	for i := range text {
		// see https://en.wikibooks.org/wiki/C%2B%2B_Programming/ASCII
		text[i] = byte(rand.Intn(95) + 32)
	}
	return text
}
