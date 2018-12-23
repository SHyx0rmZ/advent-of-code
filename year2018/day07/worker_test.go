package day07_test

import (
	"fmt"
	"github.com/SHyx0rmZ/advent-of-code/year2018/day07"
	"testing"
)

func TestQueue_Add(t *testing.T) {
	var called bool
	q := &day07.Queue{
		Size: 1,
	}
	q.Add(func() day07.Payload {
		called = true
		return nil
	})
	q.Process()
	if !called {
		t.Error("expected function to be called")
	}
}

func TestQueue_Remove(t *testing.T) {
	var is []int
	q := &day07.Queue{
		Size: 1,
	}
	var f func(int) day07.Payload
	f = func(i int) day07.Payload {
		if i >= 5 {
			return func() day07.Payload {
				is = append(is, i)
				return nil
			}
		}
		return func() day07.Payload {
			is = append(is, i)
			return f(i + 1)
		}
	}
	q.Add(f(0))
	for q.Process() {
	}
	fmt.Println(is)
}
