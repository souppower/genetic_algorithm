package genetic_algorithm

import (
	"bytes"
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
	newOrganism := o.crossOver(o1, target)

	if bytes.Compare(newOrganism.DNA, o.DNA) == 0 || bytes.Compare(newOrganism.DNA, o1.DNA) == 0 {
		t.Errorf("New organism contains DNA identical to a parent. DNA: %s", newOrganism.DNA)
	}
}
