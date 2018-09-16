package klingon

import (
	"errors"
	"fmt"
	"github.com/pmukhin/klingon-translator/klingon/parser"
	"github.com/pmukhin/klingon-translator/klingon/stapi"
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

	stapiClient := stapi.New()

	foundCharacters := stapiClient.Characters().Search(name)
	if len(foundCharacters) < 1 {
		return errors.New(fmt.Sprintf("character %s is not found via Stapi.co", name))
	}

	character := stapiClient.Characters().Get(foundCharacters[0].UID)
	if character == nil {
		return errors.New(
			fmt.Sprintf(
				"unexpected error: %s is not found via Stapi.co by UID %s",
				name,
				foundCharacters[0].UID,
			),
		)
	}

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
