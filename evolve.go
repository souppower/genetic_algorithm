package genetic_algorithm

func Evolve(target []byte) []byte {
	NewOrganism(target)
	return target
}
