package ttt_test

import (
  . "github.com/andrewzures/tictactoe/ttt"
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
)

var _ = Describe("Console UI", func() {
  var messages MessagesInterface

  BeforeEach(func(){
    consoleMessages := new(ConsoleMessages)
    consoleMessages.BuildMessages()
    messages = MessagesInterface(consoleMessages)
  })

  It("returns string representing board state", func() {
    consoleBoard := "  |   |  \n---------\n  |   |  \n---------\n  |   |  \n"
    board := GenerateBoard("x", []string{"","","","","","","","",""})
    Expect(messages.DisplayBoard(board)).To(Equal(consoleBoard))
  })

})


