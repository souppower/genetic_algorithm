package genetic_algorithm

import (
	"bytes"
	"fmt"
	"testing"
)

func TestEvolve(t *testing.T) {
	target := []byte("To be or not to be")
	res, generation := Evolve(target)
	fmt.Printf("Got %s, generation: %d", res, generation)

	if bytes.Compare(res, target) != 0 {
		t.Errorf("Expected %s, but got %s", target, res)
	}
}

func TestReproduce(t *testing.T) {
	target := []byte("abcd")
	population := []*Organism{
		{
			DNA:     []byte("abcd"),
		},
		{
			DNA:     []byte("abdc"),
		},
		{
			DNA:     []byte("abxx"),
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
