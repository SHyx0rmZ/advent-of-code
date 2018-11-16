package day18_test

import (
	"github.com/SHyx0rmZ/advent-of-code/year2017/day18"
	"testing"
)

func TestISAHasInstructions(t *testing.T) {
	for _, instruction := range []string{
		"snd",
		"rcv",
		"add",
		"mul",
		"mod",
		"set",
		"jgz",
	} {
		t.Run(instruction, func(t *testing.T) {
			f, ok := day18.ISA[instruction]
			if !ok {
				t.Fatalf("unknown instruction: %s", instruction)
			}
			if f == nil {
				t.Errorf("nil instruction: %s", instruction)
			}
		})
	}
}
