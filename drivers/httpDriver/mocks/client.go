package mocks

import http "net/http"

type HttpClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

func (m *HttpClient) Do(req *http.Request) (*http.Response, error) {
	if m.DoFunc != nil {
		return m.DoFunc(req)
	}
	return &http.Response{}, nil
}
