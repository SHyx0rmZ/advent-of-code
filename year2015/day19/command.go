package day19

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"strconv"
	"strings"

	"github.com/SHyx0rmZ/advent-of-code"
)

type problem struct {
}

func Problem() aoc.ReaderAwareProblem {
	return &problem{}
}

func (p problem) PartOneWithReader(r io.Reader) (string, error) {
	rules := make(map[string][]string)
	var molecule string
	s := bufio.NewScanner(r)
	for s.Scan() {
		if len(s.Text()) == 0 {
			break
		}
		ps := strings.Split(s.Text(), " => ")
		rules[ps[0]] = append(rules[ps[0]], ps[1])
	}
	if err := s.Err(); err != nil {
		return "", err
	}
	if s.Scan() {
		molecule = s.Text()
	}
	if err := s.Err(); err != nil {
		return "", err
	}
	variants := make(map[string]struct{})
	for i := range molecule {
		for j := range molecule[i:] {
			k := molecule[i:i+j+1]
			rs, ok := rules[k]
			if !ok {
				continue
			}
			for _, r := range rs {
				s := molecule[:i] + strings.Replace(molecule[i:], k, r, 1)
				variants[s] = struct{}{}
			}
		}
	}
	return strconv.Itoa(len(variants)), nil
}

func (p problem) PartTwoWithReader(r io.Reader) (string, error) {
	rules := make(map[string][]string)
	var molecule string
	s := bufio.NewScanner(r)
	for s.Scan() {
		if len(s.Text()) == 0 {
			break
		}
		ps := strings.Split(s.Text(), " => ")
		rules[ps[1]] = append(rules[ps[1]], ps[0])
	}
	if err := s.Err(); err != nil {
		return "", err
	}
	if s.Scan() {
		molecule = s.Text()
	}
	if err := s.Err(); err != nil {
		return "", err
	}
	var max int
	for k := range rules {
		if len(k) > max {
			max = len(k)
		}
	}
	seen := make(map[string]int)
	var step func(molecule string, steps int) int
	step = func(molecule string, steps int) int {
		if molecule == "e" {
			fmt.Println(steps)
			return steps
		}
		if n, ok := seen[molecule]; ok {
			return n
		}
		//fmt.Println(molecule)
		variants := make(map[string]struct{})
		for i := range molecule {
			for k, rs := range rules {
				for _, r := range rs {
					s := molecule[:len(molecule)-i-1] + strings.Replace(molecule[len(molecule)-i-1:], k, r, 1)
					variants[s] = struct{}{}
				}
			}
		}
		delete(variants, molecule)
		if len(variants) == 0 {
			seen[molecule] = steps + len(molecule)
			return steps + len(molecule)
		}
		ks := make([]string, 0, len(variants))
		for k := range variants {
			ks = append(ks, k)
		}
		rand.Shuffle(len(ks), func(i, j int) {
			ks[i], ks[j] = ks[j], ks[i]
		})
		//sort.Slice(ks, func(i, j int) bool {
		//	if len(ks[i]) == len(ks[j]) {
		//		return ks[i] < ks[j]
		//	}
		//	return len(ks[i]) < len(ks[j])
		//})
		min := step(ks[0], steps+1)
		for _, k := range ks[1:] {
			s := step(k, steps+1)
			if s < min {
				min = s
			}
		}
		seen[molecule] = min
		return min
	}
	n := step(molecule, 0)
	return strconv.Itoa(n), nil
}
