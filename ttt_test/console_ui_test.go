package ttt_test

import (
	. "github.com/andrewzures/tictactoe/ttt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
   "bytes"
)

var _ = Describe("Console UI", func() {
    var console ConsoleUI
    var factory Factory
    var writer bytes.Buffer
    var board []int

    BeforeEach(func(){
      console = ConsoleUI{&writer}
      factory = Factory(new(TTTFactory))
      board = make([]int, 9, 9)
    })

    It("prints using variable io.Writer options ", func() {
      console.Print("hello world")
      Expect(writer.String()).To(ContainSubstring("hello world\n"))
    })

    //TODO refactor with or without break? int to string
    It("prints a board", func() {
      console.DisplayBoard(board, 3)
      Expect(writer.String()).To(ContainSubstring("0 0 0 \n0 0 0 \n0 0 0 \n"))
    })

    It("lists available player types", func() {
      console.DisplayPlayerTypes(factory.PlayerTypes())
      Expect(writer.String()).To(ContainSubstring("Human"))
      Expect(writer.String()).To(ContainSubstring("Computer"))
    })

    AfterEach(func(){
      writer.Reset()
    })

})

