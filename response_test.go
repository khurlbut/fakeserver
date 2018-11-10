package mockhttp_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/khurlbut/fakeserver"
)

var _ = Describe("Response Tests", func() {
	var r *Response

	BeforeEach(func() {
		r = NewResponse()
	})

	It("should not create a Nil Response", func() {
		Ω(*r).ShouldNot(BeNil())
	})

	It("should initialize the Respons Header", func() {
		Ω(r.Header).ShouldNot(BeNil())
	})

	It("should set the status code", func() {
		r.Status(200)
		Ω(r.StatusCode).Should(Equal(200))
	})

	It("should set a header value", func() {
		r.SetHeader("key", "value")
		Ω(r.Header.Get("key")).Should(Equal("value"))
	})

	It("should overwrite an existing header when using Set", func() {
		r.SetHeader("key", "value1")
		Ω(r.Header.Get("key")).Should(Equal("value1"))

		r.SetHeader("key", "value2")
		Ω(r.Header.Get("key")).Should(Equal("value2"))
	})

	It("should add a header value", func() {
		r.AddHeader("key", "value")
		Ω(r.Header.Get("key")).Should(Equal("value"))
	})

	/*
	 * This demonstrates a couple interesting points:
	 * 1. Get only returns the first value Add'd - it ignores subsequent Add's
	 * 2. To access multiple Add'd values you must interact directly with the Header map
	 * 3. Internally, keys are converted to Uppercase format
	 *    Note that "key" becomes "Key":w
	 */
	It("should append to an existing header when using Add", func() {
		r.AddHeader("key", "value1")
		Ω(r.Header.Get("key")).Should(Equal("value1"))

		r.AddHeader("key", "value2")
		Ω(r.Header.Get("key")).Should(Equal("value1"))

		Ω(r.Header["Key"][0]).Should(Equal("value1"))
		Ω(r.Header["Key"][1]).Should(Equal("value2"))
	})

	It("should add a Body (string) to the Response", func() {
		r.BodyString("this is a body string")
		Ω(string(r.BodyBuffer)).Should(Equal("this is a body string"))
	})
})
