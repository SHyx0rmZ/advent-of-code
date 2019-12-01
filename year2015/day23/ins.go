package day23

type in struct {
	Mnemonic string
	X, Y     op
	Func     func(c *cpu, x, y arg)
}

var ins = []in{
	{"hlf", op{reg, nil}, op{___, nil}, (*cpu).hlr},
	{"tpl", op{reg, nil}, op{___, nil}, (*cpu).tpl},
	{"inc", op{reg, nil}, op{___, nil}, (*cpu).inc},
	{"jmp", op{off, nil}, op{___, nil}, (*cpu).jmp},
	{"jie", op{reg, nil}, op{off, nil}, (*cpu).jie},
	{"jio", op{reg, nil}, op{off, nil}, (*cpu).jio},
}
