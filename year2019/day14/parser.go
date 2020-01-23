package day14

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

type parser struct {
	chN   chan int
	chS   chan string
	chR   chan struct{}
	chErr chan error
}

type ingredient struct {
	Amount    int
	Component string
}

type recipe struct {
	Out ingredient
	In  []ingredient
}

func Parse(r io.Reader) ([]recipe, error) {
	p := parser{
		chN:   make(chan int),
		chS:   make(chan string),
		chR:   make(chan struct{}),
		chErr: make(chan error),
	}
	return p.parse(r)
}

func (p *parser) parse(r io.Reader) ([]recipe, error) {
	var rs []recipe
	go p.scan(r)
	for {
		var n int
		select {
		case n = <-p.chN:
		case err := <-p.chErr:
			if err != nil {
				return nil, err
			}
			return rs, nil
		}
		var i []ingredient
		for {
			i = append(i, ingredient{
				Amount:    n,
				Component: <-p.chS,
			})
			select {
			case err := <-p.chErr:
				return nil, err
			case n = <-p.chN:
				continue
			case <-p.chR:
			}
			break
		}
		select {
		case n = <-p.chN:
		case err := <-p.chErr:
			return nil, err
		}
		rs = append(rs, recipe{
			In: i,
			Out: ingredient{
				Amount:    n,
				Component: <-p.chS,
			},
		})
	}
}

func (p *parser) scan(r io.Reader) {
	s := bufio.NewScanner(r)
	for s.Scan() {
		ws := strings.Split(s.Text(), " ")
		for _, w := range ws[:len(ws)-1] {
			switch {
			case w == "=>":
				p.chR <- struct{}{}
			case w[0] >= '1' && w[0] <= '9':
				n, err := strconv.Atoi(w)
				if err != nil {
					p.chErr <- err
				}
				p.chN <- n
			default:
				p.chS <- strings.TrimSuffix(w, ",")
			}
		}
		p.chS <- ws[len(ws)-1]
	}
	p.chErr <- s.Err()
}
