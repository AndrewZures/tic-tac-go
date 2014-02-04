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

    It("prints a board", func() {
      console.PrintBoard(board)
      Expect(writer.String()).To(ContainSubstring("value at 1 is 0\n"))
    })

})

