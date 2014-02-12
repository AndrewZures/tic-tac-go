package ttt_test

import (
  . "github.com/andrewzures/tictactoe/ttt"
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  "bytes"
)

var _ = Describe("Game Test", func() {
  var messages Messages
  var writer bytes.Buffer
  var reader bytes.Buffer
  var factory Factory
  var console UserInterface
  var game Game

  BeforeEach(func(){
    inOut := InOut(ConsoleIO{&writer, &reader})
    consoleMessage := new(ConsoleMessages)
    messages = Messages(consoleMessage)
    messages.BuildMessages()
    boardFormatter := BoardFormatter(new(ConsoleBoardFormatter))
    consoleui := ConsoleUI{inOut, messages, boardFormatter}
    console = UserInterface(consoleui)
    factory = Factory(new(TTTFactory))

    game = new(TTTGame)
  })

  It("has choose move response", func() {
    completeHumanGame := "1\n1\n1\n1\n2\n3\n4\n5\n6\n7\nn\n"
    SetMockInput(&reader, completeHumanGame)
    game.Run(console, factory)
    Expect(writer.String()).To(ContainSubstring("Winner is X"))
  })

  It("has choose move response", func() {
    completeComputerGame := "2\n2\n1\nn\n"
    SetMockInput(&reader, completeComputerGame)
    game.Run(console, factory)
    Expect(writer.String()).To(ContainSubstring(messages.GameTieResponse()))
  })

  It("displays empty board for human/human game", func() {
    completeHumanGame := "1\n1\n1\n1\n2\n3\n4\n5\n6\n7\nn\n"
    SetMockInput(&reader, completeHumanGame)
    game.Run(console, factory)
    Expect(writer.String()).To(ContainSubstring( "  |   |  \n---------\n  |   |  \n---------\n  |   |  \n" ))
  })

})
