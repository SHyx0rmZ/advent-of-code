package lib_test

import (
	"github.com/SHyx0rmZ/advent-of-code/pkg/lib"
	"testing"
)

func TestDict_Get(t *testing.T) {
	d := lib.Dict()
	_, ok := d.Get(1)
	if ok {
		t.Errorf("got %t, want %t", ok, false)
	}
}

func TestDict_Set(t *testing.T) {
	d := lib.Dict()
	d.Set(1, "a")
}

func TestDict_Set2(t *testing.T) {
	d := lib.Dict()
	d.Set(1, "a")
	v, ok := d.Get(1)
	if !ok || v != "a" {
		t.Errorf("got (%#v, %t), want (%#v, %t)", v, ok, "a", true)
	}
}

func TestDict_Set3(t *testing.T) {
	d := lib.Dict()
	d.Set(1, "a")
	d.Set(2, "b")
}

func TestDict_Set4(t *testing.T) {
	d := lib.Dict()
	d.Set(1, "a")
	d.Set(1, "b")
	v, ok := d.Get(1)
	if !ok || v != "b" {
		t.Errorf("got (%#v, %t), want (%#v, %t)", v, ok, "b", true)
	}
}

func TestDict_Set5(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Error("expected panic")
		}
	}()
	d := lib.Dict()
	d.Set(1, "a")
	d.Set("b", "b")
}

func TestDict_Set6(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Error("expected panic")
		}
	}()
	d := lib.Dict()
	d.Set(1, "a")
	d.Set(2, 2)
}

func TestDict_Keys(t *testing.T) {
	d := lib.Dict()
	d.Set("a", 1)
	d.Set("b", 2)
	d.Set("c", 3)
	ks := d.Keys()
	if len(ks) != 3 || ks[0] != "a" || ks[1] != "b" || ks[2] != "c" {
		t.Errorf("got %+v, want %+v", ks, []interface{}{"a", "b", "c"})
	}
	ks = d.Keys()
	if len(ks) != 3 || ks[0] != "a" || ks[1] != "b" || ks[2] != "c" {
		t.Errorf("got %+v, want %+v", ks, []interface{}{"a", "b", "c"})
	}
}
