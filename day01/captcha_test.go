package day01_test

import (
	"testing"

	"github.com/SHyx0rmZ/advent-of-code/day01"
)

func TestCaptcha1122(t *testing.T) {
	r := day01.CaptchaNext("1122")
	if r != 3 {
		t.Errorf("got %d, want %d", r, 3)
	}
}

func TestCaptcha1111(t *testing.T) {
	r := day01.CaptchaNext("1111")
	if r != 4 {
		t.Errorf("got %d, want %d", r, 4)
	}
}

func TestCaptcha1234(t *testing.T) {
	r := day01.CaptchaNext("1234")
	if r != 0 {
		t.Errorf("got %d, want %d", r, 0)
	}
}

func TestCaptcha91212129(t *testing.T) {
	r := day01.CaptchaNext("91212129")
	if r != 9 {
		t.Errorf("got %d, want %d", r, 9)
	}
}

func TestCaptchaHalf1212(t *testing.T) {
	r := day01.CaptchaHalf("1212")
	if r != 6 {
		t.Errorf("got %d, want %d", r, 6)
	}
}

func TestCaptchaHalf1221(t *testing.T) {
	r := day01.CaptchaHalf("1221")
	if r != 0 {
		t.Errorf("got %d, want %d", r, 0)
	}
}

func TestCaptchaHalf123425(t *testing.T) {
	r := day01.CaptchaHalf("123425")
	if r != 4 {
		t.Errorf("got %d, want %d", r, 4)
	}
}

func TestCaptchaHalf123123(t *testing.T) {
	r := day01.CaptchaHalf("123123")
	if r != 12 {
		t.Errorf("got %d, want %d", r, 12)
	}
}

func TestCaptchaHalf12131415(t *testing.T) {
	r := day01.CaptchaHalf("12131415")
	if r != 4 {
		t.Errorf("got %d, want %d", r, 4)
	}
}
