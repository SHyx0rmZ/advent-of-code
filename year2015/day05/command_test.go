package day05_test

import (
	"strings"
	"testing"

	"github.com/SHyx0rmZ/advent-of-code/year2015/day05"
)

func TestProblem_PartOneWithReader(t *testing.T) {
	for _, tt := range []struct {
		Input  string
		Output string
	}{
		{"ugknbfddgicrmopn", "1"},
		{"aaa", "1"},
		{"jchzalrnumimnmhp", "0"},
		{"haegwjzuvuyypxyu", "0"},
		{"dvszwmarrgswjxmb", "0"},
	} {
		t.Run(tt.Input, func(t *testing.T) {
			problem := day05.Problem()
			r, err := problem.PartOneWithReader(strings.NewReader(tt.Input))
			if err != nil {
				t.Error(err)
			}
			if r != tt.Output {
				t.Errorf("expected %+v, got %+v", tt.Output, r)
			}
		})
	}
}

func TestProblem_PartTwoWithReader(t *testing.T) {
	for _, tt := range []struct {
		Input  string
		Output string
	}{
		{"qjhvhtzxzqqjkmpb", "1"},
		{"xxyxx", "1"},
		{"uurcxstgmygtbstg", "0"},
		{"ieodomkazucvgmuy", "0"},
	} {
		t.Run(tt.Input, func(t *testing.T) {
			problem := day05.Problem()
			r, err := problem.PartTwoWithReader(strings.NewReader(tt.Input))
			if err != nil {
				t.Error(err)
			}
			if r != tt.Output {
				t.Errorf("expected %+v, got %+v", tt.Output, r)
			}
		})
	}
}
