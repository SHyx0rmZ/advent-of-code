package intcode

type In struct {
	Src <-chan int
}

func (i In) Execute(args ...Operand) {
	if len(args) < 1 {
		panic("length")
	}
	args[0].Write(<-i.Src)
}
