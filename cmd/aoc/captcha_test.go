package main

import "testing"

func TestCaptcha1122(t *testing.T) {
	r := Captcha("1122")
	if r != 3 {
		t.Errorf("got %d, want %d", r, 3)
	}
}

func TestCaptcha1111(t *testing.T) {
	r := Captcha("1111")
	if r != 4 {
		t.Errorf("got %d, want %d", r, 4)
	}
}

func TestCaptcha1234(t *testing.T) {
	r := Captcha("1234")
	if r != 0 {
		t.Errorf("got %d, want %d", r, 0)
	}
}

func TestCaptcha91212129(t *testing.T) {
	r := Captcha("91212129")
	if r != 9 {
		t.Errorf("got %d, want %d", r, 9)
	}
}
