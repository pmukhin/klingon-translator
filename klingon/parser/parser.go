package parser

import (
	"errors"
	"strings"
)

var (
	EOF = errors.New("EOF")
)

type Parser interface {
	Parse() ([]rune, error)
}

type defaultParser struct {
	src    []rune
	offset int
	len    int
}

func (p defaultParser) Parse() ([]rune, error) {
	output := make([]rune, 0, 256)
	for {
		ch, err := p.getChar()
		if err != nil {
			if err == EOF {
				return output, nil
			}
			return output, err
		}
		output = append(output, ch)
	}
	return output, nil
}

func (p *defaultParser) getChar() (rune, error) {
	for {
		p.next()
		if p.offset == p.len {
			return -1, EOF
		}
		ch := p.src[p.offset]
		switch ch {
		case 'u':
			return 0xF8E5, nil
		case 'h':
			return 0xF8D6, nil
		case 'r':
			return 0xF8E1, nil
		case 'a':
			return 0xF8D0, nil
		default:
			return -1, nil
		}
	}
}

func (p *defaultParser) next() {
	p.offset++
}

func New(input string) Parser {
	return &defaultParser{
		src:    []rune(strings.ToLower(input)),
		offset: -1,
		len:    len(input),
	}
}
