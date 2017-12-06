package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

//go:generate go-bindata ../../data/...

func main() {
	c, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	c = bytes.TrimSpace(c)

	fmt.Printf("%d\n", Captcha(string(c)))
}