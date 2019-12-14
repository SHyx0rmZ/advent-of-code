package day05

import (
	"fmt"
	"io"
	"strconv"

	"github.com/SHyx0rmZ/advent-of-code/year2019/intcode"
)

type problem struct{}

func Problem() *problem {
	return &problem{}
}

var tens = [...]int{100, 1000, 10000}

func (p problem) PartOneWithReader(r io.Reader) (string, error) {
	ns, err := intcode.NewProgram(r)
	if err != nil {
		return "", err
	}
	input := make(chan int, 1)
	input <- 1
	output := make(chan int, 1)
	var tests bool
	for i := 0; i < len(ns) && ns[i] != 99; i++ {
		var ins intcode.Instruction
		switch ns[i] % 100 {
		case 1:
			ins = &intcode.Add{}
		case 2:
			ins = &intcode.Mul{}
		case 3:
			ins = &intcode.In{}
		case 4:
			ins = &intcode.Out{}
		case 99:
			ins = &intcode.Hlt{}
		default:
			panic("op: " + strconv.Itoa(ns[i]))
		}
		var op []intcode.Operand
		for p := 0; p < intcode.Args[ns[i]%100]; p++ {
			if (ns[i]%(tens[p]*10))/tens[p] == 1 {
				op = append(op, &intcode.Immediate{Value: ns[i+1+p]})
			} else {
				op = append(op, &intcode.Register{ID: ns[i+1+p], Program: ns})
			}
		}
		i += intcode.Args[ns[i]%100]
		switch ins := ins.(type) {
		case *intcode.In:
			ins.Src = input
		case *intcode.Out:
			ins.Dst = output
		}
		ins.Execute(op...)
		select {
		case v := <-output:
			fmt.Println("out", v)
			if v == 0 {
				tests = true
			} else {
				if !tests {
					panic("no tests")
				}
				if ns[i+1] == 99 {
					return strconv.Itoa(v), nil
				}
				panic("test failed")
			}
		default:
		}
	}
	panic("no result")
}

func (p problem) PartTwoWithReader(r io.Reader) (string, error) {
	ns, err := intcode.NewProgram(r)
	if err != nil {
		return "", err
	}
	input := make(chan int, 1)
	input <- 5
	output := make(chan int, 1)
	for i := 0; i < len(ns) && ns[i] != 99; i++ {
		var ins intcode.Instruction
		switch ns[i] % 100 {
		case 1:
			ins = &intcode.Add{}
		case 2:
			ins = &intcode.Mul{}
		case 3:
			ins = &intcode.In{}
		case 4:
			ins = &intcode.Out{}
		case 5:
			ins = &intcode.JumpIfTrue{}
		case 6:
			ins = &intcode.JumpIfFalse{}
		case 7:
			ins = &intcode.LessThan{}
		case 8:
			ins = &intcode.Equals{}
		case 99:
			ins = &intcode.Hlt{}
		default:
			panic("op: " + strconv.Itoa(ns[i]))
		}
		var op []intcode.Operand
		for p := 0; p < intcode.Args[ns[i]%100]; p++ {
			if (ns[i]%(tens[p]*10))/tens[p] == 1 {
				op = append(op, &intcode.Immediate{Value: ns[i+1+p]})
			} else {
				op = append(op, &intcode.Register{ID: ns[i+1+p], Program: ns})
			}
		}
		i += intcode.Args[ns[i]%100]
		switch ins := ins.(type) {
		case *intcode.In:
			ins.Src = input
		case *intcode.Out:
			ins.Dst = output
		case *intcode.JumpIfTrue:
			ins.IP = &i
		case *intcode.JumpIfFalse:
			ins.IP = &i
		}
		ins.Execute(op...)
		select {
		case v := <-output:
			fmt.Println("out", v)
			if v == 0 {
				panic("test failed")
			} else {
				if ns[i+1] == 99 {
					return strconv.Itoa(v), nil
				}
				panic("test failed")
			}
		default:
		}
	}
	panic("no result")
}
