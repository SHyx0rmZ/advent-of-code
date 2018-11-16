package day18

type Register int
type Address byte

type Operand interface{ Value(*CPU) int }

type immediate int
type register Address

func (i immediate) Value(*CPU) int  { return int(i) }
func (r register) Value(c *CPU) int { return int(c.Registers[r.Address()]) }
func (r register) Address() Address { return Address(r) }

type Instruction struct {
	Mnemonic string
	Target   Operand
	Source   Operand
}

type Program []Instruction
