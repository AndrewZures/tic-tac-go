package ttt_test

import (
	. "github.com/andrewzures/tictactoe/ttt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
  "bytes"
)

var _ = Describe("Human Player", func() {
    var player Player
    var board Board
    var writer bytes.Buffer
    var reader bytes.Buffer

    BeforeEach(func(){
      inOut := InOut(ConsoleIO{&writer, &reader})
      console := buildConsoleUI(inOut)
      userInterface := UserInterface(console)

      human := new(HumanPlayer)
      human.NewHumanPlayer("x", "human", userInterface)
      player = Player(human)
      board = Generate3x3Board("x", make([]string, 9,9))
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

func buildConsoleUI(inOut InOut) ConsoleUI {
  messages := buildConsoleMessages()
  boardFormatter := BoardFormatter(new(ConsoleBoardFormatter))
  console := ConsoleUI{inOut, messages, boardFormatter}
  return console
}

func buildConsoleMessages() Messages {
  consoleMessages := new(ConsoleMessages)
  consoleMessages.BuildMessages()
  messages := Messages(consoleMessages)
  return messages
}
