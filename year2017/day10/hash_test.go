package day10

import (
	"testing"
)

func TestNode(t *testing.T) {
	h := newHash(8)
	n := h.nodes
	for i := 0; i < 8; i++ {
		if n.Value.(int) != i {
			t.Error("invalid Value")
		}
		if n.next == n {
			t.Error("invalid next pointer")
		}
		if n.next.prev != n {
			t.Error("invalid prev pointer on next")
		}
		n = n.next
	}
	for i := 0; i < 8; i++ {
		n = n.prev
		if n.Value.(int) != 7-i {
			t.Error("invalid Value")
		}
		if n.prev == n {
			t.Error("invalid prev pointer")
		}
		if n.prev.next != n {
			t.Error("invalid next pointer on prev")
		}
	}
}

func TestNode2(t *testing.T) {
	h := newHash(8)
	n := h.nodes
	ns := make([]*node, 8)
	for i := 0; i < 8; i++ {
		ns[i] = n
		n = n.next
	}
	//fmt.Println()
	h.Update([]byte{3})
	//for i := range ns {
	//	fmt.Printf("%d:  %p <- %p [ %d %5t ] -> %p\n", i, ns[i].prev, ns[i], ns[i].Value.(int), ns[i].toggle, ns[i].next)
	//}
	//for i := 0; i < 8; i++ {
	//	fmt.Printf(" %d", h.move(i).Value.(int))
	//}
	//fmt.Println()
	h.Update([]byte{3})
	//for i := range ns {
	//	fmt.Printf("%d:  %p <- %p [ %d %5t ] -> %p\n", i, ns[i].prev, ns[i], ns[i].Value.(int), ns[i].toggle, ns[i].next)
	//}
	if h.nodes != h.move(0) {
		t.Error("invalid start")
	}
	//for i := 0; i < 8; i++ {
	//	fmt.Printf(" %d", h.move(i).Value.(int))
	//}
	//fmt.Println()
}

func TestNode3(t *testing.T) {
	h := newHash(8)
	n := h.nodes
	ns := make([]*node, 8)
	for i := 0; i < 8; i++ {
		ns[i] = n
		n = n.next
	}
	//fmt.Println()
	h.Update([]byte{3})
	//for i := range ns {
	//	fmt.Printf("%d:  %p <- %p [ %d %5t ] -> %p\n", i, ns[i].prev, ns[i], ns[i].Value.(int), ns[i].toggle, ns[i].next)
	//}
	//for i := 0; i < 8; i++ {
	//	fmt.Printf(" %d", h.move(i).Value.(int))
	//}
	//fmt.Println()
	h.Update([]byte{6})
	//for i := range ns {
	//	fmt.Printf("%d:  %p <- %p [ %d %5t ] -> %p\n", i, ns[i].prev, ns[i], ns[i].Value.(int), ns[i].toggle, ns[i].next)
	//}
	if h.nodes != h.move(0) {
		t.Error("invalid start")
	}
	//for i := 0; i < 8; i++ {
	//	fmt.Printf(" %d", h.move(i).Value.(int))
	//}
	//fmt.Println()
}

func TestNode4(t *testing.T) {
	h := newHash(8)
	n := h.nodes
	ns := make([]*node, 8)
	for i := 0; i < 8; i++ {
		ns[i] = n
		n = n.next
	}
	//fmt.Println()
	//fmt.Println([]byte{3,6})
	//for i := range ns {
	//	fmt.Printf("%d:  %p <- %p [ %d %5t ] -> %p\n", i, ns[i].prev, ns[i], ns[i].Value.(int), ns[i].toggle, ns[i].next)
	//}
	h.Update([]byte{3, 6})
	//for i := range ns {
	//	fmt.Printf("%d:  %p <- %p [ %d %5t ] -> %p\n", i, ns[i].prev, ns[i], ns[i].Value.(int), ns[i].toggle, ns[i].next)
	//}
	//for i := 0; i < 8; i++ {
	//	fmt.Printf(" %d", h.move(i).Value.(int))
	//}
	//fmt.Println()
	if h.nodes != h.move(0) {
		t.Error("invalid start")
	}
	//for i := 0; i < 8; i++ {
	//	fmt.Printf(" %d", h.move(i).Value.(int))
	//}
	//fmt.Println()
}
