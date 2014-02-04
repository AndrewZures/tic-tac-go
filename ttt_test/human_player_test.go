package ttt_test

import (
	. "github.com/andrewzures/tictactoe/ttt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

)

var _ = Describe("Human Player", func() {
    var human *HumanPlayer
    var player Player
    var board []string

    BeforeEach(func(){
      human = new(HumanPlayer)
      human.SetSymbol("X")
      player = Player(human)
      board = make([]string, 9, 9)
    })

    It("meets player interface requirements", func() {
        Expect(player.MakeMove(board)).To(Equal(1))
    })

    It("makes a move", func() {
        Expect(player.MakeMove(board)).To(Equal(1))
    })

    It("has a symbol", func() {
        Expect(player.Symbol()).To(Equal("X"))
    })

})
