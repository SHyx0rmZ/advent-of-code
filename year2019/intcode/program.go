package intcode

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
)

type program []int

func NewProgram(r io.Reader) (program, error) {
	var p program
	bs, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	for _, d := range bytes.Split(bytes.TrimSpace(bs), []byte{','}) {
		n, err := strconv.Atoi(string(d))
		if err != nil {
			return nil, err
		}
		p = append(p, n)
	}
	return p, nil
}

func (p program) String() string {
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
