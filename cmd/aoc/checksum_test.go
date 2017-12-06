package main

import "testing"

func TestChecksum(t *testing.T) {
	r, err := Checksum([]byte("5\t1\t9\t5\n7\t5\t3\n2\t4\t6\t8\n"))
	if r != 18 || err != nil {
		t.Errorf("got (%d, %+v), want (%d, %+v)", r, err, 18, nil)
	}
}