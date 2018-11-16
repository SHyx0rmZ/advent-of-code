package day05

import "strconv"

type problem struct{}

func Problem() *problem {
	return &problem{}
}

func (problem) PartOne(data []byte) (string, error) {
	steps, err := JumpStrange(string(data))
	return strconv.Itoa(steps), err
}

func (problem) PartTwo(data []byte) (string, error) {
	steps, err := JumpEvenStranger(string(data))
	return strconv.Itoa(steps), err
}
