package day03_test

import (
	"strings"
	"testing"

	"github.com/SHyx0rmZ/advent-of-code/year2015/day03"
)

func TestProblem_PartOneWithReader(t *testing.T) {
	for _, tt := range []struct {
		Input  string
		Output string
	}{
		{">", "2"},
		{"^>v<", "4"},
		{"^v^v^v^v^v", "2"},
	} {
		t.Run(tt.Input, func(t *testing.T) {
			problem := day03.Problem()
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
		{"^v", "3"},
		{"^>v<", "3"},
		{"^v^v^v^v^v", "11"},
	} {
		t.Run(tt.Input, func(t *testing.T) {
			problem := day03.Problem()
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
