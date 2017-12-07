package day05

import (
	"bytes"
	"fmt"
	"os"

	"github.com/SHyx0rmZ/advent-of-code/input"
)

func Command() error {
	if len(os.Args) < 4 {
		panic("not enough arguments")
	}

	c, err := input.ReadInput(os.Args[3])
	if err != nil {
		return err
	}
	c = bytes.TrimSpace(c)

	var steps int

	switch os.Args[2] {
	case "strange":
		steps, err = JumpStrange(string(c))
	case "evenstranger":
		steps, err = JumpEvenStranger(string(c))
	default:
		panic("unknown sub-command: " + os.Args[2])
	}

	if err != nil {
		return err
	}

	_, err = fmt.Printf("%d\n", steps)

	return err
}
