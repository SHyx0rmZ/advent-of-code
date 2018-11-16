package day08

import (
	"go/constant"
	"go/token"
)

type instruction struct {
	Reg constant.Value
	Op  token.Token
	N   constant.Value
}

func Instruction(reg, op, n []byte, m map[constant.Value]token.Token) instruction {
	return instruction{
		Reg: constant.MakeString(string(reg)),
		Op:  m[constant.MakeString(string(op))],
		N:   constant.MakeFromLiteral(string(n), token.INT, 0),
	}
}
