package fakeserver_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/khurlbut/fakeserver"
	"io/ioutil"
	"net/http"
)

var _ = Describe("HTTP Fake Tests", func() {
	var fake *HTTPFake

	BeforeEach(func() {
		fake = New()
		fake.Start()
	})

	AfterEach(func() {
		fake.Server.Close()
	})

	It("should not be nil", func() {
		Ω(*fake).ShouldNot(BeNil())
	})

	It("should intialize empty request handlers array", func() {
		Ω(fake.RequestHandlers).ShouldNot(BeNil())
		Ω(len(fake.RequestHandlers)).Should(BeZero())
	})

	It("should initialize Server", func() {
		Ω(fake.Server).ShouldNot(BeNil())
	})

	It("should add a new Request to the array of Request Handlers", func() {
		r := fake.NewHandler()
		Ω(len(fake.RequestHandlers)).ShouldNot(BeZero())
		Ω(fake.RequestHandlers[0]).Should(Equal(r))
	})

	/*
	 * This test demonstrates that the http.Server will generate a url pointing to
	 * localhost with a random port (5 digits).
	 *
	 * URL is: http://127.0.0.1:\d{5}/path/to/page?param1=value1
	 */
	It("should resolve the full URL to the fake server for a given path", func() {
		resolvedURL := fake.ResolveURL("%s?%s=%s", "/path/to/page", "param1", "value1")
		Ω(resolvedURL).Should(MatchRegexp("http:\\/\\/127\\.0\\.0\\.1:8181\\/path\\/to\\/page\\?param1=value1"))
		// Ω(resolvedURL).Should(MatchRegexp("http:\\/\\/127\\.0\\.0\\.1:\\d{5}\\/path\\/to\\/page\\?param1=value1"))
	})

	It("should reset the Request Handler definitions", func() {
		fake.NewHandler()
		fake.Reset()
		Ω(len(fake.RequestHandlers)).Should(BeZero())
	})

	It("should return the expected response on GET", func() {
		fake.NewHandler().Get("/users").Reply(200).BodyString(`[{"username": "dreamer"}]`)

		res, _ := http.Get(fake.ResolveURL("/users"))
		defer res.Body.Close()

		body, _ := ioutil.ReadAll(res.Body)

		Ω(res.StatusCode).Should(Equal(200))
		Ω(string(body)).Should(Equal(`[{"username": "dreamer"}]`))

	})

	It("should return 404", func() {
		res, _ := http.Get(fake.ResolveURL("/path/to/nowhere"))
		defer res.Body.Close()
		Ω(res.StatusCode).Should(Equal(404))
	})

	PIt("should just do stuff...", func() {
	})
})
