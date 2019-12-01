package day01

import (
	"bufio"
	"fmt"
	"io"
	"strconv"

	aoc "github.com/SHyx0rmZ/advent-of-code"
)

type problem struct {
}

func Problem() aoc.ReaderAwareProblem {
	return &problem{}
}

func parse(r io.Reader) ([]int, error) {
	var ns []int
	s := bufio.NewScanner(r)
	for s.Scan() {
		n, err := strconv.Atoi(s.Text())
		if err != nil {
			return nil, err
		}
		ns = append(ns, n)
	}
	return ns, nil
}

func (p problem) PartOneWithReader(r io.Reader) (string, error) {
	ns, err := parse(r)
	if err != nil {
		return "", err
	}
	var t int
	for _, n := range ns {
		t += (n / 3) - 2
	}
	return strconv.Itoa(t), nil
}

func (p problem) PartTwoWithReader(r io.Reader) (string, error) {
	ns, err := parse(r)
	if err != nil {
		return "", err
	}
	var t int
	for _, n := range ns {
		m := ((n) / 3) - 2
		fmt.Println(ns, m)
		x := m
		for {
			f := (x / 3) - 2
			fmt.Println(f, x, m)
			if f < 0 {
				break
			}
			m += f
			x = f
		}
		t += m
	}
	return strconv.Itoa(t), nil
}
