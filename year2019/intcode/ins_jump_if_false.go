package intcode

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
