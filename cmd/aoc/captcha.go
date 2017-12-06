package main

import (
	"bytes"
	"container/ring"
	"fmt"
	"io/ioutil"
	"os"
)

type captchaState struct {
	Last byte
	Sum int
}

func (s *captchaState) count(v interface{}) {
	b := v.(byte)
	if s.Last == b {
		s.Sum += int(b - '0')
	}

	s.Last = b
}

type captchaStateHalf struct {
	Ring *Ring
	Last byte
	Sum int
}

func newCaptchaStateHalf(r *ring.Ring) *captchaStateHalf {
	return &captchaStateHalf{
		Ring: (*Ring)(r),
	}
}

func (s *captchaStateHalf) Count() int {
	s.Ring.Do(s.count)
	return s.Sum
}

func (s *captchaStateHalf) count(r *ring.Ring) {
	n := (*Ring)(r).Advance(r.Len() / 2)
	if n.Value.(byte) == r.Value.(byte) {
		s.Sum += int(r.Value.(byte) - '0')
	}
}

type Ring ring.Ring

func (r *Ring) Do(f func(*ring.Ring)) {
	x := (*ring.Ring)(r)
	if x != nil {
		f(x)
		for p := x.Next(); p != x; p = p.Next() {
			f(p)
		}
	}
}

func (r *Ring) Advance(n int) *Ring {
	if r == nil {
		return r
	}

	x := (*ring.Ring)(r)

	for i := 0; i < n; i++ {
		x = x.Next()
	}

	return (*Ring)(x)
}

func Captcha(c string) int {
	r := ring.New(len(c))

	for i := 0; i < len(c); i++ {
		r.Value = c[i]
		r = r.Next()
	}

	s := captchaState{
		Last: r.Prev().Value.(byte),
	}
	r.Do(s.count)

	return s.Sum
}

func CaptchaHalf(c string) int {
	r := ring.New(len(c))

	for i := 0; i < len(c); i++ {
		r.Value = c[i]
		r = r.Next()
	}

	s := newCaptchaStateHalf(r)

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