package day08

import (
	"bytes"
	"fmt"
	"go/constant"
	"go/token"
)

var (
	conds = map[constant.Value]token.Token{
		constant.MakeString(">"):  token.GTR,
		constant.MakeString("<"):  token.LSS,
		constant.MakeString(">="): token.GEQ,
		constant.MakeString("<="): token.LEQ,
		constant.MakeString("=="): token.EQL,
		constant.MakeString("!="): token.NEQ,
	}
	ops = map[constant.Value]token.Token{
		constant.MakeString("inc"): token.ADD,
		constant.MakeString("dec"): token.SUB,
	}
)

type problem struct{}

func Problem() *problem {
	return &problem{}
}

func (p problem) PartOne(data []byte) (string, error) {
	prog, err := p.parse(data)
	if err != nil {
		return "", err
	}
	c := CPU()
	c.Run(prog)
	var l int64
	for _, r := range c.Registers() {
		n, ok := constant.Int64Val(r)
		if !ok {
			panic("not an int64")
		}
		if n > l {
			l = n
		}
	}
	return fmt.Sprintf("%d", l), nil
}

func (p problem) PartTwo(data []byte) (string, error) {
	prog, err := p.parse(data)
	if err != nil {
		return "", err
	}
	var l int64
	c := CPU()
	c.SingleStep(prog, func(r constant.Value) {
		n, ok := constant.Int64Val(r)
		if !ok {
			panic("not an int64")
		}
		if n > l {
			l = n
		}
	})
	return fmt.Sprintf("%d", l), nil
}

func (problem) parse(data []byte) (program, error) {
	var p program
	for _, line := range bytes.Split(data, []byte("\n")) {
		if len(line) == 0 {
			continue
		}

		parts := bytes.Split(line, []byte(" "))

		p = append(p, struct {
			Instruction instruction
			Condition   instruction
		}{
			Instruction: Instruction(parts[0], parts[1], parts[2], ops),
			Condition:   Instruction(parts[4], parts[5], parts[6], conds),
		})
	}
	return p, nil
}
