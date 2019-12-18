package intcode

type instruction int

const (
	ADD = 1
	MUL = 2
	IN  = 3
	OUT = 4
	HLT = 99
)

const (
	R0 = 0
	R1 = 0
	I0 = 100
	I1 = 1000
)

type Instruction interface {
	Execute(args ...Operand)
}

type Operand interface {
	Read() int
	Write(int)
}

type Immediate struct {
	Value int
}

func (i Immediate) Read() int    { return i.Value }
func (i *Immediate) Write(v int) { i.Value = v }

type Register struct {
	ID      int
	Program Program
}

func (r Register) Read() int   { return r.Program[r.ID] }
func (r Register) Write(v int) { r.Program[r.ID] = v }

type Relative struct {
	Base    *int
	Offset  int
	Program Program
}

func (r Relative) Read() int   { return r.Program[*r.Base+r.Offset] }
func (r Relative) Write(v int) { r.Program[*r.Base+r.Offset] = v }

const (
	ADDRR instruction = ADD | R0 | R1
	ADDRI instruction = ADD | R0 | I1
	ADDIR instruction = ADD | I0 | R1
	ADDII instruction = ADD | I0 | I1
	MULRR instruction = MUL | R0 | R1
	MULRI instruction = MUL | R0 | I1
	MULIR instruction = MUL | I0 | R1
	MULII instruction = MUL | I0 | I1
	INR   instruction = IN | R0
	OUTR  instruction = OUT | R0
)

var Args = map[int]int{
	1:  3,
	2:  3,
	3:  1,
	4:  1,
	5:  2,
	6:  2,
	7:  3,
	8:  3,
	9:  1,
	99: 0,
}

var instructions = [...]instruction{
	ADDRR,
	ADDRI,
	ADDIR,
	ADDII,
	MULRR,
	MULRI,
	MULIR,
	MULII,
	INR,
	OUTR,
	HLT,
}
