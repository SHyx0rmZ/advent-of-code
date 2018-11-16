package day18

import (
	"fmt"
	"math"
	"strings"
	"sync"
	"sync/atomic"
)

type state int

const (
	busy state = iota
	sending
	waiting
)

type CPU struct {
	Registers  map[Address]int
	Program    Program
	PC         int
	lPC        int
	Deadlock   *atomic.Value
	ID         int
	renderFunc func(*CPU, int, int)
	once       sync.Once
	state      state
	Sender
	Receiver
}

func (c *CPU) Jump(d int) {
	c.lPC = c.PC
	c.PC += d - 1
}

var mu sync.Mutex

func (c *CPU) Render(x, y int) {
	if c.renderFunc == nil {
		return
	}

	c.renderFunc(c, x, y)
}

func RenderCPU(c *CPU, x, y int) {
	mu.Lock()
	defer mu.Unlock()

	c.once.Do(func() {
		fmt.Printf("\033[%d;%dH╔══════════════════════╦════════════════════════════╗", y+1, x)
		fmt.Printf("\033[%d;%dH║       Registers      ║        Instructions        ║", y+2, x)
		fmt.Printf("\033[%d;%dH╟───┬──────────────────╫────┬───────────────────────╢", y+3, x)
		fmt.Printf("\033[%d;%dH║   │                  ║    │                       ║", y+4, x)
		fmt.Printf("\033[%d;%dH║   │                  ║    │                       ║", y+5, x)
		fmt.Printf("\033[%d;%dH║   │                  ║    │                       ║", y+6, x)
		fmt.Printf("\033[%d;%dH║   │                  ║    │                       ║", y+7, x)
		fmt.Printf("\033[%d;%dH║   │                  ║    │                       ║", y+8, x)
		fmt.Printf("\033[%d;%dH╟───┼──────────────────╢    │                       ║", y+9, x)
		fmt.Printf("\033[%d;%dH║ R │                  ║    │                       ║", y+10, x)
		fmt.Printf("\033[%d;%dH║ S │                  ║    │                       ║", y+11, x)
		fmt.Printf("\033[%d;%dH║ Q │                  ║    │                       ║", y+12, x)
		fmt.Printf("\033[%d;%dH╠═══╪══════════════════╣    │                       ║", y+13, x)
		fmt.Printf("\033[%d;%dH║ ∙ │                  ║    │                       ║", y+14, x)
		fmt.Printf("\033[%d;%dH╚═══╧══════════════════╩════╧═══════════════════════╝", y+15, x)

		fmt.Printf("\033[%d;%dHCPU #%d", y+14, x+11, c.ID)
	})

	for i, r := range []Address{'a', 'b', 'f', 'i', 'p'} {
		ins := c.Program[c.PC]

		reg := "\033[1;32m%s\033[0m"
		val := "\033[1;34m%16d\033[0m"

		if t, ok := ins.Target.(register); ok && Address(t) == r {
			switch ins.Mnemonic {
			case "jgz":
			case "snd":
			case "rcv":
				reg = "\033[5;1;37m%s\033[0m"
				val = "\033[8;1;37m%16d\033[0m"
			default:
				reg = "\033[1;37m%s\033[0m"
				val = "\033[1;37m%16d\033[0m"
			}
		}

		fmt.Printf("\033[%d;%dH%s", y+4+i, x+2, fmt.Sprintf(reg, strings.ToUpper(string(r))))
		fmt.Printf("\033[%d;%dH%s", y+4+i, x+6, fmt.Sprintf(val, c.Registers[r]))
	}

	l := 0
	for i := range c.Program {
		if math.Abs(float64(i)-float64(c.PC)) > 5 {
			if !((c.PC < 5 && i <= 10) || (c.PC > len(c.Program)-5) && i >= len(c.Program)-10) {
				continue
			}
		}

		ins := c.Program[i]

		fmt.Printf("\033[%d;%dH%02d", y+4+l, x+25, i+1)
		fmt.Printf("\033[%d;%dH   ", y+4+l, x+30)

		switch {
		case i == c.PC && c.PC == c.lPC:
			fallthrough
		case i == c.PC && c.PC == c.lPC+1:
			fmt.Printf("\033[%d;%dH > \033[1;37m", y+4+l, x+30)
		case i == c.PC && c.PC > c.lPC:
			fmt.Printf("\033[%d;%dH└> \033[1;37m", y+4+l, x+30)
		case i == c.PC && c.PC < c.lPC:
			fmt.Printf("\033[%d;%dH┌> \033[1;37m", y+4+l, x+30)
		case i == c.lPC && c.PC > c.lPC && c.PC != c.lPC+1:
			fmt.Printf("\033[%d;%dH┌─ ", y+4+l, x+30)
		case i == c.lPC && c.PC < c.lPC:
			fmt.Printf("\033[%d;%dH└─ ", y+4+l, x+30)
		case i > c.lPC && i < c.PC:
			fallthrough
		case i < c.lPC && i > c.PC:
			fmt.Printf("\033[%d;%dH│  ", y+4+l, x+30)
		}

		fmt.Printf("%s ", strings.ToUpper(ins.Mnemonic))

		reg := "\033[1;32m%6s\033[0m"
		imm := "\033[1;34m%6d\033[0m"

		if c.PC == i {
			reg = "\033[1;37m%6s\033[0m"
			imm = "\033[1;37m%6d\033[0m"
		}

		switch op := ins.Target.(type) {
		case register:
			fmt.Printf(reg, strings.ToUpper(string(op)))
		case immediate:
			fmt.Printf(imm, int(op))
		}

		if ins.Source != nil {
			fmt.Printf(", ")

			switch op := ins.Source.(type) {
			case register:
				fmt.Printf(reg, strings.ToUpper(string(op)))
			case immediate:
				fmt.Printf(imm, int(op))
			}
		} else {
			fmt.Printf("        ")
		}

		if c.PC == i {
			fmt.Printf("\033[0m")
		}

		l++
	}

	fmt.Printf("\033[%d;%dH%16d", y+10, x+6, c.Registers['r'])
	fmt.Printf("\033[%d;%dH%16d", y+11, x+6, c.Registers['s'])
	fmt.Printf("\033[%d;%dH%16d", y+12, x+6, len(c.Receiver.(*queue).Values))

	switch c.state {
	case busy:
		fmt.Printf("\033[%d;%dH \033[1;38m▓\033[0m", y+14, x+1)
	case sending:
		fmt.Printf("\033[%d;%dH \033[1;37m█\033[0m", y+14, x+1)
	case waiting:
		fmt.Printf("\033[%d;%dH \033[1;30m░\033[0m", y+14, x+1)
	}

	fmt.Printf("\033[%d;%dH", y+18, 1)
}
