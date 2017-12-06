package day04

import (
	"strings"
	"os"
	"io/ioutil"
	"bytes"
	"fmt"
	"sort"
	"strconv"
)

func Command() error {
	if len(os.Args) < 4 {
		panic("not enough arguments")
	}

	var f *os.File
	var err error

	if os.Args[3] == "-" {
		f = os.Stdin
	} else {
		f, err = os.OpenFile(os.Args[3], os.O_RDONLY, 0755)
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

	var sum int

	switch os.Args[2] {
	case "duplicate":
		for _, passphrase := range bytes.Split(c, []byte("\n")) {
			if Valid(string(passphrase)) {
				sum++
			}
		}
	case "anagram":
		for _, passphrase := range bytes.Split(c, []byte("\n")) {
			if ValidEx(string(passphrase)) {
				sum++
			}
		}
	default:
		panic("unknown sub-command: " + os.Args[2])
	}

	if err != nil {
		return err
	}

	_, err = fmt.Printf("%d\n", sum)

	return err
}

func Valid(p string) bool {
	set := make(map[string]struct{})
	for _, word := range strings.Split(p, " ") {
		_, ok := set[word]
		if ok {
			return false
		}
		set[word] = struct{}{}
	}
	return true
}

func ValidEx(p string) bool {
	set := make(map[string]struct{})
	for _, word := range strings.Split(p, " ") {
		set2 := make(map[rune]int)
		for _, r := range word {
			set2[r]++
		}
		var keys []int
		for r := range set2 {
			keys = append(keys, int(r))
		}
		sort.Ints(keys)
		key := ""
		for _, k := range keys {
			key += string(rune(k)) + strconv.Itoa(set2[rune(k)])
		}
		_, ok := set[key]
		if ok {
			return false
		}
		set[key] = struct{}{}
	}
	return true
}
