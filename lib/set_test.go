package lib_test

import (
	"github.com/SHyx0rmZ/advent-of-code/lib"
	"testing"
)

func TestSet_Add(t *testing.T) {
	s := lib.Set()
	s.Add(1)
}

func TestSet_Add2(t *testing.T) {
	s := lib.Set()
	s.Add(1)
	r := s.Contains(1)
	if !r {
		t.Errorf("got %t, want %t", r, true)
	}
}

func TestSet_Add3(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Error("expected panic")
		}
	}()

	s := lib.Set()
	s.Add(1)
	s.Add("")
}

func TestSet_Contains(t *testing.T) {
	s := lib.Set()
	r := s.Contains(1)
	if r {
		t.Errorf("got %t, want %t", r, false)
	}
}

func TestSet_Delete(t *testing.T) {
	s := lib.Set()
	s.Add(1)
	s.Delete(1)
	r := s.Contains(1)
	if r {
		t.Errorf("got %t, want %t", r, false)
	}
}

func TestSet_Delete2(t *testing.T) {
	s := lib.Set()
	s.Delete(1)
}

func TestSet_Elements(t *testing.T) {
	s := lib.Set()
	r := s.Elements()
	if len(r) != 0 {
		t.Errorf("got %#v, want %#v", r, []interface{}{})
	}
}

func TestSet_Elements2(t *testing.T) {
	s := lib.Set()
	s.Add(1)
	r := s.Elements()
	if len(r) != 1 || r[0] != 1 {
		t.Errorf("got %#v, want %#v", r, []interface{}{1})
	}
}
