package genetic_algorithm

import (
	"math/rand"
	"time"
)

const MutationRate = 0.01

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

func (o *Organism) crossOver(other *Organism, target []byte) *Organism {
	child := &Organism{
		DNA: make([]byte, len(o.DNA)),
	}

	rand.Seed(time.Now().UTC().UnixNano())
	mid := rand.Intn(len(o.DNA))

	for i := 0; i < len(o.DNA); i++ {
		if i <= mid {
			child.DNA[i] = o.DNA[i]
		} else {
			child.DNA[i] = other.DNA[i]
		}
	}

	child.calculateFitness(target)

	return child
}

func (o *Organism) mutate() {
	for i := 0; i < len(o.DNA); i++ {
		if rand.Float64() < MutationRate {
			o.DNA[i] = generateRandomLetter()
		}
	}
}

func generateRandomText(length int) []byte {
	text := make([]byte, length)
	for i := range text {
		// see https://en.wikibooks.org/wiki/C%2B%2B_Programming/ASCII
		text[i] = generateRandomLetter()
	}
	return text
}

func generateRandomLetter() byte {
	rand.Seed(time.Now().UTC().UnixNano())
	return byte(rand.Intn(95) + 32)
}
