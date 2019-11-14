package genetic_algorithm

import (
	"bytes"
	"testing"
)

func TestEvolve(t *testing.T) {
	target := []byte("To be or not to be")
	match, generation := Evolve(target)

	t.Logf("Generation: %d", generation)

	if bytes.Compare(match, target) != 0 {
		t.Errorf("Expected %s, but got %s", target, match)
	}
}

func TestReproduce(t *testing.T) {
	target := []byte("abcd")
	pool := []*Organism{
		{
			DNA: []byte("abcd"),
		},
		{
			DNA: []byte("abdc"),
		},
		{
			DNA: []byte("abxx"),
		},
	}

	res := reproduce(pool, target)

	if len(res) != PopulationSize {
		t.Errorf("Expected %v, but got %v", PopulationSize, len(res))
	}

}
