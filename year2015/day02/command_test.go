package day02_test

import (
	"strings"
	"testing"

	"github.com/SHyx0rmZ/advent-of-code/year2015/day02"
)

func TestProblem_PartOneWithReader(t *testing.T) {
	for _, tt := range []struct {
		Input  string
		Output string
	}{
		{"2x3x4", "58"},
		{"1x1x10", "43"},
	} {
		t.Run(tt.Input, func(t *testing.T) {
			problem := day02.Problem()
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
		{"2x3x4", "34"},
		{"1x1x10", "14"},
	} {
		t.Run(tt.Input, func(t *testing.T) {
			problem := day02.Problem()
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
