package aoc

type Problem interface {
	PartOne([]byte) (string, error)
	PartTwo([]byte) (string, error)
}
