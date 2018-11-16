package day02_test

import (
	"testing"

	"github.com/SHyx0rmZ/advent-of-code/year2017/day02"
)

func TestChecksum(t *testing.T) {
	r, err := day02.ChecksumMinMax([]byte("5\t1\t9\t5\n7\t5\t3\n2\t4\t6\t8\n"))
	if r != 18 || err != nil {
		t.Errorf("got (%d, %+v), want (%d, %+v)", r, err, 18, nil)
	}
}

func TestChecksumDivision(t *testing.T) {
	r, err := day02.ChecksumDivision([]byte("5\t9\t2\t8\n9\t4\t7\t3\n3\t8\t6\t5"))
	if r != 9 || err != nil {
		t.Errorf("got (%d, %+v), want (%d, %+v)", r, err, 9, nil)
	}
}
