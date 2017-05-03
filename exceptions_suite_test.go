package exceptions_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestExceptions(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Exceptions Suite")
}
