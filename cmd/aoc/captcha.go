package main

import "container/ring"

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