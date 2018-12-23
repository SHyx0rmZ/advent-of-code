package day07

import (
	"bufio"
	"io"
)

type rule struct {
	ID         rune
	Dependency rune
}

func ParseRules(r io.Reader) <-chan rule {
	chs := make(chan rule)
	go func() {
		chr := make(chan rune)
		defer close(chr)
		s := bufio.NewScanner(r)
		s.Split(bufio.ScanWords)
		go func() {
			defer close(chs)
			for {
				dep, ok := <-chr
				if !ok {
					break
				}
				id, ok := <-chr
				if !ok {
					panic("format error")
				}
				chs <- rule{id, dep}
			}
		}()
		for s.Scan() {
			switch s.Text() {
			case "Step", "must", "be", "finished", "before", "step", "can", "begin.":
			default:
				for _, ru := range s.Text() {
					chr <- ru
				}
			}
		}
	}()
	return chs
}
