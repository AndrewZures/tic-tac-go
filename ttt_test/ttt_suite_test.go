package ttt_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestTtt(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Ttt Suite")
}
