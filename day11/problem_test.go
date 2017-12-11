package day11_test

import (
	"github.com/SHyx0rmZ/advent-of-code/day11"
	"testing"
)

func TestProblem_PartOne1(t *testing.T) {
	r, err := day11.Problem().PartOne([]byte("ne,ne,ne"))
	if r != "3" || err != nil {
		t.Errorf("got (%s, %+v), want (%s, %+v)", r, err, "3", nil)
	}
}

func TestProblem_PartOne2(t *testing.T) {
	r, err := day11.Problem().PartOne([]byte("ne,ne,sw,sw"))
	if r != "0" || err != nil {
		t.Errorf("got (%s, %+v), want (%s, %+v)", r, err, "0", nil)
	}
}

func TestProblem_PartOne3(t *testing.T) {
	r, err := day11.Problem().PartOne([]byte("ne,ne,s,s"))
	if r != "2" || err != nil {
		t.Errorf("got (%s, %+v), want (%s, %+v)", r, err, "2", nil)
	}
}

func TestProblem_PartOne4(t *testing.T) {
	r, err := day11.Problem().PartOne([]byte("se,sw,se,sw,sw"))
	if r != "3" || err != nil {
		t.Errorf("got (%s, %+v), want (%s, %+v)", r, err, "3", nil)
	}
}
