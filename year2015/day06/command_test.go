package day06_test

import (
	"strings"
	"testing"

	"github.com/SHyx0rmZ/advent-of-code/year2015/day06"
)

func TestProblem_PartOneWithReader(t *testing.T) {
	for _, tt := range []struct {
		Input  string
		Output string
	}{
		{"turn on 0,0 through 999,999", "1000000"},
		{"toggle 0,0 through 999,0", "1000"},
		{"turn off 499,499 through 500,500", "0"},
	} {
		t.Run(tt.Input, func(t *testing.T) {
			problem := day06.Problem()
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
		{"turn on 0,0 through 0,0", "1"},
		{"toggle 0,0 through 999,999", "2000000"},
	} {
		t.Run(tt.Input, func(t *testing.T) {
			problem := day06.Problem()
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
