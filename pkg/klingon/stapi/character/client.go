package character

import (
	"fmt"
	"github.com/pmukhin/klingon-translator/pkg/klingon/stapi/http"
	"net/url"
	"strings"
)

const (
	httpFormContentType = "application/x-www-form-urlencoded"
)

var (
	defaultCharacterSpecies = Species{
		UID:  "",
		Name: "Human",
	}
)

// UID ...
type UID string

type searchResponse struct {
	Characters []Short
}

type fullResponse struct {
	Character *Full `json:"character"`
}

// CharactersClient is client sourcing Characters
// from Stapi.co
type CharactersClient interface {
	Search(name string) ([]Short, error)
	Get(uid string) (*Full, error)
}

type defaultCharactersClient struct {
	client  http.HttpClient
	baseUrl string
}

// Search sends a post request to characters resource
func (d defaultCharactersClient) Search(name string) ([]Short, error) {
	defRet := make([]Short, 0)
	nameSlice := []string{name} // just for not to create this slice twice
	values := url.Values{
		"name":  nameSlice,
		"title": nameSlice,
	}

	response, err := d.client.Post(
		d.baseUrl+"/api/v1/rest/character/search",
		httpFormContentType,
		strings.NewReader(values.Encode()),
	)

	if err != nil {
		return defRet, serverUnavailableErr(err)
	}

	var sr searchResponse
	if err := http.ReadAsJson(response.Body, &sr); err != nil {
		return defRet, err
	}

	return sr.Characters, nil
}

// Get sends a request to characters resource
func (d defaultCharactersClient) Get(uid string) (*Full, error) {
	response, err := d.client.Get(d.baseUrl + "/api/v1/rest/character?uid=" + uid)
	if err != nil {
		return nil, serverUnavailableErr(err)
	}

	var cr fullResponse
	if err := http.ReadAsJson(response.Body, &cr); err != nil {
		return nil, err
	}

	return d.normalize(cr.Character), nil
}

func serverUnavailableErr(err error) error {
	return fmt.Errorf("server is not available: %s", err.Error())
}

func (d defaultCharactersClient) normalize(character *Full) *Full {
	if len(character.CharacterSpecies) > 0 {
		return character
	}
	character.CharacterSpecies = append(character.CharacterSpecies, defaultCharacterSpecies)

	return character
}

// New is defaultCharactersClient constructor
func New(baseUrl string, client http.HttpClient) CharactersClient {
	return &defaultCharactersClient{baseUrl: baseUrl, client: client}
}
