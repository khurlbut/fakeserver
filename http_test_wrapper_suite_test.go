package mockhttp_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestHttpTestWrapper(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "HttpTestWrapper Suite")
}
