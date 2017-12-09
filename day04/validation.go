package day04

import "strings"

func Valid(p string) bool {
	if p == "" {
		return false
	}
	s := set{
		M: make(map[string]struct{}),
	}
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
	s := set{
		M: make(map[string]struct{}),
	}
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
