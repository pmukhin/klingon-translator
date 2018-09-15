package parser

import (
	"errors"
	"fmt"
)

var (
	EOF = errors.New("EOF")
)

type Parser interface {
	Parse() ([]rune, error)
}

type defaultParser struct {
	src    []byte
	offset int
	len    int
}

func makeErr(r byte, pos int) error {
	return errors.New(
		fmt.Sprintf(
			"unexpected character %s in pos %d",
			string(r),
			pos,
		),
	)
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
		if p.offset >= p.len {
			return -1, EOF
		}
		ch := p.src[p.offset]
		switch ch {
		case 'a', 'A':
			return 0xF8D0, nil
		case 'b', 'B':
			return 0xF8D1, nil
		case 'c', 'C':
			nextCh := p.peek()
			if nextCh == 'h' || nextCh == 'H' {
				p.next()
				return 0xF8D2, nil
			} else {
				return 0, makeErr(nextCh, p.offset+1)
			}
		case 'd', 'D':
			return 0xF8D3, nil
		case 'e', 'E':
			return 0xF8D4, nil
		case 'g', 'G':
			nextCh := p.peek()
			if nextCh == 'h' || nextCh == 'H' {
				p.next()
				return 0xF8D5, nil
			} else {
				return 0, makeErr(nextCh, p.offset+1)
			}
		case 'h', 'H':
			return 0xF8D6, nil
		case 'i', 'I':
			return 0xF8D7, nil
		case 'j', 'J':
			return 0xF8D8, nil
		case 'l', 'L':
			return 0xF8D9, nil
		case 'm', 'M':
			return 0xF8DA, nil
		case 'n', 'N':
			nextCh := p.peek()
			if nextCh == 'g' || nextCh == 'G' {
				p.next()
				return 0xF8DC, nil
			} else {
				return 0xF8DB, nil
			}
		case 'o', 'O':
			return 0xF8DD, nil
		case 'p', 'P':
			return 0xF8DE, nil
		case 'q':
			return 0xF8DF, nil
		case 'Q':
			return 0xF8E0, nil
		case 'r', 'R':
			return 0xF8E1, nil
		case 's', 'S':
			return 0xF8E2, nil
		case 't', 'T':
			nextCh := p.peek()
			if nextCh == 'l' || nextCh == 'L' {
				p.next()
				nextNextCh := p.peek()
				if nextNextCh == 'h' || nextNextCh == 'H' {
					p.next()
					return 0xF8E4, nil
				} else {
					return 0, makeErr(nextCh, p.offset)
				}
			} else {
				return 0xF8E3, nil
			}
		case 'u', 'U':
			return 0xF8E5, nil
		case 'v', 'V':
			return 0xF8E6, nil
		case 'w', 'W':
			return 0xF8E7, nil
		case 'y', 'Y':
			return 0xF8E8, nil
		case '\'':
			return 0xF8E9, nil
		case ' ':
			return ' ', nil
		default:
			return -1, makeErr(p.peek(), p.offset+1)
		}
	}
}

func (p *defaultParser) next() {
	p.offset++
}

func (p defaultParser) peek() byte {
	if p.offset+1 >= p.len {
		return 0
	}
	return p.src[p.offset+1]
}

func New(input string) Parser {
	src := []byte(input)
	return &defaultParser{
		src:    src,
		offset: -1,
		len:    len(input),
	}
}
