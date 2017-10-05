package problem2_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestProblem2(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Problem2 Suite")
}
