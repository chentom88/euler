package problem1_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestProblem1(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Problem1 Suite")
}
