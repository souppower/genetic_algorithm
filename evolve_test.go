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

func TestReproduce(t *testing.T) {
	target := []byte("abcd")
	population := []*Organism{
		{
			DNA:     []byte("abcd"),
			Fitness: 0.1,
		},
		{
			DNA:     []byte("abdc"),
			Fitness: 0,
		},
		{
			DNA:     []byte("abxx"),
			Fitness: 0.23,
		},
	}

	res := reproduce(population, target)

	if len(res) != len(population) {
		t.Errorf("Expected %v, but got %v", len(population), len(res))
	}

}

