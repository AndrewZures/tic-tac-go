package ttt_test

import (
	. "github.com/andrewzures/tictactoe/ttt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Human Player", func() {
    var human *HumanPlayer
    var board []int

    BeforeEach(func(){
      human = new(HumanPlayer)
      human.SetSymbol("X")
      board = make([]int, 9, 9)
    })

    It("makes a move", func() {
        Expect(human.MakeMove(board)).To(Equal(1))
    })

    It("has a symbol", func() {
        Expect(human.GetSymbol()).To(Equal("X"))
    })

    It("can be interfaced with Player", func() {
        Expect(human.MakeMove(board)).To(Equal(1))
    })

    It("can be interfaced with Player again", func() {
        Expect(human.MakeMove(board)).To(Equal(1))
    })
})
