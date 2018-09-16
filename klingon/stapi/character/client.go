package character

type UID string

type SearchResponse struct {
	Characters []Short
}

// CharacterResponse
type characterResponse struct {
	Character Full `json:"character"`
}

type CharactersClient interface {
	Search(name string) []Short
	Get(uid string) *Full
}

type defaultCharactersClient struct{}

func (defaultCharactersClient) Search(name string) []Short {
	panic("implement me")
}

func (defaultCharactersClient) Get(uid string) *Full {
	panic("implement me")
}

func New() CharactersClient {
	return &defaultCharactersClient{}
}
