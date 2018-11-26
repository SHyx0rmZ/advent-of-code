package lib

import (
	"strconv"
)

func StringsToInts(strings []string) ([]int, error) {
	ints := make([]int, len(strings))
	for i, s := range strings {
		n, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		ints[i] = n
	}
	return ints, nil
}