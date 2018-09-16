package character

import (
	"github.com/pmukhin/klingon-translator/klingon/util"
	"net/http"
	"net/url"
)

type UID string

type SearchResponse struct {
	Characters []Short
}

// CharacterResponse
type CharacterResponse struct {
	Character *Full `json:"character"`
}

type CharactersClient interface {
	Search(name string) ([]Short, error)
	Get(uid string) (*Full, error)
}

type defaultCharactersClient struct {
	baseUrl string
}

func (d defaultCharactersClient) Search(name string) ([]Short, error) {
	defRet := make([]Short, 0)
	response, err := http.PostForm(d.baseUrl+"/api/v1/rest/character/search", url.Values{
		"name": []string{name},
	})

	if err != nil {
		return defRet, err
	}

	var sr SearchResponse
	if err := util.ReadResponse(response, &sr); err != nil {
		return defRet, err
	}

	return sr.Characters, nil
}

func (d defaultCharactersClient) Get(uid string) (*Full, error) {
	response, err := http.Get(d.baseUrl + "/api/v1/rest/character?uid=" + uid)
	if err != nil {
		return nil, err
	}

	var cr CharacterResponse
	if err := util.ReadResponse(response, &cr); err != nil {
		return nil, err
	}

	return cr.Character, nil
}

func New(baseUrl string) CharactersClient {
	return &defaultCharactersClient{baseUrl: baseUrl}
}
