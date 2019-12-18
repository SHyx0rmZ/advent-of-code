//+build ignore

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type instruction struct {
	Mnemonic  string
	Arguments int
	Code      []int
}

func main() {
	f, err := os.Open("microcode.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		a := strings.SplitN(s.Text(), ":", 2)
		b := strings.SplitN(a[0], ",", 2)
		c := strings.TrimSpace(a[1])
		parse(c)
		fmt.Printf("mnemonic: %s, args: %s, body: %s\n", strings.TrimSpace(b[0]), strings.TrimSpace(b[1]), c)
	}
	if err := s.Err(); err != nil {
		log.Fatalln(err)
	}
}

func parse(code string) {
	for _, expr := range strings.Split(code, ";") {
		var elemns []string
		for _, elem := range strings.Split(strings.TrimSpace(expr), " ") {
			elemns = append(elemns, strings.TrimSpace(elem))
		}
		fmt.Println(elemns)
		switch {
		case elemns[0] == "C":
		case elemns[0] == "PC":
		case strings.HasPrefix(elemns[0], "R"):
		default:
			panic("invalid")
		}
	}
}
