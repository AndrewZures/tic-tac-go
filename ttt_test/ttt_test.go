package ttt_test

import (
  "github.com/andrewzures/tictactoe/ttt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ttt", func() {

    It("adds two numbers", func() {
        Expect(ttt.Add(6,5)).To(Equal(5))
    })

})
