package mockhttp

import "net/http"

type Response struct {
	StatusCode int
	Header     http.Header
	BodyBuffer []byte
}

func NewResponse() *Response {
	return &Response{
		Header: make(http.Header),
	}
}

func (r *Response) Status(status int) *Response {
	r.StatusCode = status
	return r
}

func (r *Response) SetHeader(key string, value string) *Response {
	r.Header.Set(key, value)
	return r
}

func (r *Response) AddHeader(key string, value string) *Response {
	r.Header.Add(key, value)
	return r
}

func (r *Response) BodyString(body string) *Response {
	r.BodyBuffer = []byte(body)
	return r
}
