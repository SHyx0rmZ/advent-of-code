package day14_test

import (
	"github.com/SHyx0rmZ/advent-of-code/year2017/day14"
	"testing"
)

func TestProblem_PartOne(t *testing.T) {
	r, err := day14.Problem().PartOne([]byte("flqrgnkx"))
	if r != "8108" || err != nil {
		t.Errorf("got (%#v, %+v), want (%#v, %+v)", r, err, "8108", nil)
	}
}

func TestProblem_PartTwo(t *testing.T) {
	r, err := day14.Problem().PartTwo([]byte("flqrgnkx"))
	if r != "1242" || err != nil {
		t.Errorf("got (%#v, %+v), want (%#v, %+v)", r, err, "1242", nil)
	}
}
