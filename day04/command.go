package day04

import (
	"strings"
	"os"
	"io/ioutil"
	"bytes"
	"fmt"
)

func Command() error {
	if len(os.Args) < 4 {
		panic("not enough arguments")
	}

	var f *os.File
	var err error

	if os.Args[3] == "-" {
		f = os.Stdin
	} else {
		f, err = os.OpenFile(os.Args[3], os.O_RDONLY, 0755)
		if err != nil {
			return err
		}
		defer f.Close()
	}

	c, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}
	c = bytes.TrimSpace(c)

	var sum int

	switch os.Args[2] {
	case "count":
		for _, passphrase := range bytes.Split(c, []byte("\n")) {
			if Valid(string(passphrase)) {
				sum++
			}
		}
	default:
		panic("unknown sub-command: " + os.Args[2])
	}

	if err != nil {
		return err
	}

	_, err = fmt.Printf("%d\n", sum)

	return err
}

func Valid(p string) bool {
	set := make(map[string]struct{})
	for _, word := range strings.Split(p, " ") {
		_, ok := set[word]
		if ok {
			return false
		}
		set[word] = struct{}{}
	}
	return true
}
