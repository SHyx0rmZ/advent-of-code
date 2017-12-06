package day02

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

	var sum int

	switch os.Args[2] {
	case "minmax":
		sum, err = ChecksumMinMax(c)
	case "division":
		sum, err = ChecksumDivision(c)
	default:
		panic("unknown sub-command: " + os.Args[2])
	}

	if err != nil {
		return err
	}

	_, err = fmt.Printf("%d\n", sum)

	return err
}
