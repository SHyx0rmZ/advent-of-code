package main

import (
	"github.com/SHyx0rmZ/advent-of-code/day04"
	"os"
)

//go:generate go-bindata ../../data/...

var commands = map[string]func() error{
	"captcha":     captchaCommand,
	"checksum":    checksumCommand,
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
