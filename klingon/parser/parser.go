package parser

type Parser interface {
	Parse() ([]rune, error)
}

type defaultParser struct {
	src    []rune
	offset int
}

func (p defaultParser) Parse() ([]rune, error) {
	return []rune{}, nil
}

func New(input []rune) Parser {
	return &defaultParser{src: input, offset: -1}
}
