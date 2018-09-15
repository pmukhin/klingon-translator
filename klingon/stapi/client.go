package stapi

type Response struct{}

type StapiClient interface {
	Search(name []rune) Response
}

func New() StapiClient {
	return nil
}
