package klingon

import (
	"errors"
	"fmt"
	"github.com/pmukhin/klingon-translator/pkg/klingon/parser"
	"github.com/pmukhin/klingon-translator/pkg/klingon/stapi"
	"github.com/pmukhin/klingon-translator/pkg/klingon/stapi/character"
	"net/http"
	"strings"
)

type response struct {
	parsedName []rune
	species    string
}

// Main is entry point of klingon pkg
func Main(name string) error {
	lexer := parser.New(name)
	translatedName, err := lexer.Parse()

	if err != nil {
		return err
	}

	stapiClient := stapi.New(http.DefaultClient)

	foundCharacters, err := stapiClient.Characters().Search(name)
	if err != nil {
		return err
	}
	if len(foundCharacters) < 1 {
		return errors.New(fmt.Sprintf("character %s is not found via Stapi.co", name))
	}

	characterObj, err := stapiClient.Characters().Get(foundCharacters[0].UID)
	if characterObj == nil {
		return errors.New(
			fmt.Sprintf(
				"unexpected error: %s is not found via Stapi.co by UID %s",
				name,
				foundCharacters[0].UID,
			),
		)
	}

	res := response{
		parsedName: translatedName,
		species:    extractSpeciesNames(characterObj.CharacterSpecies),
	}

	render(res)

	return nil
}

func extractSpeciesNames(speciesObjs []character.Species) string {
	species := make([]string, 0)
	for _, s := range speciesObjs {
		species = append(species, s.Name)
	}
	return strings.Join(species, ", ")
}

func render(res response) {
	for _, ch := range res.parsedName {
		fmt.Printf("0x%X ", ch)
	}
	fmt.Println("\n" + res.species)
}
