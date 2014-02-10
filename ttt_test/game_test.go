package ttt_test

import (
  . "github.com/andrewzures/tictactoe/ttt"
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  "bytes"
)

var _ = Describe("Game", func() {
  var userInterface UserInterface
  var factory Factory
  var game Game
  var messages MessagesInterface
  var writer bytes.Buffer
  var reader bytes.Buffer
  var inOut InOutInterface

  BeforeEach(func(){
    inOut = InOutInterface(ConsoleIO{&writer, &reader})
    consoleMessages := new(ConsoleMessages)
    consoleMessages.BuildMessages()
    messages = MessagesInterface(consoleMessages)
    console := ConsoleUI{inOut, messages}
    userInterface = UserInterface(console)
    factory = Factory(new(TTTFactory))
    game = Game(new(TTTGame))
  })

  It("sets up a game", func() {
    SetMockInput(&reader, "1\n2\n1\n")
    emptyBoardArray := []string{"","","","","","","","",""}

    board, player1, player2 := game.SetupNewGame(userInterface,factory)

    Expect(player1.Description()).To(Equal("Human"))
    Expect(player2.Description()).To(Equal("Computer"))
    Expect(board.Description()).To(Equal("3x3 Board"))
    Expect(board.Array()).To(Equal(emptyBoardArray))
  })

})
