package day08

import "go/constant"

type cpu struct {
	regs map[constant.Value]constant.Value
}

func CPU() *cpu {
	return &cpu{
		regs: make(map[constant.Value]constant.Value),
	}
}

func (c *cpu) Evaluate(i instruction) bool {
	return constant.Compare(valueWithDefault(c.regs[i.Reg]), i.Op, i.N)
}

func (c *cpu) Execute(i instruction) {
	c.regs[i.Reg] = constant.BinaryOp(valueWithDefault(c.regs[i.Reg]), i.Op, i.N)
}

func (c cpu) Registers() map[constant.Value]constant.Value {
	return c.regs
}

func (c *cpu) Run(p program) {
	c.SingleStep(p, func(r constant.Value) {})
}

func (c *cpu) SingleStep(p program, cb func(r constant.Value)) {
	for _, pc := range p {
		if c.Evaluate(pc.Condition) {
			c.Execute(pc.Instruction)
			cb(c.regs[pc.Instruction.Reg])
		}
	}
}

func valueWithDefault(v constant.Value) constant.Value {
	if v == nil {
		return constant.MakeInt64(0)
	}
	return v
}
