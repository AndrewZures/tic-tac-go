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
  var messages Messages
  var rules Rules
  var writer bytes.Buffer
  var reader bytes.Buffer

  BeforeEach(func(){
    rules = new(BasicRules)

    inOut := InOut(ConsoleIO{&writer, &reader})
    consoleMessages := new(ConsoleMessages)
    consoleMessages.BuildMessages()
    messages = Messages(consoleMessages)
    boardFormatter := BoardFormatter(new(ConsoleBoardFormatter))

    console = ConsoleUI{inOut, messages, boardFormatter}
    factory = Factory(new(TTTFactory))
  })


    It("prints intro message", func() {
      console.DisplayIntroMessage()
      Expect(writer.String()).To(ContainSubstring("Welcome to Tic Tac Go"))
    })

    It("prints exit message", func() {
      console.DisplayExitMessage()
      Expect(writer.String()).To(ContainSubstring("So Long!"))
    })


  Context("when receiving input from user", func() {

    It("validates input from use is integer", func() {
      SetMockInput(&reader, "3")
      Expect(console.GetIntegerFromUser()).To(Equal(3))
    })

    It("requeries user until integer is provided", func() {
      SetMockInput(&reader, "firsstring\nsecond123string2\n1\n")
      Expect(console.GetIntegerFromUser()).To(Equal(1))
    })

    It("reads multiple integer inputs from user", func() {
      SetMockInput(&reader, "firsstring\nsecond123string2\n3\n6\n")
      Expect(console.GetIntegerFromUser()).To(Equal(3))
      Expect(console.GetIntegerFromUser()).To(Equal(6))
    })
  })

  Context("when displaying information to user", func() {

    It("indicates invalid choice to user", func() {
      console.PrintChoiceInvalid()
      Expect(writer.String()).To(ContainSubstring("Whoops, that choice is invalid"))
      Expect(writer.String()).To(ContainSubstring("Try Again"))
    })

    It("indicates invalid choice to user if integer expected and not received", func() {
      SetMockInput(&reader, "firsstring\n1\n")
      console.GetIntegerFromUser()
      Expect(writer.String()).To(ContainSubstring("Whoops, that choice is invalid"))
      Expect(writer.String()).To(ContainSubstring("Try Again"))
    })

    It("Displays Game Winner", func() {
      console.DisplayWinner("Player 1")
      Expect(writer.String()).To(ContainSubstring("Player 1"))
    })

    It("Displays Tie Game", func() {
      console.DisplayWinner("tie")
      Expect(writer.String()).To(ContainSubstring("Tie"))
    })

    It("Asks player for move", func() {
      player := new(HumanPlayer)
      player.NewHumanPlayer("X", "Player 1", console)
      console.AskUserForMove(player)
      Expect(writer.String()).To(ContainSubstring("Choose a Move"))

    })

    It("Asks user for board type", func() {
      console.PrintBoardTypeQuestion()
      Expect(writer.String()).To(ContainSubstring("Choose Board Type:"))
    })

    It("Asks user for player type", func() {
      console.PrintPlayerTypeQuestion("Player 1")
      Expect(writer.String()).To(ContainSubstring("Choose Type for: Player 1"))
    })

    It("Asks user for new game", func() {
      SetMockInput(&reader, "y\nn\n")
      result := console.AskForNewGame()
      Expect(writer.String()).To(ContainSubstring("Would you like to start a new game"))
      Expect(result).To(Equal(true))

      result = console.AskForNewGame()
      Expect(result).To(Equal(false))
    })

    Context ("when displaying Player Type Information", func() {

      It("lists available player types", func() {
        console.DisplayPlayerTypes(factory.PlayerTypes(console, rules))
        Expect(writer.String()).To(ContainSubstring("Human"))
        Expect(writer.String()).To(ContainSubstring("Computer"))
      })
    })

    Context ("when displaying Board Type Information", func() {

      It("lists available board types", func() {
        console.DisplayBoardTypes(factory.BoardTypes())
        Expect(writer.String()).To(ContainSubstring("3x3 Board"))
        Expect(writer.String()).To(ContainSubstring("4x4 Board"))
        Expect(writer.String()).To(ContainSubstring("5x5 Board"))
      })

    })

    It ("validates chosen move against available moves", func() {
        board := Generate3x3Board([]string{"x","o","x","o","","x","","o",""})
        result := console.ValidateMove(8, board)
        Expect(result).To(Equal(true))

        result = console.ValidateMove(1, board)
        Expect(result).To(Equal(false))
      })

  })

  Context("when communicating with program", func() {
    var playerTypes []Player
    var boardTypes []Board

    BeforeEach(func() {
      playerTypes = factory.PlayerTypes(console, rules)
      boardTypes = factory.BoardTypes()
    })

    It("returns selected player choice", func() {
      SetMockInput(&reader, "1")
      selectedPlayer := console.SelectPlayerChoice(playerTypes, "player 1")
      Expect(selectedPlayer.Description()).To(Equal("Human"))

      SetMockInput(&reader, "2")
      selectedPlayer = console.SelectPlayerChoice(playerTypes, "player 1")
      Expect(selectedPlayer.Description()).To(Equal("Computer"))
    })

    It("only returns valid player choice", func() {
      SetMockInput(&reader, "-1\n100\na\n2\n1\n")
      selectedPlayer := console.SelectPlayerChoice(playerTypes, "player 1")
      Expect(selectedPlayer.Description()).To(Equal("Computer"))
    })

    It("returns multiptle players", func() {
      SetMockInput(&reader, "a\n1\n2\n2\n")
      selectedPlayer := console.SelectPlayerChoice(playerTypes, "player 1")
      Expect(selectedPlayer.Description()).To(Equal("Human"))

      anotherPlayer := console.SelectPlayerChoice(playerTypes, "player 2")
      Expect(anotherPlayer.Description()).To(Equal("Computer"))
    })

    It("returns board type template", func() {
      SetMockInput(&reader, "1\n")
      console.SelectBoardChoice(boardTypes)
      Expect(writer.String()).To(ContainSubstring("3x3 Board"))
    })

    It("re-queries user if choice not within range", func() {
      SetMockInput(&reader, "5\n1\n")
      console.SelectBoardChoice(boardTypes)
      Expect(writer.String()).To(ContainSubstring("Whoops, that choice is invalid"))
    })

  })

  AfterEach(func(){
    writer.Reset()
    reader.Reset()
  })

})


