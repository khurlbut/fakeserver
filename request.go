package http_test_wrapper

import (
	"net/url"
	"strings"
)

type Request struct {
	Method       string
	URL          *url.URL
	Response     *Response
	CustomHandle Responder
}

func NewRequest() *Request {
	return &Request{
		URL:      &url.URL{},
		Response: NewResponse(),
	}
}

func (r *Request) Get(path string) *Request {
	return r.method("GET", path)
}

func (r *Request) Post(path string) *Request {
	return r.method("POST", path)
}

func (r *Request) Put(path string) *Request {
	return r.method("PUT", path)
}

func (r *Request) Patch(path string) *Request {
	return r.method("PATCH", path)
}

func (r *Request) Delete(path string) *Request {
	return r.method("DELETE", path)
}

func (r *Request) Head(path string) *Request {
	return r.method("HEAD", path)
}

func (r *Request) Handle(handle Responder) {
	r.CustomHandle = handle
}

func (r *Request) Reply(status int) *Response {
	r.Response.Status(status)
	return r.Response
}

func (r *Request) method(method, path string) *Request {
	r.URL.Path = normalize(path)
	r.Method = strings.ToUpper(method)
	return r
}

func normalize(p string) string {
	if strings.HasSuffix(p, "/") {
		return p
	}
	return p + "/"
}
