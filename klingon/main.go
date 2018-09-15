package klingon

import (
	"fmt"
	"github.com/pmukhin/klingon-translator/klingon/parser"
)

const (
	searchUri = "http://stapi.co/api/v1/rest/character/search"
)

type response struct {
	parsedName    []rune
	characterName string
}

func Main(name string) error {
	lexer := parser.New(name)
	translatedName, err := lexer.Parse()

	if err != nil {
		return err
	}

	//client := stapi.New()
	//_ = client.Search(translatedName)

	res := response{
		parsedName:    translatedName,
		characterName: "",
	}

	render(res)

	return nil
}

func render(res response) {
	for _, ch := range res.parsedName {
		fmt.Printf("0x%X ", ch)
	}
	fmt.Println()
}
