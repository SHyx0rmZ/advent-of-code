package day19_test

import (
	"github.com/SHyx0rmZ/advent-of-code/year2015/day19"
	"strings"
	"testing"
)

func TestProblem_PartOneWithReader(t *testing.T) {
	for _, tt := range []struct {
		Input  string
		Output string
	}{
		{
			Input: `H => HO
H => OH
O => HH

HOH
`,
			Output: "4",
		},
		{
			Input: `H => HO
H => OH
O => HH

HOHOHO
`,
			Output: "7",
		},
	} {
		t.Run(tt.Input, func(t *testing.T) {
			problem := day19.Problem()
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
		{
			Input: `e => H
e => O
H => HO
H => OH
O => HH

HOH
`,
			Output: "3",
		},
		{
			Input: `e => H
e => O
H => HO
H => OH
O => HH

HOHOHO
`,
			Output: "6",
		},
	} {
		t.Run(tt.Input, func(t *testing.T) {
			problem := day19.Problem()
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

func BenchmarkProblem_PartTwoWithReader(t *testing.B) {
	for _, tt := range []struct {
		Input  string
		Output string
	}{
		{
			Input: `e => H
e => O
H => HO
H => OH
O => HH

HOH
`,
			Output: "3",
		},
		{
			Input: `e => H
e => O
H => HO
H => OH
O => HH

HOHOHO
`,
			Output: "6",
		},
	} {
		t.Run(tt.Input, func(t *testing.B) {
			problem := day19.Problem()
			r, err := problem.PartTwoWithReader(strings.NewReader(tt.Input))
			if err != nil {
				t.Error(err)
			}
			if r != tt.Output {
				t.Errorf("expected %+v, got %+v", tt.Output, r)
			}
			sr := strings.NewReader(tt.Input)
			t.ResetTimer()
			for i := 0; i < t.N; i++ {
				problem.PartTwoWithReader(sr)
			}
		})
	}
}
