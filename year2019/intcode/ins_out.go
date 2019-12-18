package intcode

type Out struct {
	Dst chan<- int
}

func (o Out) Execute(args ...Operand) {
	if len(args) < 1 {
		panic("length")
	}
	o.Dst <- args[0].Read()
}
