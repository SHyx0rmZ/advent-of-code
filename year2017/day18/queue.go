package day18

import (
	"runtime"
	"sync"
)

type Receiver interface {
	Receive(c *CPU) int
}

type Sender interface {
	Send(value int)
}

type queue struct {
	Values     []int
	Deadlocked chan bool
	Continue   chan struct{}
	Cancel     chan struct{}
	sync.Mutex
}

func (q *queue) Receive(c *CPU) int {
	//fmt.Printf("<- ")

	if len(q.Values) == 0 {
		//c.Deadlock.Store(true)
		q.Deadlocked <- true
		select {
		case <-q.Continue:
			//fmt.Printf("%p true continue\n", q)
		case <-q.Cancel:
			//fmt.Printf("%p true cancel\n", q)
			close(q.Deadlocked)
			close(q.Continue)
			close(q.Cancel)
			runtime.Goexit()
		}
		for len(q.Values) == 0 {
			runtime.Gosched()
		}
		//c.Deadlock = false
		q.Deadlocked <- false
		select {
		case <-q.Continue:
			//fmt.Printf("%p false continue\n", q)
		case <-q.Cancel:
			//fmt.Printf("%p false cancel\n", q)
			close(q.Deadlocked)
			close(q.Continue)
			close(q.Cancel)
			runtime.Goexit()
		}
	}

	q.Lock()
	defer q.Unlock()

	value := q.Values[0]
	q.Values = q.Values[1:]

	//fmt.Printf("%d\n", value)

	return value
}

func (q *queue) Send(value int) {
	q.Lock()
	defer q.Unlock()

	//fmt.Printf("-> %d\n", value)

	q.Values = append(q.Values, value)
}
