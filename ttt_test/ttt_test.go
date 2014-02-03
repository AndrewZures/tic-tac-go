package ttt_test

import (
	"github.com/andrewzures/tictactoe/ttt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ttt", func() {

    It("adds two numbers", func() {
        Expect(ttt.Add(2,3)).To(Equal(5))
    })

    It("adds two numbers also", func() {
        Expect(ttt.Add(4,5)).To(Equal(9))
    })

})
