package day01_test

import (
	"strings"
	"testing"

	"github.com/SHyx0rmZ/advent-of-code/year2015/day01"
)

func TestProblem_PartOneWithReader(t *testing.T) {
	for _, tt := range []struct {
		Input  string
		Output string
	}{
		{"(())", "0"},
		{"()()", "0"},
		{"))(((((", "3"},
		{"())", "-1"},
		{"))(", "-1"},
		{")))", "-3"},
		{")())())", "-3"},
	} {
		t.Run(tt.Input, func(t *testing.T) {
			problem := day01.Problem()
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
		{")", "1"},
		{"()())", "5"},
	} {
		t.Run(tt.Input, func(t *testing.T) {
			problem := day01.Problem()
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
