package day20

import (
	"bufio"
	"io"
	"strconv"

	"github.com/SHyx0rmZ/advent-of-code"
)

type problem struct {
}

func Problem() aoc.ReaderAwareProblem {
	return &problem{}
}

func (p problem) PartOneWithReader(r io.Reader) (string, error) {
	s, err := bufio.NewReader(r).ReadString('\n')
	if err != nil {
		return "", err
	}
	n, err := strconv.Atoi(s[:len(s)-1])
	if err != nil {
		return "", err
	}
	m := make(map[int]int, n)
	for i := 1; i < n; i++ {
		for j := i; j < n; j += i {
			m[j] += i * 10
		}
		p := m[i]
		//fmt.Println("House", i, "got", p, "presents")
		if p >= n {
			return strconv.Itoa(i), nil
		}
	}
	return "?", nil
}

func (p problem) PartTwoWithReader(r io.Reader) (string, error) {
	s, err := bufio.NewReader(r).ReadString('\n')
	if err != nil {
		return "", err
	}
	n, err := strconv.Atoi(s[:len(s)-1])
	if err != nil {
		return "", err
	}
	m := make(map[int]int, n)
	for i := 1; i < n; i++ {
		for j, k := i, 1; k < 51; j, k = j + i, k + 1 {
			m[j] += i * 11
		}
		p := m[i]
		//fmt.Println("House", i, "got", p, "presents")
		if p >= n {
			return strconv.Itoa(i), nil
		}
	}
	return "?", nil
}
