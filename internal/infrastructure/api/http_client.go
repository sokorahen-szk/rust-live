package api

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const httpClientTimeout = 5

type HttpClient struct {
	req    *http.Request
	params url.Values
}

type RequestHeader struct {
	key   string
	value string
}

type RequestParam struct {
	key   string
	value string
}

type GetJsonResponse struct {
	v interface{}
}

func NewHttpClient(method string, url string, body io.Reader) *HttpClient {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		panic(err)
	}

	httpClient := &HttpClient{
		req: req,
	}

	httpClient.params = req.URL.Query()

	return httpClient
}

func (r *HttpClient) AddHeaders(headers []RequestHeader) {
	for _, header := range headers {
		r.req.Header.Add(header.key, header.value)
	}
}

func (r *HttpClient) AddParams(params []RequestParam) {
	for _, param := range params {
		r.params.Add(param.key, param.value)
	}
}

func (r *HttpClient) GetJson(v interface{}) (*GetJsonResponse, error) {
	r.req.URL.RawQuery = r.params.Encode()

	client := r.createClient()
	httpResponse, err := client.Do(r.req)
	if err != nil {
		return nil, err
	}

	defer httpResponse.Body.Close()

	body, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(body), v)
	if err != nil {
		return nil, nil
	}

	response := &GetJsonResponse{v: v}
	return response, nil
}

func (r *HttpClient) createClient() *http.Client {
	return &http.Client{
		Timeout: time.Duration(httpClientTimeout * time.Second),
	}
}
