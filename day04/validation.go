package day04

import (
	"strings"

	"github.com/SHyx0rmZ/advent-of-code/lib"
)

func Valid(p string) bool {
	if p == "" {
		return false
	}
	s := lib.Set()
	for _, word := range strings.Split(p, " ") {
		if s.Contains(word) {
			return false
		}
		s.Add(word)
	}
	return true
}

func ValidEx(p string) bool {
	if p == "" {
		return false
	}
	s := lib.Set()
	for _, word := range strings.Split(p, " ") {
		h := hash{
			M: make(map[rune]int),
		}
		for _, r := range word {
			h.Add(r)
		}
		word := h.String()
		if s.Contains(word) {
			return false
		}
		s.Add(word)
	}
	return true
}
