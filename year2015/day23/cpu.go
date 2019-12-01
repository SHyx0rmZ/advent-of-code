package day23

type cpu struct {
	PC        int
	Registers map[register]int
}

func (c *cpu) run(p program) {
	for c.PC < len(p) {
		i := p[c.PC]
		i.Func(c, i.X.arg, i.Y.arg)
		c.PC++
	}
}

func (c *cpu) hlr(r, _ arg) {
	c.Registers[r.(register)] >>= 1
}

func (c *cpu) tpl(r, _ arg) {
	c.Registers[r.(register)] *= 3
}

func (c *cpu) inc(r, _ arg) {
	c.Registers[r.(register)]++
}

func (c *cpu) jmp(off, _ arg) {
	c.PC += int(off.(offset)) - 1
}

func (c *cpu) jie(r, off arg) {
	if c.Registers[r.(register)]%2 != 0 {
		return
	}
	c.PC += int(off.(offset)) - 1
}

func (c *cpu) jio(r, off arg) {
	if c.Registers[r.(register)] != 1 {
		return
	}
	c.PC += int(off.(offset)) - 1
}
