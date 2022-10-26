package http

type HttpClientInterface interface {
	AddHeaders(headers []RequestHeader)
	AddParams(params []RequestParam)
	Get(reqUrl string, v interface{}) (*GetResponse, error)
}
