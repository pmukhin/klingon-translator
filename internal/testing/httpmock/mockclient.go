package httpmock

import (
	"bytes"
	"container/list"
	"io"
	"net/http"
)

type mockReadCloser struct {
	*bytes.Buffer
}

func (mockReadCloser) Close() error {
	return nil
}

func newReadCloser(data string) io.ReadCloser {
	return &mockReadCloser{Buffer: bytes.NewBufferString(data)}
}

type MockClientResponse struct {
	Status int
	Body   string
}

type MockClient struct {
	responseQueue *list.List
}

func (m *MockClient) Get(uri string) (*http.Response, error) {
	if m.responseQueue.Len() == 0 {
		panic("response queue is empty")
	}
	return m.responseQueue.Front().Value.(*http.Response), nil
}

func (m *MockClient) Post(uri, contentType string, body io.Reader) (*http.Response, error) {
	if m.responseQueue.Len() == 0 {
		panic("response queue is empty")
	}
	return m.responseQueue.Front().Value.(*http.Response), nil
}

func New(q []MockClientResponse) *MockClient {
	queue := list.New()

	for _, mockRes := range q {
		queue.PushBack(&http.Response{
			Status:     "",
			StatusCode: mockRes.Status,
			Body:       newReadCloser(mockRes.Body),
		})
	}

	return &MockClient{responseQueue: queue}
}
