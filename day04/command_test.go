package day04_test

import (
	"testing"
	"github.com/SHyx0rmZ/advent-of-code/day04"
)

func TestValidEE(t *testing.T) {
	r := day04.Valid("aa bb cc dd ee")
	if r != true {
		t.Errorf("got %t, want %t", r, true)
	}
}

func TestValidAA(t *testing.T) {
	r := day04.Valid("aa bb cc dd aa")
	if r != false {
		t.Errorf("got %t, want %t", r, false)
	}
}
