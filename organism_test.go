package genetic_algorithm

import (
	"testing"
)

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

func TestCrossOver(t *testing.T) {
	target := []byte("oooo")
	o := &Organism{
		DNA: []byte("abcd"),
	}
	o1 := &Organism{
		DNA: []byte("xxxx"),
	}
	expected := float64(0)
	newOrganism := o.crossOver(o1, target)

	if newOrganism.Fitness != expected {
		t.Errorf("Expected %f, but got %f", expected, o.Fitness)
	}
}
