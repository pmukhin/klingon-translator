package parser

import (
	"errors"
	"fmt"
)

var (
	// end of line error
	EOF = errors.New("EOF")
)

// Parser tokenizes input into chars and translates each character
// into a corresponding Klingon character
type Parser interface {
	Parse() ([]rune, error)
}

// defaultParser is a default Parser implementation
type defaultParser struct {
	src    []byte
	offset int
	len    int
}

// Parse combines tokens into a slice of runes
func (p defaultParser) Parse() ([]rune, error) {
	output := make([]rune, 0, 256)
	for {
		ch, err := p.scan()
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

// scan is a core method here
// same pattern is used as golang lexer uses
// https://github.com/golang/go/blob/master/src/go/scanner/scanner.go
func (p *defaultParser) scan() (rune, error) {
	for {
		p.next()
		if p.offset >= p.len {
			// oops, all input is eaten
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
				p.next() // eat 'h' or 'H'
				return 0xF8D2, nil
			} else {
				// standalone c(C) is not a valid char in Klingon hence error
				return 0, p.makeErr()
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
				// standalone g(G) is not a valid char in Klingon hence error
				return 0, p.makeErr()
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
				// both n(N) and ng(nG, NG, Ng) are valid, so no error here
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
			// as t(T) valid standalone char let's first check if there's no l(L) in front
			nextCh := p.peek()
			if nextCh == 'l' || nextCh == 'L' {
				p.next()               // eat l(L)
				nextNextCh := p.peek() // as standalone l(L) is not valid, let's dig deeper
				if nextNextCh == 'h' || nextNextCh == 'H' {
					p.next() // bingo, it's tlh
					return 0xF8E4, nil
				} else {
					// tl(TL, tL, Tl) are not valid, so error
					return 0, p.makeErr()
				}
			} else {
				// if no then just t(T)
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
			return -1, p.makeErr()
		}
	}
}

// makeErr constructs an error from the context
func (p defaultParser) makeErr() error {
	return errors.New(
		fmt.Sprintf(
			"unexpected character %s in pos %d",
			string(p.peek()),
			p.offset+1,
		),
	)
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

// New is a constructor for a defaultParser
func New(input string) Parser {
	src := []byte(input)
	return &defaultParser{
		src:    src,
		offset: -1,
		len:    len(input),
	}
}
