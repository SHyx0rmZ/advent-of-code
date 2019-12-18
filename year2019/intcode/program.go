package intcode

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
)

var tens = [...]int{100, 1000, 10000}

type Program map[int]int

func NewProgram(r io.Reader) (Program, error) {
	p := make(Program)
	bs, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	for i, d := range bytes.Split(bytes.TrimSpace(bs), []byte{','}) {
		n, err := strconv.Atoi(string(d))
		if err != nil {
			return nil, err
		}
		p[i] = n
	}
	return p, nil
}

func (p Program) Copy() Program {
	ns := make(Program, len(p))
	for k, v := range p {
		ns[k] = v
	}
	return ns
}

func (p Program) Run(input <-chan int, output chan<- int) {
	p.Copy().RunInPlace(input, output)
}

func (p Program) RunInPlace(input <-chan int, output chan<- int) {
	defer close(output)
	for i := 0; i < len(p); i++ {
		var ins Instruction
		switch p[i] % 100 {
		case 1:
			ins = &Add{}
		case 2:
			ins = &Mul{}
		case 3:
			ins = &In{}
		case 4:
			ins = &Out{}
		case 5:
			ins = &JumpIfTrue{}
		case 6:
			ins = &JumpIfFalse{}
		case 7:
			ins = &LessThan{}
		case 8:
			ins = &Equals{}
		case 99:
			ins = &Hlt{}
		default:
			panic("op: " + strconv.Itoa(p[i]))
		}
		var op []Operand
		for a := 0; a < Args[p[i]%100]; a++ {
			if (p[i]%(tens[a]*10))/tens[a] == 1 {
				op = append(op, &Immediate{Value: p[i+1+a]})
			} else {
				op = append(op, &Register{ID: p[i+1+a], Program: p})
			}
		}
		i += Args[p[i]%100]
		switch ins := ins.(type) {
		case *In:
			ins.Src = input
		case *Out:
			ins.Dst = output
		case *JumpIfTrue:
			ins.IP = &i
		case *JumpIfFalse:
			ins.IP = &i
		case *Hlt:
			return
		}
		ins.Execute(op...)
	}
	panic("no result")
}

func (p Program) String() string {
	var sb strings.Builder
	params1 := func(i int) int {
		sb.WriteByte(' ')
		sb.WriteByte('(')
		sb.WriteString(strconv.Itoa(p[i+1]))
		sb.WriteByte(')')
		sb.WriteByte('\n')
		return i + 1
	}
	params3 := func(i int) int {
		sb.WriteByte(' ')
		if p[i]&I0 == I0 {
			sb.WriteByte('$')
		} else {
			sb.WriteByte('(')
		}
		sb.WriteString(strconv.Itoa(p[i+1]))
		if p[i]&I0 != I0 {
			sb.WriteByte(')')
		}
		sb.WriteByte(',')
		sb.WriteByte(' ')
		if p[i]&I1 == I1 {
			sb.WriteByte('$')
		} else {
			sb.WriteByte('(')
		}
		sb.WriteString(strconv.Itoa(p[i+2]))
		if p[i]&I1 != I1 {
			sb.WriteByte(')')
		}
		sb.WriteByte(',')
		sb.WriteByte(' ')
		sb.WriteString(strconv.Itoa(p[i+3]))
		sb.WriteByte('\n')
		return i + 3
	}
	for i := 0; i < len(p); i++ {
		sb.WriteString(fmt.Sprintf("%04d: ", i))
		switch p[i] & 99 {
		case ADD:
			sb.WriteString("ADD")
			i = params3(i)
		case MUL:
			sb.WriteString("MUL")
			i = params3(i)
		case IN:
			sb.WriteString("IN")
			i = params1(i)
		case OUT:
			sb.WriteString("OUT")
			i = params1(i)
		case HLT:
			sb.WriteString("HLT\n")
		default:
			sb.WriteString("DAT")
			sb.WriteByte(' ')
			sb.WriteByte('$')
			sb.WriteString(strconv.Itoa(p[i]))
			sb.WriteByte('\n')
		}
	}
	return sb.String()
}
