package day18

var ISA = map[string]func(cpu *CPU, target, source Operand){
	"snd": func(cpu *CPU, target, source Operand) {
		cpu.Registers['s']++
		cpu.Send(target.Value(cpu))
	},
	"set": func(cpu *CPU, target, source Operand) {
		cpu.Registers[target.(register).Address()] = source.Value(cpu)
	},
	"add": func(cpu *CPU, target, source Operand) {
		cpu.Registers[target.(register).Address()] += source.Value(cpu)
	},
	"mul": func(cpu *CPU, target, source Operand) {
		cpu.Registers[target.(register).Address()] *= source.Value(cpu)
	},
	"mod": func(cpu *CPU, target, source Operand) {
		cpu.Registers[target.(register).Address()] %= source.Value(cpu)
	},
	"rcv": func(cpu *CPU, target, source Operand) {
		cpu.Registers['r']++
		cpu.Registers[target.(register).Address()] = cpu.Receive(cpu)
	},
	"jgz": func(cpu *CPU, target, source Operand) {
		if target.Value(cpu) > 0 {
			cpu.Jump(source.Value(cpu))
		}
	},
}
