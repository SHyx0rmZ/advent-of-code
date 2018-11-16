package day01

func CaptchaNext(c string) int {
	s := state{
		Ring: NewRing(c),
		Func: (*state).next,
	}

	return s.Count()
}

func CaptchaHalf(c string) int {
	s := state{
		Ring: NewRing(c),
		Func: (*state).half,
	}

	return s.Count()
}
