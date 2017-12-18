package day18_test

import (
	"testing"

	"github.com/SHyx0rmZ/advent-of-code/day18"
)

var data = []byte(`set a 1
add a 2
mul a a
mod a 5
snd a
set a 0
rcv a
jgz a -1
set a 1
jgz a -2
`)

func TestProblem_PartOne(t *testing.T) {
	r, err := day18.Problem().PartOne(data)
	if r != "4" || err != nil {
		t.Errorf("got (%#v, %+v), want (%#v, %+v)", r, err, "4", nil)
	}
}

func TestProblem_PartTwo(t *testing.T) {
	r, err := day18.Problem().PartTwo([]byte(""))
	if r != "" || err != nil {
		t.Errorf("got (%#v, %+v), want (%#v, %+v)", r, err, "", nil)
	}
}
