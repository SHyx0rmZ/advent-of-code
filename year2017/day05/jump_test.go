package day05_test

import (
	"github.com/SHyx0rmZ/advent-of-code/year2017/day05"
	"testing"
)

func TestJumpStrange(t *testing.T) {
	r, err := day05.JumpStrange("0\n3\n0\n1\n-3")
	if r != 5 || err != nil {
		t.Errorf("got (%d, %+v), want (%d, %+v)", r, err, 5, nil)
	}
}

func TestJumpEvenStranger(t *testing.T) {
	r, err := day05.JumpEvenStranger("0\n3\n0\n1\n-3")
	if r != 10 || err != nil {
		t.Errorf("got (%d, %+v), want (%d, %+v)", r, err, 10, nil)
	}
}
