package day03

var directions = []position{
	{X: +1, Y: +0},
	{X: +0, Y: +1},
	{X: -1, Y: +0},
	{X: +0, Y: -1},
}

func spiral() <-chan position {
	c := make(chan position)
	go func() {
		p := position{}
		g := 0
		for {
			for i := 0; i <= g/2; i++ {
				c <- p
				p = p.Add(directions[g%4])
			}
			g++
		}
	}()
	return c
}
