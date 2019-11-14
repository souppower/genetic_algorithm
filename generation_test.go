package genetic_algorithm

import (
	"sort"
	"testing"
)

func TestCreatePopulation(t *testing.T) {
	target := []byte("To be or not to be")
	res := NewPopulation(target)

	if len(res) != PopulationSize {
		t.Errorf("Expected %d, but got %d", PopulationSize, res)
	}
}

func TestGenerateNextGeneration(t *testing.T) {
	oo := []*Organism{
		{
			Fitness: 0,
		},
		{
			Fitness: 0.1,
		},
		{
			Fitness: 0.1,
		},
		{
			Fitness: 0.2,
		},
		{
			Fitness: 0.2,
		},
		{
			Fitness: 0.2,
		},
	}
	res := GenerateNextGeneration(oo)

	if !containsMoreHighFitnessThanLowerFitness(res) {
		t.Errorf("Should contain more high fitness organisms than low fitness organisms.")
	}
}

func TestContainsMoreHighFitnessThanLowerFitness(t *testing.T) {
	oo := []*Organism{
		{
			Fitness: 0,
		},
		{
			Fitness: 0.1,
		},
		{
			Fitness: 0.2,
		},
		{
			Fitness: 0.1,
		},
		{
			Fitness: 0.2,
		},
		{
			Fitness: 0.2,
		},
	}
	expected := true

	res := containsMoreHighFitnessThanLowerFitness(oo)

	if res != expected {
		t.Errorf("High fitness organisms should exceed lower fitness organisms. Expected %t, but got %t", expected, res)
	}
}

func TestContainsMoreHighFitnessThanLowerFitnessInvalid(t *testing.T) {
	oo := []*Organism{
		{
			Fitness: 0.1,
		},
		{
			Fitness: 0.1,
		},
		{
			Fitness: 0.2,
		},
		{
			Fitness: 0.2,
		},
	}
	expected := false

	res := containsMoreHighFitnessThanLowerFitness(oo)

	if res != expected {
		t.Errorf("Expected %t, but got %t", expected, res)
	}
}

func containsMoreHighFitnessThanLowerFitness(oo []*Organism) bool {
	memo := make(map[int]int, 0)
	for _, r := range oo {
		ratio := int(r.Fitness * 100)
		memo[ratio]++
	}

	keys := make([]int, 0)
	for k := range memo {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	for i := 0; i < len(keys)-1; i++ {
		comp := memo[keys[i]]
		for j := i + 1; j < len(keys); j++ {
			if memo[keys[j]] <= comp {
				return false
			}
		}
	}

	return true
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

	res := getOrganismWithBestFitness(population)

	if res.Fitness != expected {
		t.Errorf("Expected %f, but got %f", expected, res.Fitness)
	}
}
