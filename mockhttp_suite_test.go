package mockhttp_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMockHttp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Package 'github.com/khurlbut/mockhttp'")
}
