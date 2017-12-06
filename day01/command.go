package day01

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

	switch os.Args[2] {
	case "next":
		_, err = fmt.Printf("%d\n", CaptchaNext(string(c)))
	case "half":

		_, err = fmt.Printf("%d\n", CaptchaHalf(string(c)))
	default:
		panic("unknown sub-command: " + os.Args[2])
	}

	return err
}
