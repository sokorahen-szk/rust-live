package http

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
	Key   string
	Value string
}

type RequestParam struct {
	Key   string
	Value string
}

type GetResponse struct {
	Data interface{}
}

var globalParams map[string]string

func NewHttpClient(method string, body io.Reader) HttpClientInterface {
	req, err := http.NewRequest(method, "", body)
	if err != nil {
		panic(err)
	}

	httpClient := &HttpClient{
		req: &http.Request{},
	}

	httpClient.params = req.URL.Query()

	globalParams = map[string]string{}

	return httpClient
}

func (r *HttpClient) AddHeaders(headers []RequestHeader) {
	for _, header := range headers {
		r.req.Header.Add(header.Key, header.Value)
	}
}

func (r *HttpClient) AddParams(params []RequestParam) {
	for _, param := range params {
		globalParams[param.Key] = param.Value
		r.params.Add(param.Key, param.Value)
	}
}

func (r *HttpClient) Get(reqUrl string, v interface{}) (*GetResponse, error) {
	changedUrl, err := url.Parse(reqUrl)
	if err != nil {
		return nil, err
	}
	r.req.URL = changedUrl
	r.req.URL.RawQuery = r.params.Encode()

	r.deleteParamAll()

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
		return nil, err
	}

	//if httpResponse.StatusCode != http.StatusOK {
	//	return nil, fmt.Errorf("%d", httpResponse.StatusCode)
	//}

	response := &GetResponse{Data: v}
	return response, nil
}

func (r *HttpClient) deleteParam(key string) {
	r.params.Del(key)
	delete(globalParams, key)
}

func (r *HttpClient) deleteParamAll() {
	for key, _ := range globalParams {
		r.deleteParam(key)
	}
}

func (r *HttpClient) createClient() *http.Client {
	return &http.Client{
		Timeout: time.Duration(httpClientTimeout * time.Second),
	}
}
