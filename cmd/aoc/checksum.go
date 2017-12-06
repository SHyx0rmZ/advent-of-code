package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
)

func Checksum(data []byte) (int, error) {
	sum := 0
	for len(data) > 0 {
		low := math.MaxInt64
		high := math.MinInt64
		for {
			i := bytes.IndexAny(data, "\t \n")
			if i < 1 {
				break
			}
			n, err := strconv.Atoi(string(data[0:i]))
			if err != nil {
				return 0, err
			}
			if n > high {
				high = n
			}
			if n < low {
				low = n
			}
			for data[i] == '\t' || data[i] == ' ' {
				i++
			}
			data = data[i:]
		}
		sum += high - low
		i := bytes.IndexAny(data, "\n")
		if i != 0 {
			break
		}
		data = data[1:]
	}
	return sum, nil
}

func checksumCommand() error {
	if len(os.Args) < 3 {
		panic("not enough arguments")
	}

	var f *os.File
	var err error

	if os.Args[2] == "-" {
		f = os.Stdin
	} else {
		f, err = os.OpenFile(os.Args[2], os.O_RDONLY, 0755)
		if err != nil {
			return err
		}
		defer f.Close()
	}

	c, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}
	c = bytes.TrimSpace(c)

	sum, err := Checksum(c)
	if err != nil {
		return err
	}

	_, err = fmt.Printf("%d\n", sum)

	return err
}