package genetic_algorithm

import "testing"

func TestCalculateFitness(t *testing.T) {
	o := &Organism{
		DNA: []byte("abdc"),
	}

	o.calculateFitness([]byte("abcd"))
	expected := float64(0.5)

	if o.Fitness != expected {
		t.Errorf("Expected %f, but got %f", expected, o.Fitness)
	}

}

func TestCalculateFitnessIdentical(t *testing.T) {
	o := &Organism{
		DNA: []byte("abcd"),
	}

	o.calculateFitness([]byte("abcd"))
	expected := float64(1)

	if o.Fitness != expected {
		t.Errorf("Expected %f, but got %f", expected, o.Fitness)
	}
}
