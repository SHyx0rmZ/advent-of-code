package day06

import (
	"bytes"
	"strconv"
)

type problem struct{}

func Problem() *problem {
	return &problem{}
}

func (p problem) PartOne(data []byte) (string, error) {
	banks, err := p.parse(data)
	if err != nil {
		return "", err
	}

	_, steps := p.balanceMemory(banks)
	return strconv.Itoa(steps), nil
}

func (p problem) PartTwo(data []byte) (string, error) {
	banks, err := p.parse(data)
	if err != nil {
		return "", err
	}

	length, _ := p.balanceMemory(banks)
	return strconv.Itoa(length), nil
}

func (problem) balanceMemory(banks []int) (int, int) {
	m := &memory{}
	m.Reload(banks)
	var steps int
	s := set{
		M: make(map[string]int),
	}
	for {
		steps++
		m.Balance()
		if s.Contains(m.String()) {
			break
		}
		s.Add(m.String())
	}
	return s.LoopSize(m.String()), steps
}

func (problem) parse(data []byte) ([]int, error) {
	var banks []int
	for len(data) > 0 {
		i := bytes.IndexAny(data, "\t \n")
		if i < 1 {
			// if there are no bytes before the whitespace,
			// we must have reached the end of the line
			break
		}
		// parse column
		b, err := strconv.Atoi(string(data[0:i]))
		if err != nil {
			return nil, err
		}
		banks = append(banks, b)
		for data[i] == '\t' || data[i] == ' ' {
			i++
		}
		data = data[i:]
	}
	return banks, nil
}
