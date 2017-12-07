package day05

import (
	"strconv"
	"strings"
)

func Jump(s string) (int, error) {
	var list []int
	var p int
	var steps int
	for _, line := range strings.Split(s, "\n") {
		i, err := strconv.Atoi(line)
		if err != nil {
			return 0, err
		}
		list = append(list, i)
	}
	for {
		j := list[p]
		list[p]++
		p += j
		steps++

		if p < 0 || p > len(list)-1 {
			break
		}
	}
	return steps, nil
}
