package ttt_test

import (
	. "github.com/andrewzures/tictactoe/ttt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
  "bytes"
)

var _ = Describe("Human Player", func() {
    var human *HumanPlayer
    var player Player
    var board Board
    var writer bytes.Buffer
    var reader bytes.Buffer

    BeforeEach(func(){
      human = new(HumanPlayer)
      console := ConsoleUI{&writer, &reader}
      userInterface := UserInterface(console)
      human.NewHumanPlayer("x", "human", userInterface)
      player = Player(human)
      board = Board(new(BasicBoard))
      board.NewBoard("x")
    })

    It("meets player interface requirements", func() {
        SetMockInput(&reader, "1\n")
        Expect(player.MakeMove(board)).To(Equal(1))
    })

    It("makes a move", func() {
        SetMockInput(&reader, "1")
        Expect(player.MakeMove(board)).To(Equal(1))
    })

    It("has a symbol", func() {
        Expect(player.Symbol()).To(Equal("x"))
    })

})

