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

type Add struct{}

func (Add) Execute(args ...Operand) {
	if len(args) < 3 {
		panic("length")
	}
	args[2].Write(args[0].Read() + args[1].Read())
}

type Mul struct{}

func (Mul) Execute(args ...Operand) {
	if len(args) < 3 {
		panic("length")
	}
	args[2].Write(args[0].Read() * args[1].Read())
}

type In struct {
	Src <-chan int
}

func (i In) Execute(args ...Operand) {
	if len(args) < 1 {
		panic("length")
	}
	args[0].Write(<-i.Src)
}

type Out struct {
	Dst chan<- int
}

func (o Out) Execute(args ...Operand) {
	if len(args) < 1 {
		panic("length")
	}
	o.Dst <- args[0].Read()
}

type JumpIfTrue struct {
	IP *int
}

func (j JumpIfTrue) Execute(args ...Operand) {
	if len(args) < 2 {
		panic("length")
	}
	if args[0].Read() != 0 {
		*j.IP = args[1].Read() - 1
	}
}

type JumpIfFalse struct {
	IP *int
}

func (j JumpIfFalse) Execute(args ...Operand) {
	if len(args) < 2 {
		panic("length")
	}
	if args[0].Read() == 0 {
		*j.IP = args[1].Read() - 1
	}
}

type LessThan struct{}

func (LessThan) Execute(args ...Operand) {
	if len(args) < 3 {
		panic("length")
	}
	if args[0].Read() < args[1].Read() {
		args[2].Write(1)
	} else {
		args[2].Write(0)
	}
}

type Equals struct{}

func (Equals) Execute(args ...Operand) {
	if len(args) < 3 {
		panic("length")
	}
	if args[0].Read() == args[1].Read() {
		args[2].Write(1)
	} else {
		args[2].Write(0)
	}
}

type Hlt struct{}

func (Hlt) Execute(args ...Operand) {
	panic(nil)
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
	Program []int
}

func (r Register) Read() int   { return r.Program[r.ID] }
func (r Register) Write(v int) { r.Program[r.ID] = v }

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
