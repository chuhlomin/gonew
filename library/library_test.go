package main // should be "package library", but gonew doesn't support it

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type httpClientMock struct {
	mock.Mock
}

func (h *httpClientMock) Do(req *http.Request) (*http.Response, error) {
	args := h.Called(req)
	return args.Get(0).(*http.Response), args.Error(1)
}

type closingBuffer struct {
	*bytes.Buffer
}

func (cb *closingBuffer) Close() error {
	return nil
}

func TestSomething(t *testing.T) {
	// record HTTP request
	httpClient := new(httpClientMock)
	httpClient.
		On("Do", mock.MatchedBy(func(req *http.Request) bool {
			return req.Method == http.MethodPost &&
				req.URL.String() == "https://server.io/api" &&
				req.Header.Get("User-Agent") == "library/1.0 (go1.16.5/darwin)" &&
				req.Header.Get("Content-Type") == "application/json" &&
				req.Header.Get("Accept") == "application/json" &&
				string(mustRead(req.Body)) == `{"field":"req"}`
		})).
		Return(
			&http.Response{
				StatusCode: http.StatusOK,
				Body: &closingBuffer{
					bytes.NewBufferString(`{"field":"resp"}`),
				},
			},
			nil,
		)

	client, err := NewMyStruct(
		httpClient,
		"library/1.0 (go1.16.5/darwin)",
		"https://server.io/api",
	)
	assert.NoError(t, err)

	resp, err := client.DoSomething(MyRequest{
		Field: "req",
	})

	assert.NoError(t, err)
	assert.Equal(
		t,
		&MyResponse{
			Field: "resp",
		},
		resp,
	)
}

func mustRead(r io.ReadCloser) []byte {
	defer r.Close()
	b, err := io.ReadAll(r)
	if err != nil {
		panic(err)
	}
	return b
}
