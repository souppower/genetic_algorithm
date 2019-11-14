package genetic_algorithm

import (
	"bytes"
	"fmt"
	"testing"
)

func TestEvolve(t *testing.T) {
	t.Skip("Skipping long-running test.")

	target := []byte("abcd")
	match, generation := Evolve(target)

	t.Logf("Got %s, generation: %d", match, generation)

	if bytes.Compare(match, target) != 0 {
		t.Errorf("Expected %s, but got %s", target, match)
	}
}

func TestReproduce(t *testing.T) {
	target := []byte("abcd")
	population := []*Organism{
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

	res := reproduce(population, target)

	for _, r := range res {
		fmt.Println("=========================")
		fmt.Printf("Fitness: %f\n", r.Fitness)
		fmt.Printf("DNA: %+v\n", string(r.DNA))
	}

	if len(res) != len(population) {
		t.Errorf("Expected %v, but got %v", len(population), len(res))
	}

}
