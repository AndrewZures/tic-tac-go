package ttt_test

import (
	. "github.com/andrewzures/tictactoe/ttt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
   "bytes"
)

var _ = Describe("Console UI", func() {
    var console ConsoleUI
    var writer bytes.Buffer
    var board []int

    BeforeEach(func(){
      console = ConsoleUI{&writer}
      board = make([]int, 9, 9)
    })

    It("prints using variable io.Writer options ", func() {
      console.Print("hello world")
      Expect(writer.String()).To(ContainSubstring("hello world\n"))
    })

    //TODO refactor with or without break?
    It("prints a board", func() {
      console.PrintBoard(board, 3)
      Expect(writer.String()).To(ContainSubstring("0 0 0 \n0 0 0 \n0 0 0 \n"))
    })

    AfterEach(func(){
      writer.Reset()
    })

})

