package day16

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"math/bits"
	"strconv"

	"github.com/SHyx0rmZ/advent-of-code"
)

type problem struct {
}

func Problem() aoc.ReaderAwareProblem {
	return &problem{}
}

type access interface {
	Read(*computer) int
	Write(*computer, int)
}

type immediateAccess struct {
	Value int
	Type  accessType
}

type registerAccess struct {
	Index int
	Type  accessType
}

func (a immediateAccess) Read(c *computer) int { return a.Value }
func (a immediateAccess) Write(*computer, int) {}

func (a registerAccess) Read(c *computer) int     { return c.Registers[a.Index] }
func (a registerAccess) Write(c *computer, v int) { c.Registers[a.Index] = v }

type accessType int

const (
	___unused accessType = iota
	_register
	immediate
	__ignored
)

func (t accessType) String() string {
	switch t {
	case _register:
		return "R"
	case immediate:
		return "I"
	default:
		return "-"
	}
}

func add(c *computer, p ParamedOpCode) { p.Destination.Write(c, p.Source1.Read(c)+p.Source2.Read(c)) }
func mul(c *computer, p ParamedOpCode) { p.Destination.Write(c, p.Source1.Read(c)*p.Source2.Read(c)) }
func and(c *computer, p ParamedOpCode) { p.Destination.Write(c, p.Source1.Read(c)&p.Source2.Read(c)) }
func _or(c *computer, p ParamedOpCode) { p.Destination.Write(c, p.Source1.Read(c)|p.Source2.Read(c)) }
func mov(c *computer, p ParamedOpCode) { p.Destination.Write(c, p.Source1.Read(c)) }
func tgt(c *computer, p ParamedOpCode) {
	var v int
	if p.Source1.Read(c) > p.Source2.Read(c) {
		v = 1
	}
	p.Destination.Write(c, v)
}
func teq(c *computer, p ParamedOpCode) {
	var v int
	if p.Source1.Read(c) == p.Source2.Read(c) {
		v = 1
	}
	p.Destination.Write(c, v)
}

func lip(c *computer, p ParamedOpCode) { c.IP = p.Source1.Read(c) }

type OpCode struct {
	Mnemonic    string
	Operation   func(c *computer, p ParamedOpCode)
	Source1     accessType
	Source2     accessType
	Destination accessType
}

var opcodes = []OpCode{
	{"addr", add, _register, _register, _register},
	{"addi", add, _register, immediate, _register},
	{"mulr", mul, _register, _register, _register},
	{"muli", mul, _register, immediate, _register},
	{"banr", and, _register, _register, _register},
	{"bani", and, _register, immediate, _register},
	{"borr", _or, _register, _register, _register},
	{"bori", _or, _register, immediate, _register},
	{"setr", mov, _register, ___unused, _register},
	{"seti", mov, immediate, ___unused, _register},
	{"gtir", tgt, immediate, _register, _register},
	{"gtri", tgt, _register, immediate, _register},
	{"gtrr", tgt, _register, _register, _register},
	{"eqir", teq, immediate, _register, _register},
	{"eqri", teq, _register, immediate, _register},
	{"eqrr", teq, _register, _register, _register},
	//{"#ip", lip, immediate, __ignored, __ignored},
}

type ParamedOpCode struct {
	OpCode
	Source1     access
	Source2     access
	Destination access
}

func selectA(t accessType, chn <-chan int) access {
	switch t {
	default:
		return immediateAccess{0, __ignored}
	case ___unused:
		return immediateAccess{<-chn, ___unused}
	case _register:
		return registerAccess{<-chn, _register}
	case immediate:
		return immediateAccess{<-chn, immediate}
	}
}

func selectI(cho <-chan OpCode, chn <-chan int, chp chan<- ParamedOpCode) {
	for o := range cho {
		p := ParamedOpCode{
			o,
			selectA(o.Source1, chn),
			selectA(o.Source2, chn),
			selectA(o.Destination, chn),
		}
		chp <- p
	}
	close(chp)
}

type computer struct {
	Registers [6]int
	IP        int
	Program   []func(c *computer)
}

func buildP(p ParamedOpCode) func(c *computer) {
	return func(c *computer) {
		//fmt.Println(c.Registers[c.IP], p.Mnemonic, p.Source1, p.Source2, p.Destination)
		p.Operation(c, p)
	}
}

func buildC(chp <-chan ParamedOpCode, chc chan<- *computer, sv int) {
	c := new(computer)
	c.Registers[0] = sv
	for p := range chp {
		if p.Mnemonic[0] == '#' {
			p.Operation(c, p)
		} else {
			c.Program = append(c.Program, buildP(p))
		}
	}
	chc <- c
	close(chc)
}

func (p problem) PartOneWithReader(r io.Reader) (string, error) {
	bs, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}
	ps := bytes.SplitN(bs, []byte("\n\n\n\n"), 2)
	pps := bytes.Split(ps[0], []byte("\n\n"))
	var rs int
	for _, p := range pps {
		ps = bytes.SplitN(p, []byte("\n"), 3)
		var before []int
		var after []int
		var instruction []int
		for _, n := range bytes.Split(bytes.Split(bytes.Split(ps[0], []byte("Before: ["))[1], []byte("]"))[0], []byte(", ")) {
			cn, err := strconv.Atoi(string(n))
			if err != nil {
				return "", err
			}
			before = append(before, cn)
		}
		for _, n := range bytes.Split(ps[1], []byte(" ")) {
			cn, err := strconv.Atoi(string(n))
			if err != nil {
				return "", err
			}
			instruction = append(instruction, cn)
		}
		for _, n := range bytes.Split(bytes.Split(bytes.Split(ps[2], []byte("After:  ["))[1], []byte("]"))[0], []byte(", ")) {
			cn, err := strconv.Atoi(string(n))
			if err != nil {
				return "", err
			}
			after = append(after, cn)
		}
		var pc int
		for i := range opcodes {
			chn := make(chan int)
			cho := make(chan OpCode)
			chp := make(chan ParamedOpCode)
			chc := make(chan *computer)
			go selectI(cho, chn, chp)
			go buildC(chp, chc, 0)
			cho <- opcodes[i]
			chn <- instruction[1]
			chn <- instruction[2]
			chn <- instruction[3]
			close(cho)
			close(chn)
			c := <-chc
			for i := range before {
				c.Registers[i] = before[i]

			}
			c.Program[0](c)
			var invalid bool
			for i := range after {
				if c.Registers[i] != after[i] {
					invalid = true
					break
				}
			}
			if !invalid {
				pc++
			}
		}
		if pc >= 3 {
			rs++
		}
	}
	return strconv.Itoa(rs), nil
}

func (p problem) PartTwoWithReader(r io.Reader) (string, error) {
	bs, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}
	ps := bytes.SplitN(bs, []byte("\n\n\n\n"), 2)
	pps := bytes.Split(ps[0], []byte("\n\n"))
	var om []uint
	for _ = range opcodes {
		om = append(om, 0xffff)
	}
	for _, p := range pps {
		ps := bytes.SplitN(p, []byte("\n"), 3)
		var before []int
		var after []int
		var instruction []int
		for _, n := range bytes.Split(bytes.Split(bytes.Split(ps[0], []byte("Before: ["))[1], []byte("]"))[0], []byte(", ")) {
			cn, err := strconv.Atoi(string(n))
			if err != nil {
				return "", err
			}
			before = append(before, cn)
		}
		for _, n := range bytes.Split(ps[1], []byte(" ")) {
			cn, err := strconv.Atoi(string(n))
			if err != nil {
				return "", err
			}
			instruction = append(instruction, cn)
		}
		for _, n := range bytes.Split(bytes.Split(bytes.Split(ps[2], []byte("After:  ["))[1], []byte("]"))[0], []byte(", ")) {
			cn, err := strconv.Atoi(string(n))
			if err != nil {
				return "", err
			}
			after = append(after, cn)
		}
		var pm uint
		for i := range opcodes {
			chn := make(chan int)
			cho := make(chan OpCode)
			chp := make(chan ParamedOpCode)
			chc := make(chan *computer)
			go selectI(cho, chn, chp)
			go buildC(chp, chc, 0)
			cho <- opcodes[i]
			chn <- instruction[1]
			chn <- instruction[2]
			chn <- instruction[3]
			close(cho)
			close(chn)
			c := <-chc
			for i := range before {
				c.Registers[i] = before[i]

			}
			c.Program[0](c)
			var invalid bool
			for i := range after {
				if c.Registers[i] != after[i] {
					invalid = true
					break
				}
			}
			if !invalid {
				pm |= 1 << uint(i)
			}
		}
		om[instruction[0]] &= pm
	}
	for {
		var reduced bool
		for i := range om {
			if bits.OnesCount(om[i]) != 1 {
				reduced = true
				for j := range om {
					if i == j {
						continue
					}
					if bits.OnesCount(om[j]) == 1 {
						om[i] &= ^om[j]
					}
				}
			}
		}
		if !reduced {
			break
		}
	}
	fmt.Println()
	chn := make(chan int)
	cho := make(chan OpCode)
	chp := make(chan ParamedOpCode)
	chc := make(chan *computer)
	go selectI(cho, chn, chp)
	go buildC(chp, chc, 0)
	for _, l := range bytes.Split(ps[1], []byte("\n")) {
		if len(l) == 0 {
			continue
		}
		is := bytes.Split(l, []byte(" "))
		var instruction []int
		for _, n := range is {
			cn, err := strconv.Atoi(string(n))
			if err != nil {
				return "", err
			}
			instruction = append(instruction, cn)
		}
		cho <- opcodes[bits.TrailingZeros(om[instruction[0]])]
		chn <- instruction[1]
		chn <- instruction[2]
		chn <- instruction[3]
	}
	close(cho)
	close(chn)
	c := <-chc
	c.IP = 4
	for c.Registers[c.IP] >= 0 && c.Registers[c.IP] < len(c.Program) {
		c.Program[c.Registers[c.IP]](c)
		c.Registers[c.IP] += 1
	}
	fmt.Println(c)
	return strconv.Itoa(0), nil
}

func (p problem) shared(r io.Reader, sv int) (string, error) {
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanWords)
	chn := make(chan int)
	cho := make(chan OpCode)
	chp := make(chan ParamedOpCode)
	chc := make(chan *computer)
	go selectI(cho, chn, chp)
	go buildC(chp, chc, sv)
scan:
	for s.Scan() {
		for _, op := range opcodes {
			if s.Text() == op.Mnemonic {
				cho <- op
				continue scan
			}
		}
		i, err := strconv.Atoi(s.Text())
		if err != nil {
			panic(err)
		}
		chn <- i
	}
	close(cho)
	close(chn)
	if err := s.Err(); err != nil {
		return "", err
	}
	c := <-chc
	for c.Registers[c.IP] >= 0 && c.Registers[c.IP] < len(c.Program) {
		c.Program[c.Registers[c.IP]](c)
		c.Registers[c.IP] += 1
	}
	fmt.Println(c)
	return "", nil
}
