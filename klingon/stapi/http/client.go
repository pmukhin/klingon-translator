package http

import (
	"io"
	"net/http"
)

// HttpClient is internal http client representation
// mainly required for testabilty
type HttpClient interface {
	Get(uri string) (*http.Response, error)
	Post(uri, contentType string, body io.Reader) (*http.Response, error)
}

