package day13_test

import (
	"github.com/SHyx0rmZ/advent-of-code/year2017/day13"
	"testing"
)

func TestProblem_PartOne(t *testing.T) {
	r, err := day13.Problem().PartOne([]byte("0: 3\n1: 2\n4: 4\n6: 4\n"))
	if r != "24" || err != nil {
		t.Errorf("got (%#v, %+v), want (%#v, %+v)", r, err, "24", nil)
	}
}

func TestProblem_PartTwo(t *testing.T) {
	r, err := day13.Problem().PartTwo([]byte("0: 3\n1: 2\n4: 4\n6: 4\n"))
	if r != "10" || err != nil {
		t.Errorf("got (%#v, %+v), want (%#v, %+v)", r, err, "10", nil)
	}
}
