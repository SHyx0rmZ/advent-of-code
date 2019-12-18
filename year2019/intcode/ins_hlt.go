package intcode

type Hlt struct{}

func (Hlt) Execute(args ...Operand) {
	panic(nil)
}
