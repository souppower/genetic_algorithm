package genetic_algorithm

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
