package day04

import (
	"sort"
	"strconv"
)

type hash struct {
	M map[rune]int
}

func (h *hash) Add(r rune) {
	h.M[r]++
}

func (h *hash) String() string {
	var runes []int
	for r := range h.M {
		runes = append(runes, int(r))
	}
	sort.Ints(runes)
	word := ""
	for _, r := range runes {
		word += string(rune(r)) + strconv.Itoa(h.M[rune(r)])
	}
	return word
}
