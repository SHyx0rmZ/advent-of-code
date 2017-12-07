package day05

import (
	"strconv"
	"strings"
)

func JumpStrange(s string) (int, error) {
	list, err := newJumpList(s)
	if err != nil {
		return 0, err
	}
	return evalJumpList(list, incOnly), nil
}

func JumpEvenStranger(s string) (int, error) {
	list, err := newJumpList(s)
	if err != nil {
		return 0, err
	}
	return evalJumpList(list, incDec), nil
}

func evalJumpList(list []int, op func(int) int) int {
	var p int
	var steps int
	for {
		o := list[p]
		list[p] += op(o)
		p += o
		steps++

		if p < 0 || p > len(list)-1 {
			break
		}
	}
	return steps
}

func incDec(o int) int {
	if o >= 3 {
		return -1
	}
	return 1
}

func incOnly(o int) int {
	return 1
}

func newJumpList(s string) ([]int, error) {
	var list []int
	for _, line := range strings.Split(s, "\n") {
		i, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		list = append(list, i)
	}
	return list, nil
}
