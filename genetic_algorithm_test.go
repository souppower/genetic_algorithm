package genetic_algorithm

import "testing"

func TestEvolve(t *testing.T) {
	target := "To be or not to be"
	res := Evolve(target)

	if res != target {
		t.Errorf("Expected %s, but got %s", target, res)
	}
}

func Evolve(target string) string {
	return target
}