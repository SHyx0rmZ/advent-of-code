package day05

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strings"

	"github.com/SHyx0rmZ/advent-of-code"
)

type problem struct{}

func Problem() aoc.ReaderAwareProblem {
	return &problem{}
}

var vowels = regexp.MustCompile(`([aeiou].*){3,}`)
var repeating = regexp.MustCompile(`aa|bb|cc|dd|ee|ff|gg|hh|ii|jj|kk|ll|mm|nn|oo|pp|qq|rr|ss|tt|uu|vv|ww|xx|yy|zz`)
var forbidden = regexp.MustCompile(`ab|cd|pq|xy`)

func (p problem) PartOneWithReader(r io.Reader) (string, error) {
	var i int
	s := bufio.NewScanner(r)
	for s.Scan() {
		if forbidden.MatchString(s.Text()) {
			continue
		}
		if !vowels.MatchString(s.Text()) {
			continue
		}
		if !repeating.MatchString(s.Text()) {
			continue
		}
		i++
	}
	return fmt.Sprintf("%d", i), s.Err()
}

var repeatSpaced, repeatNonOverlapping *regexp.Regexp

func init() {
	chars := "abcdefghijklmnopqrstuvwxyz"
	s := strings.Builder{}
	for _, r := range chars {
		s.WriteString(fmt.Sprintf("%c.%c", r, r))
		if r != 'z' {
			s.WriteByte('|')
		}
	}
	repeatSpaced = regexp.MustCompile(s.String())
	s = strings.Builder{}
	for _, r1 := range chars {
		for _, r2 := range chars {
			s.WriteString(fmt.Sprintf("%c%c.*%c%c", r1, r2, r1, r2))
			if r2 != 'z' {
				s.WriteByte('|')
			}
		}
		if r1 != 'z' {
			s.WriteByte('|')
		}
	}
	repeatNonOverlapping = regexp.MustCompile(s.String())
}

func (p problem) PartTwoWithReader(r io.Reader) (string, error) {
	var i int
	s := bufio.NewScanner(r)
	for s.Scan() {
		if !repeatSpaced.MatchString(s.Text()) {
			continue
		}
		if !repeatNonOverlapping.MatchString(s.Text()) {
			continue
		}
		i++
	}
	return fmt.Sprintf("%d", i), s.Err()
}
