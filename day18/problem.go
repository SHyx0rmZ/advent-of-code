package day18

import (
	"bytes"
	"fmt"
	"math/rand"
	"reflect"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

const timeFactor = time.Microsecond

type problem struct{}

func Problem() *problem {
	return &problem{}
}

func (p problem) PartOne(data []byte) (string, error) {
	program, err := p.parse(data)
	if err != nil {
		return "", err
	}

	cpu := &CPU{
		Registers: make(map[Address]int),
		Program:   program,
	}

	checkpoint := setUpQueue(cpu, cpu)

	var r int
	ISA["snd"] = func(cpu *CPU, target, source Operand) {
		cpu.Registers['s']++
		address := Address(target.(register))
		r = cpu.Registers[address]
	}
	fmt.Printf("\033[2J")

	go func() {
		for {
			//instruction := cpu.Program[cpu.PC]
			//ISA[instruction.Mnemonic](cpu, instruction.Target, instruction.Source)
			//cpu.PC++

			instruction := program[cpu.PC]
			if instruction.Mnemonic == "snd" {
				cpu.state = sending
			} else if instruction.Mnemonic != "rcv" {
				cpu.state = busy
			}
			if instruction.Mnemonic == "rcv" || instruction.Mnemonic == "jgz" {
				if instruction.Mnemonic == "rcv" {
					cpu.state = waiting
				}
				cpu.Render(cpu.ID*61+7, 2)
				time.Sleep(time.Duration(rand.Intn(200)+300) * timeFactor)
			}
			ISA[instruction.Mnemonic](cpu, instruction.Target, instruction.Source)
			if instruction.Mnemonic != "rcv" && (instruction.Mnemonic != "jgz" || instruction.Target.Value(cpu) != 0) {
				cpu.Render(cpu.ID*61+7, 2)
				time.Sleep(time.Duration(rand.Intn(200)+300) * timeFactor)
			}
			cpu.Jump(2)
		}
	}()

	checkpoint.Deadlocked()
	checkpoint.Cancel()

	return fmt.Sprintf("%d", r), nil
}

func setUpQueue(sender, receiver *CPU) checkpoint {
	q := &queue{
		Values:     nil,
		Continue:   make(chan struct{}),
		Cancel:     make(chan struct{}),
		Deadlocked: make(chan bool),
	}
	sender.Sender = q
	receiver.Receiver = q
	receiver.Deadlock = new(atomic.Value)
	receiver.Deadlock.Store(false)
	return checkpoint{
		chanDeadlocked: q.Deadlocked,
		chanContinue:   q.Continue,
		chanCancel:     q.Cancel,
		cpu:            receiver.Deadlock,
		C:              receiver,
	}
}

func deadlocked(cpus []*CPU) bool {
	for _, cpu := range cpus {
		if cpu.state != waiting {
			return false
		}
	}
	fmt.Println("opajuws")
	return true
}

func (p problem) PartTwo(data []byte) (string, error) {
	program, err := p.parse(data)
	if err != nil {
		return "", err
	}

	cpus := []*CPU{
		{Registers: map[Address]int{'p': 0}, ID: 0},
		{Registers: map[Address]int{'p': 1}, ID: 1},
	}

	checkpoints := []checkpoint{
		setUpQueue(cpus[0], cpus[1]),
		setUpQueue(cpus[1], cpus[0]),
	}

	var r int

	ISA["snd"] = func(cpu *CPU, target, source Operand) {
		if cpu == cpus[1] {
			r++
		}
		cpu.Registers['s']++
		cpu.Send(target.Value(cpu))
	}

	fmt.Printf("\033[2J")
	cpus[0].Program = program
	cpus[0].Render(7, 2)
	cpus[1].Program = program
	cpus[1].Render(7+61, 2)

	wg := &sync.WaitGroup{}

	for i, cpu := range cpus {
		wg.Add(1)

		go func(cpu *CPU, other *CPU) {
			defer wg.Done()
			defer func() { cpu.Deadlock.Store(true) }()

			for cpu.PC < len(program) && !cpu.Deadlock.Load().(bool) {
				instruction := program[cpu.PC]
				if instruction.Mnemonic == "snd" {
					cpu.state = sending
				} else if instruction.Mnemonic != "rcv" {
					cpu.state = busy
				}
				//fmt.Printf("%s\n", instruction.Mnemonic)
				if instruction.Mnemonic == "rcv" || instruction.Mnemonic == "jgz" {
					if instruction.Mnemonic == "rcv" {
						cpu.state = waiting
					}
					cpu.Render(cpu.ID*61+7, 2)
					if other.Deadlock.Load().(bool) {
						time.Sleep(time.Duration(rand.Intn(50)+75) * timeFactor)
					} else {
						time.Sleep(time.Duration(rand.Intn(200)+300) * timeFactor)
					}
				}
				ISA[instruction.Mnemonic](cpu, instruction.Target, instruction.Source)
				if instruction.Mnemonic != "rcv" && (instruction.Mnemonic != "jgz" || instruction.Target.Value(cpu) != 0) {
					cpu.Render(cpu.ID*61+7, 2)
					if other.Deadlock.Load().(bool) {
						time.Sleep(time.Duration(rand.Intn(50)+75) * timeFactor)
					} else {
						time.Sleep(time.Duration(rand.Intn(200)+300) * timeFactor)
					}
				}
				cpu.Jump(2)

			}
		}(cpu, cpus[(i+1)%len(cpus)])
	}

	//for len(checkpoints) > 0 {
	//fmt.Printf("%#v\n", len(checkpoints))
	var cases []reflect.SelectCase
	for i := range checkpoints {
		cases = append(cases, reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(checkpoints[i].chanDeadlocked),
		})
	}

	go func() {
		for {
			i, _, _ := reflect.Select(cases)
			//fmt.Printf("%d %d %t %t\n", len(checkpoints), i, v.Bool(), ok)
			//checkpoints[i].cpu.Store(v.Bool())
			if deadlocked(cpus) {
				break
				//checkpoints[i].Cancel()
				//checkpoints = append(checkpoints[:i], checkpoints[i+1:]...)
				//continue
			}
			checkpoints[i].Continue()
		}
	}()

	//fmt.Println(r)

	wg.Wait()

	return fmt.Sprintf("%d", r), nil
}

func (problem) parse(data []byte) (Program, error) {
	var es Program
	for _, line := range bytes.Split(data, []byte("\n")) {
		if len(line) == 0 {
			continue
		}
		ps := bytes.Split(line, []byte(" "))
		target := operandFromBytes(ps[1])
		var source Operand
		if len(ps) > 2 {
			source = operandFromBytes(ps[2])
		}
		es = append(es, Instruction{
			Mnemonic: string(ps[0]),
			Target:   target,
			Source:   source,
		})
	}
	return es, nil
}

func operandFromBytes(bs []byte) Operand {
	i, err := strconv.Atoi(string(bs))
	if err != nil {
		return register(bs[0])
	}
	return immediate(i)
}
