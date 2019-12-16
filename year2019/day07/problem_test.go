package day07_test

import (
	"strings"
	"testing"

	"github.com/SHyx0rmZ/advent-of-code/year2019/day07"
)

func TestProblem_PartOne(t *testing.T) {
	for _, tt := range []struct {
		Input  string
		Output string
	}{
		{
			"3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0",
			"4,3,2,1,0",
		},
		{
			"3,23,3,24,1002,24,10,24,1002,23,-1,23,101,5,23,23,1,24,23,23,4,23,99,0,0",
			"0,1,2,3,4",
		},
		{
			"3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33,1002,33,7,33,1,33,31,31,1,32,31,31,4,31,99,0,0,0",
			"1,0,4,3,2",
		},
	} {
		t.Run(tt.Input, func(t *testing.T) {
			r, err := day07.Problem().PartOneWithReader(strings.NewReader(tt.Input))
			if r != tt.Output || err != nil {
				t.Errorf("got (%#v, %+v), want (%#v, %+v)", r, err, tt.Output, nil)
			}
		})
	}
}

func TestProblem_PartTwo(t *testing.T) {
	r, err := day07.Problem().PartTwoWithReader(strings.NewReader(""))
	if r != "" || err != nil {
		t.Errorf("got (%#v, %+v), want (%#v, %+v)", r, err, "", nil)
	}
}
