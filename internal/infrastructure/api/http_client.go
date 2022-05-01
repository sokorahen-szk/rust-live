package api

import (
	"io"
	"net/http"
)

type HttpClient struct {
	req *http.Request
}

type Header struct {
	key   string
	value string
}

func NewHttpClient(method string, url string, body io.Reader) *HttpClient {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		panic(err)
	}
	return &HttpClient{
		req: req,
	}
}

func (r *HttpClient) AddHeader(headers []Header) {
	for _, header := range headers {
		r.req.Header.Add(header.key, header.value)
	}
}
