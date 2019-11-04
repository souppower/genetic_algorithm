package genetic_algorithm

import (
	"bytes"
	"testing"
)

func TestEvolve(t *testing.T) {
	target := []byte("To be or not to be")
	res := Evolve(target)

	if bytes.Compare(res, target) != 0 {
		t.Errorf("Expected %s, but got %s", target, res)
	}
}

func TestCreatePopulation(t *testing.T) {
	target := []byte("To be or not to be")
	res := createPopulation(target)

	if len(res) != PopulationSize {
		t.Errorf("Expected %d, but got %d", PopulationSize, res)
	}
}
