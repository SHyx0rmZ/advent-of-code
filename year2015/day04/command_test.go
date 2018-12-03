package day04_test

import (
	"strings"
	"testing"

	"github.com/SHyx0rmZ/advent-of-code/year2015/day04"
)

func TestProblem_PartOneWithReader(t *testing.T) {
	for _, tt := range []struct {
		Input  string
		Output string
	}{
		{"abcdef", "609043"},
		{"pqrstuv", "1048970"},
	} {
		t.Run(tt.Input, func(t *testing.T) {
			problem := day04.Problem()
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
