package intcode

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
