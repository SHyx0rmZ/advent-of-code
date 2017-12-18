package day18

import (
	"bytes"
	"fmt"
	"strconv"
)

type problem struct{}

func Problem() *problem {
	return &problem{}
}

type Ins interface {
	Execute(c *CPU)
}

type Ins1 struct {
	X byte
}

type Ins2 struct {
	X  byte
	Yr *byte
	Yi *int
}

type CPU struct {
	Register map[byte]int
	Result   chan int

	sound int
	done  bool
	pc    int
}

func (c *CPU) Execute(is []Ins) int {
	go func() {
		for {
			i := is[c.pc]
			i.Execute(c)
			if c.done {
				return
			}
			c.pc++
		}
	}()
	return <-c.Result
}

type SND Ins1
type SET Ins2
type ADD Ins2
type MUL Ins2
type MOD Ins2
type RCV Ins1
type JGZ Ins2

func (i SND) Execute(c *CPU) {
	c.sound = c.Register[i.X]
}

func (i SET) Execute(c *CPU) {
	var v int
	if i.Yi != nil {
		v = *i.Yi
	} else {
		v = c.Register[*i.Yr]
	}
	c.Register[i.X] = v
}

func (i ADD) Execute(c *CPU) {
	var v int
	if i.Yi != nil {
		v = *i.Yi
	} else {
		v = c.Register[*i.Yr]
	}
	c.Register[i.X] += v
}

func (i MUL) Execute(c *CPU) {
	var v int
	if i.Yi != nil {
		v = *i.Yi
	} else {
		v = c.Register[*i.Yr]
	}
	c.Register[i.X] *= v
}

func (i MOD) Execute(c *CPU) {
	var v int
	if i.Yi != nil {
		v = *i.Yi
	} else {
		v = c.Register[*i.Yr]
	}
	c.Register[i.X] = c.Register[i.X] % v
}

func (i RCV) Execute(c *CPU) {
	if c.Register[i.X] != 0 {
		c.Result <- c.sound
		c.done = true
	}
}

func (i JGZ) Execute(c *CPU) {
	var v int
	if i.Yi != nil {
		v = *i.Yi
	} else {
		v = c.Register[*i.Yr]
	}
	if c.Register[i.X] > 0 {
		c.pc += v - 1
	}
}

func (p problem) PartOne(data []byte) (string, error) {
	ins, err := p.parse(data)
	if err != nil {
		return "", err
	}

	c := &CPU{
		Register: make(map[byte]int),
		Result:   make(chan int),
	}

	r := c.Execute(ins)

	return fmt.Sprintf("%d", r), nil
}

func (p problem) PartTwo(data []byte) (string, error) {
	_, err := p.parse(data)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%d", 0), nil
}

func (problem) parse(data []byte) ([]Ins, error) {
	var es []Ins
	for _, line := range bytes.Split(data, []byte("\n")) {
		if len(line) == 0 {
			continue
		}
		ps := bytes.Split(line, []byte(" "))
		i := new(int)
		if len(ps) > 2 {
			var err error
			*i, err = strconv.Atoi(string(ps[2]))
			if err != nil {
				i = nil
			}
		}
		switch string(ps[0]) {
		case "snd":
			es = append(es, SND{ps[1][0]})
		case "set":
			es = append(es, SET{ps[1][0], &ps[2][0], i})
		case "add":
			es = append(es, ADD{ps[1][0], &ps[2][0], i})
		case "mul":
			es = append(es, MUL{ps[1][0], &ps[2][0], i})
		case "mod":
			es = append(es, MOD{ps[1][0], &ps[2][0], i})
		case "rcv":
			es = append(es, RCV{ps[1][0]})
		case "jgz":
			es = append(es, JGZ{ps[1][0], &ps[2][0], i})
		}
	}
	return es, nil
}
