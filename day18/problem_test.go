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

var data2 = []byte(`snd 1
snd 2
snd p
rcv a
rcv b
rcv c
rcv d
`)

func TestProblem_PartOne(t *testing.T) {
	r, err := day18.Problem().PartOne(data)
	if r != "4" || err != nil {
		t.Errorf("got (%#v, %+v), want (%#v, %+v)", r, err, "4", nil)
	}
}

//func TestProblem_PartTwo(t *testing.T) {
//	r, err := day18.Problem().PartTwo(data)
//	if r != "1" || err != nil {
//		t.Errorf("got (%#v, %+v), want (%#v, %+v)", r, err, "1", nil)
//	}
//}

func TestProblem_PartTwo2(t *testing.T) {
	r, err := day18.Problem().PartTwo(data2)
	if r != "3" || err != nil {
		t.Errorf("got (%#v, %+v), want (%#v, %+v)", r, err, "3", nil)
	}
}
