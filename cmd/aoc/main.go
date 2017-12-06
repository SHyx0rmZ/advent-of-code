package main

import (
	"os"

	"github.com/SHyx0rmZ/advent-of-code/day01"
	"github.com/SHyx0rmZ/advent-of-code/day02"
	"github.com/SHyx0rmZ/advent-of-code/day04"
)

var commands = map[string]func() error{
	"captcha":     day01.Command,
	"checksum":    day02.Command,
	"passphrases": day04.Command,
}

func main() {
	if len(os.Args) < 2 {
		panic("no command specified")
	}

	c, ok := commands[os.Args[1]]
	if !ok {
		panic("unknown command: " + os.Args[1])
	}
	err := c()
	if err != nil {
		panic(err)
	}
}
