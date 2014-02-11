package ttt_test

import (
  . "github.com/andrewzures/tictactoe/ttt"
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
)

var _ = Describe("Console UI", func() {
  var messages Messages
  var boardFormatter BoardFormatter

  BeforeEach(func(){
    consoleMessages := new(ConsoleMessages)
    consoleMessages.BuildMessages()
    messages = Messages(consoleMessages)
    boardFormatter = BoardFormatter(new(ConsoleBoardFormatter))
  })

  It("returns string representing board state", func() {
    consoleBoard := "  |   |  \n---------\n  |   |  \n---------\n  |   |  \n"
    board := Generate3x3Board("x", []string{"","","","","","","","",""})
    Expect(boardFormatter.FormatBoard(board, messages)).To(Equal(consoleBoard))
  })

  It("returns board state with default x and o symbols", func() {
    consoleBoard := "X | O |  \n---------\n  |   |  \n---------\n  |   |  \n"
    board := Generate3x3Board("x", []string{"x","o","","","","","","",""})
    Expect(boardFormatter.FormatBoard(board, messages)).To(Equal(consoleBoard))
  })

  It("returns properly formatted 4x4 board", func() {
    consoleBoard := "  |   |   |  \n------------\n  |   |   |  \n------------\n  |   |   |  \n------------\n  |   |   |  \n"
    board := GenerateBoard(16, 4, "x", "4x4 Board", []string{"","","","","","","","","","","","","","","",""})
    Expect(boardFormatter.FormatBoard(board, messages)).To(Equal(consoleBoard))
  })

})


