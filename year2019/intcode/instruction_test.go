package intcode_test

import (
	"testing"

	"github.com/SHyx0rmZ/advent-of-code/year2019/intcode"
)

func TestNewProgram(t *testing.T) {
	for _, tt := range []struct {
		Instruction intcode.Instruction
		Operands    []intcode.Operand
		Check       func([]intcode.Operand) bool
	}{
		{
			intcode.Add{},
			[]intcode.Operand{&intcode.Immediate{3}, &intcode.Immediate{4}, &intcode.Immediate{0}},
			func(operands []intcode.Operand) bool {
				return operands[2].Read() == 7
			},
		},
		{
			intcode.Mul{},
			[]intcode.Operand{&intcode.Immediate{3}, &intcode.Immediate{4}, &intcode.Immediate{0}},
			func(operands []intcode.Operand) bool {
				return operands[2].Read() == 12
			},
		},
	} {
		t.Run(t.Name(), func(t *testing.T) {
			tt.Instruction.Execute(tt.Operands...)
			if !tt.Check(tt.Operands) {
				t.Fail()
			}
		})
	}
}
