package ttt_test

import (
  . "github.com/andrewzures/tictactoe/ttt"
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
)

var _ = Describe("Console UI", func() {
  var messages MessagesInterface
  var boardFormatter BoardFormatterInterface

  BeforeEach(func(){
    consoleMessages := new(ConsoleMessages)
    consoleMessages.BuildMessages()
    messages = MessagesInterface(consoleMessages)
    boardFormatter = BoardFormatterInterface(new(ConsoleBoardFormatter))
  })

  It("returns string representing board state", func() {
    consoleBoard := "  |   |  \n---------\n  |   |  \n---------\n  |   |  \n"
    board := GenerateBoard("x", []string{"","","","","","","","",""})
    Expect(boardFormatter.FormatBoard(board, messages)).To(Equal(consoleBoard))
  })

  It("returns board state with default x and o symbols", func() {
    consoleBoard := "X | O |  \n---------\n  |   |  \n---------\n  |   |  \n"
    board := GenerateBoard("x", []string{"x","o","","","","","","",""})
    Expect(boardFormatter.FormatBoard(board, messages)).To(Equal(consoleBoard))
  })

})


