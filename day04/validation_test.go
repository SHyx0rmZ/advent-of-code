package day04_test

import (
	"github.com/SHyx0rmZ/advent-of-code/day04"
	"testing"
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

func TestValidAAA(t *testing.T) {
	r := day04.Valid("aa bb cc dd aaa")
	if r != true {
		t.Errorf("got %t, want %t", r, true)
	}
}

func TestValidExFGHIJ(t *testing.T) {
	r := day04.ValidEx("abcde fghij")
	if r != true {
		t.Errorf("got %t, want %t", r, true)
	}
}

func TestValidExECDAB(t *testing.T) {
	r := day04.ValidEx("abcde xyz ecdab")
	if r != false {
		t.Errorf("got %t, want %t", r, false)
	}
}

func TestValidExABJ(t *testing.T) {
	r := day04.ValidEx("a ab abc abd abf abj")
	if r != true {
		t.Errorf("got %t, want %t", r, true)
	}
}

func TestValidExOOOO(t *testing.T) {
	r := day04.ValidEx("iiii oiii ooii oooi oooo")
	if r != true {
		t.Errorf("got %t, want %t", r, true)
	}
}

func TestValidExIIIO(t *testing.T) {
	r := day04.ValidEx("oiii ioii iioi iiio")
	if r != false {
		t.Errorf("got %t, want %t", r, false)
	}
}
