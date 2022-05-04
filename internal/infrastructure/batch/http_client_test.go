package batch

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestHttpBinOrgGet struct {
	Args    *TestHttpBinOrgGetParam
	Url     string `json:"url"`
	Headers *TestHttpBinOrgGetHeaders
}

type TestHttpBinOrgGetParam struct {
	Test string `json:"test"`
}

type TestHttpBinOrgGetHeaders struct {
	Hoge string `json:"hoge"`
	Fuga string `json:"fuga"`
}

func Test_NewHttpClient_Get(t *testing.T) {
	a := assert.New(t)

	url := "https://httpbin.org/get"

	t.Run("Getパラメータなし", func(t *testing.T) {
		client := NewHttpClient(http.MethodGet, nil)
		httpBinOrgGet := &TestHttpBinOrgGet{}

		res, err := client.Get(url, httpBinOrgGet)
		a.NoError(err)
		a.NotNil(res)

		actual := res.Data.(*TestHttpBinOrgGet)
		a.Equal(url, actual.Url)
		a.Equal("", actual.Args.Test)
		a.Equal("", actual.Headers.Hoge)
		a.Equal("", actual.Headers.Fuga)
	})
	t.Run("Getパラメータあり", func(t *testing.T) {
		client := NewHttpClient(http.MethodGet, nil)
		httpBinOrgGet := &TestHttpBinOrgGet{}

		client.AddParams([]RequestParam{{"test", "abcd"}})
		res, err := client.Get(url, httpBinOrgGet)
		a.NoError(err)
		a.NotNil(res)

		actual := res.Data.(*TestHttpBinOrgGet)
		a.Equal(fmt.Sprintf("%s?test=abcd", url), actual.Url)
		a.Equal("abcd", actual.Args.Test)
		a.Equal("", actual.Headers.Hoge)
		a.Equal("", actual.Headers.Fuga)
	})
	t.Run("Getパラメータあり, ヘッダー追加あり", func(t *testing.T) {
		client := NewHttpClient(http.MethodGet, nil)
		httpBinOrgGet := &TestHttpBinOrgGet{}

		client.AddParams([]RequestParam{{"test", "abcd"}})
		client.AddHeaders([]RequestHeader{{"Hoge", "test"}})
		client.AddHeaders([]RequestHeader{{"Fuga", "test2"}})
		res, err := client.Get(url, httpBinOrgGet)
		a.NoError(err)
		a.NotNil(res)

		actual := res.Data.(*TestHttpBinOrgGet)
		a.Equal(fmt.Sprintf("%s?test=abcd", url), actual.Url)
		a.Equal("abcd", actual.Args.Test)
		a.Equal("test", actual.Headers.Hoge)
		a.Equal("test2", actual.Headers.Fuga)
	})
}
