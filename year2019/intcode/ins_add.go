package intcode

type Add struct{}

func (Add) Execute(args ...Operand) {
	if len(args) < 3 {
		panic("length")
	}
	args[2].Write(args[0].Read() + args[1].Read())
}
