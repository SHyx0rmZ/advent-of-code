package day02

import "strconv"

type problem struct{}

func Problem() *problem {
	return &problem{}
}

func (problem) PartOne(data []byte) (string, error) {
	sum, err := ChecksumMinMax(data)
	return strconv.Itoa(sum), err
}

func (problem) PartTwo(data []byte) (string, error) {
	sum, err := ChecksumDivision(data)
	return strconv.Itoa(sum), err
}
