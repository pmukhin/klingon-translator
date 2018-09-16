package klingon

import (
	"errors"
	"fmt"
	"github.com/pmukhin/klingon-translator/klingon/parser"
	"github.com/pmukhin/klingon-translator/klingon/stapi"
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

	foundCharacters, err := stapiClient.Characters().Search(name)
	if err != nil {
		return err
	}
	if len(foundCharacters) < 1 {
		return errors.New(fmt.Sprintf("character %s is not found via Stapi.co", name))
	}

	character, err := stapiClient.Characters().Get(foundCharacters[0].UID)
	if character == nil {
		return errors.New(
			fmt.Sprintf(
				"unexpected error: %s is not found via Stapi.co by UID %s",
				name,
				foundCharacters[0].UID,
			),
		)
	}

	species := "Human"
	if len(character.CharacterSpecies) > 0 {
		species = character.CharacterSpecies[0].Name
	}

	res := response{
		parsedName:    translatedName,
		characterName: species,
	}

	render(res)

	return nil
}

func render(res response) {
	for _, ch := range res.parsedName {
		fmt.Printf("0x%X ", ch)
	}
	fmt.Println()
	fmt.Println(res.characterName)
}
