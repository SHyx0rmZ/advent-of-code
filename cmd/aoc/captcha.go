package main

import (
	"bytes"
	"container/ring"
	"fmt"
	"io/ioutil"
	"os"
)

type captchaState struct {
	Ring *Ring
	Func func(*captchaState) func(*Ring)

	sum int
}

func (s *captchaState) Count() int {
	if s.Func != nil {
		s.Ring.Do(s.Func(s))
	}
	return s.sum
}

func (s *captchaState) half() func(*Ring) {
	return func(r *Ring) {
		c := r.Advance(r.Len() / 2)
		if c.Value == r.Value {
			s.sum += int(r.Value.(rune) - '0')
		}
	}
}

func (s *captchaState) next() func(*Ring) {
	return func(r *Ring) {
		c := r.Next()
		if c.Value == r.Value {
			s.sum += int(r.Value.(rune) - '0')
		}
	}
}

type Ring ring.Ring

func NewRing(s string) *Ring {
	r := ring.New(len(s))
	for _, c := range s {
		r.Value = c
		r = r.Next()
	}
	return (*Ring)(r)
}

func (r *Ring) Advance(n int) *Ring {
	p := r
	if p != nil {
		for i := 0; i < n; i++ {
			p = p.Next()
		}
	}
	return p
}

func (r *Ring) Do(f func(*Ring)) {
	if r != nil {
		f(r)
		for p := r.Next(); p != r; p = p.Next() {
			f(p)
		}
	}
}

func (r *Ring) Len() int {
	return (*ring.Ring)(r).Len()
}

func (r *Ring) Next() *Ring {
	return (*Ring)((*ring.Ring)(r).Next())
}

func (r *Ring) Prev() *Ring {
	return (*Ring)((*ring.Ring)(r).Prev())
}

func Captcha(c string) int {
	s := captchaState{
		Ring: NewRing(c),
		Func: (*captchaState).next,
	}

	return s.Count()
}

func CaptchaHalf(c string) int {
	s := captchaState{
		Ring: NewRing(c),
		Func: (*captchaState).half,
	}

	return s.Count()
}

func captchaCommand() error {
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

	switch os.Args[2] {
	case "next":
		_, err = fmt.Printf("%d\n", Captcha(string(c)))
	case "half":

		_, err = fmt.Printf("%d\n", CaptchaHalf(string(c)))
	default:
		panic("unknown sub-command: " + os.Args[2])
	}

	return err
}
