package day23

import (
	"strconv"
)

type arg interface {
	arg()
	String() string
}

type typ int

const (
	___ typ = iota
	reg
	off
)

type op struct {
	typ
	arg
}

type register string
type offset int

func (r register) arg() {}
func (o offset) arg()   {}

func (r register) String() string { return string(r) }
func (o offset) String() string   { return strconv.Itoa(int(o)) }
func (o op) String() string {
	if o.arg == nil {
		return "-"
	}
	return o.arg.String()
}
