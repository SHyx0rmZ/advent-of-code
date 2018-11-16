package day12_test

import (
	"github.com/SHyx0rmZ/advent-of-code/year2017/day12"
	"testing"
)

func TestProblem_PartOne(t *testing.T) {
	r, err := day12.Problem().PartOne([]byte("0 <-> 2\n1 <-> 1\n2 <-> 0, 3, 4\n3 <-> 2, 4\n4 <-> 2, 3, 6\n5 <-> 6\n6 <-> 4, 5\n"))
	if r != "6" || err != nil {
		t.Errorf("got (%#v, %+v), want (%#v, %+v)", r, err, "6", nil)
	}
}

func TestProblem_PartTwo(t *testing.T) {
	r, err := day12.Problem().PartTwo([]byte("0 <-> 2\n1 <-> 1\n2 <-> 0, 3, 4\n3 <-> 2, 4\n4 <-> 2, 3, 6\n5 <-> 6\n6 <-> 4, 5\n"))
	if r != "2" || err != nil {
		t.Errorf("got (%#v, %+v), want (%#v, %+v)", r, err, "2", nil)
	}
}
