package day05

import (
	"bytes"
	"fmt"
	"github.com/SHyx0rmZ/advent-of-code/input"
	"os"
)

func Command() error {
	if len(os.Args) < 3 {
		panic("not enough arguments")
	}

	c, err := input.ReadInput(os.Args[2])
	if err != nil {
		return err
	}
	c = bytes.TrimSpace(c)

	steps, err := Jump(string(c))
	if err != nil {
		return err
	}

	_, err = fmt.Printf("%d\n", steps)

	return err
}
