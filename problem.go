package aoc

import "io"

type Problem interface {
	PartOne([]byte) (string, error)
	PartTwo([]byte) (string, error)
}

type ReaderAwareProblem interface {
	PartOneWithReader(io.Reader) (string, error)
	PartTwoWithReader(io.Reader) (string, error)
}
