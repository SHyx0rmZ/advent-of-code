package day10

type counter struct {
	Chan chan struct{}
	N    int
}

func Counter(N int) *counter {
	return &counter{
		Chan: make(chan struct{}),
		N:    N,
	}
}

func (c *counter) Done() bool {
	_, ok := <-c.Chan
	return !ok
}

func (c *counter) Run() {
	for c.N > 0 {
		c.Chan <- struct{}{}
		c.N--
	}
	close(c.Chan)
}
