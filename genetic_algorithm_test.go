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
	res := NewPopulation(target)

	if len(res) != PopulationSize {
		t.Errorf("Expected %d, but got %d", PopulationSize, res)
	}
}

func TestNewGeneration(t *testing.T) {
	target := []byte("To be or not to be")
	population := NewPopulation(target)
	res := GenerateNextGeneration(population, target)

	if len(res) != PopulationSize {
		t.Errorf("Expected %d, but got %d", PopulationSize, len(res))
	}
}

func TestGetBestFitness(t *testing.T) {
	population := []*Organism{
		{
			Fitness: 0.1,
		},
		{
			Fitness: 0,
		},
		{
			Fitness: 0.23,
		},
	}
	expected := 0.23

	res := getBestFitness(population)

	if res != expected {
		t.Errorf("Expected %f, but got %f", expected, res)
	}
}