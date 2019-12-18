package intcode

type AdjustRelativeBase struct {
	Base *int
}

func (a AdjustRelativeBase) Execute(args ...Operand) {
	if len(args) < 1 {
		panic("length")
	}
	*a.Base += args[0].Read()
}
