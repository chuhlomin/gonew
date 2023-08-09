// Package library ...
package library

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type httpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// MyStruct ...
type MyStruct struct {
	userAgent  string
	url        string
	httpClient httpClient
}

// NewMyStruct creates a new MyStruct ...
func NewMyStruct(
	httpClient httpClient,
	userAgent string,
	url string,
) (*MyStruct, error) {
	return &MyStruct{
		httpClient: httpClient,
		userAgent:  userAgent,
		url:        url,
	}, nil
}

// MyRequest ...
type MyRequest struct {
	Field string `json:"field"`
}

// MyResponse ...
type MyResponse struct {
	Field string `json:"field"`
}

// DoSomething ...
func (s *MyStruct) DoSomething(r MyRequest) (*MyResponse, error) {
	b, err := json.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("failed to encode request: %w", err)
	}

	req, err := http.NewRequest("POST", s.url, bytes.NewBuffer(b))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("User-Agent", s.userAgent)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}

	defer resp.Body.Close()

	var response MyResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &response, nil
}
