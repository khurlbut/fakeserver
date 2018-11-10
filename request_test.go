package http_test_wrapper_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "kdh.com/http_test_wrapper"
)

var _ = Describe("Request Tests", func() {
	var r *Request

	BeforeEach(func() {
		r = NewRequest()
	})

	It("should not create a Nil Request", func() {
		Ω(*r).ShouldNot(BeNil())
	})

	It("should initialize a URL pointer", func() {
		Ω(r.URL).ShouldNot(BeNil())
	})

	It("should initialize a Response pointer", func() {
		Ω(r.Response).ShouldNot(BeNil())
	})

	It("should mutate into a GET request with a path to GET", func() {
		r.Get("path/to/get")
		Ω(r.Method).Should(Equal("GET"))
		Ω(r.URL.Path).Should(Equal("path/to/get"))
	})

	It("should mutate into a POST request with a path to POST", func() {
		r.Post("path/to/post")
		Ω(r.Method).Should(Equal("POST"))
		Ω(r.URL.Path).Should(Equal("path/to/post"))
	})

	It("should mutate into a PUT request with a path to PUT", func() {
		r.Put("path/to/put")
		Ω(r.Method).Should(Equal("PUT"))
		Ω(r.URL.Path).Should(Equal("path/to/put"))
	})

	It("should mutate into a PATCH request with a path to PATCH", func() {
		r.Patch("path/to/patch")
		Ω(r.Method).Should(Equal("PATCH"))
		Ω(r.URL.Path).Should(Equal("path/to/patch"))
	})

	It("should mutate into a DELETE request with a path to DELETE", func() {
		r.Delete("path/to/delete")
		Ω(r.Method).Should(Equal("DELETE"))
		Ω(r.URL.Path).Should(Equal("path/to/delete"))
	})

	It("should mutate into a HEAD request with a path to HEAD", func() {
		r.Head("path/to/head")
		Ω(r.Method).Should(Equal("HEAD"))
		Ω(r.URL.Path).Should(Equal("path/to/head"))
	})

	It("should set a custom Responder onto the CustomHandler", func() {
		Ω(r.CustomHandle).Should(BeNil())
		r.Handle(DefaultResponder)
		Ω(r.CustomHandle).ShouldNot(BeNil())
	})

	Specify("that Reply will set a status on the contained Response", func() {
		res := r.Reply(200)
		Ω(res.StatusCode).Should(Equal(200))
	})

})
