package stapi

import (
	"github.com/pmukhin/klingon-translator/pkg/klingon/stapi/character"
	"github.com/pmukhin/klingon-translator/pkg/klingon/stapi/http"
)

const (
	baseUrl = "http://stapi.co"
)

// Client is a top level interface to access Stapi.co
// Currently it contains access to Characters and Species
type Client interface {
	Characters() character.CharactersClient
}

type defaultClient struct {
	characters character.CharactersClient
}

// Characters returns Characters client
func (d defaultClient) Characters() character.CharactersClient {
	return d.characters
}

// New is a constructor for default implementation of Client
func New(client http.HttpClient) Client {
	return &defaultClient{
		characters: character.New(baseUrl, client),
	}
}
