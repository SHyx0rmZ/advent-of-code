package day06

import (
	"bufio"
	"image"
	"io"
	"strconv"
	"strings"
	"time"
)

type Parser struct {
	*bufio.Scanner

	funcs  chan func()
	ops    chan Operation
	points chan image.Point
}

func NewParser(r io.Reader) *Parser {
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanWords)
	return &Parser{
		Scanner: s,
		funcs:   make(chan func(), 0),
		ops:     make(chan Operation, 0),
		points:  make(chan image.Point, 0),
	}
}

func (p *Parser) Operations() <-chan Operation { return p.ops }

func (p *Parser) on()     { p.ops <- TurnOn{Area{<-p.points, <-p.points}} }
func (p *Parser) off()    { p.ops <- TurnOff{Area{<-p.points, <-p.points}} }
func (p *Parser) toggle() { p.ops <- Toggle{Area{<-p.points, <-p.points}} }

func (p *Parser) Parse() error {
	go func() {
		for f := range p.funcs {
			f()
		}
	}()
	for p.Scan() {
		switch p.Text() {
		case "on":
			p.funcs <- p.on
		case "off":
			p.funcs <- p.off
		case "toggle":
			p.funcs <- p.toggle
		case "turn":
		case "through":
		default:
			ps := strings.Split(p.Text(), ",")
			x, err := strconv.Atoi(ps[0])
			if err != nil {
				return err
			}
			y, err := strconv.Atoi(ps[1])
			if err != nil {
				return err
			}
			p.points <- image.Point{x, y}
		}
	}
	time.Sleep(1 * time.Second)
	close(p.funcs)
	close(p.points)
	close(p.ops)
	return p.Err()
}
