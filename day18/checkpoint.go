package day18

import "sync/atomic"

type checkpoint struct {
	chanContinue   chan<- struct{}
	chanCancel     chan<- struct{}
	chanDeadlocked <-chan bool
	cpu            *atomic.Value
	C              *CPU
}

func (c *checkpoint) Cancel() {
	c.chanCancel <- struct{}{}
}

func (c *checkpoint) Continue() {
	c.chanContinue <- struct{}{}
}

func (c *checkpoint) Deadlocked() bool {
	return <-c.chanDeadlocked
}
