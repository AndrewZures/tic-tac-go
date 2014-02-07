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
  var reader bytes.Buffer

  BeforeEach(func(){
    console = ConsoleUI{&writer, &reader}
    factory = Factory(new(TTTFactory))
  })

  It("implements UserInterface", func() {
    playerTypes := factory.PlayerTypes()
    interfacedConsole := UserInterface(console)
    SetMockInput(&reader, "1")
    selectedPlayer := interfacedConsole.SelectPlayerChoice(playerTypes, "player 1")
    Expect(selectedPlayer.Description()).To(Equal("Human"))
  })

  Context("when receiving input from user", func() {

    XIt("reads from console", func() {
      SetMockInput(&reader, "Test Console Input")
      result := console.ReadConsole()
      Expect(result).To(Equal("Test Console Input"))
    })

    It("reads multiple times from console", func() {
      SetMockInput(&reader, "Test\nConsole\nInput\n")
      Expect(console.ReadConsole()).To(Equal("Test"))
      Expect(console.ReadConsole()).To(Equal("Console"))
      Expect(console.ReadConsole()).To(Equal("Input"))
    })

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

    Context ("when displaying Player Type Information", func() {

      It("lists available player types", func() {
        console.DisplayPlayerTypes(factory.PlayerTypes())
        Expect(writer.String()).To(ContainSubstring("Human"))
        Expect(writer.String()).To(ContainSubstring("Computer"))
      })
    })

    Context ("when displaying Board Type Information", func() {

      It("lists available board types", func() {
        console.DisplayBoardTypes(factory.BoardTypes())
        Expect(writer.String()).To(ContainSubstring("3x3 Board"))
      })

    })

    Context ("when displaying game board", func() {

      It("prints a board", func() {
        board := GenerateBoard("x", []string{"","","","","","","","",""})
        console.DisplayBoard(board)
        Expect(writer.String()).To(ContainSubstring("  |   |  \n"))
      })

      It("prints current gamestate", func() {
        board := GenerateBoard("x", []string{"","x","o","o","o","x","","x",""})
        console.DisplayBoard(board)
        Expect(writer.String()).To(ContainSubstring(  "  | X | O\n"))
        Expect(writer.String()).To(ContainSubstring("\n---------\n"))
        Expect(writer.String()).To(ContainSubstring("\nO | O | X\n"))
        Expect(writer.String()).To(ContainSubstring("\n  | X |  \n"))
      })
    })

  })

  Context("when communicating with rest of program", func() {
    var playerTypes []Player
    var boardTypes []Board

    BeforeEach(func() {
      playerTypes = factory.PlayerTypes()
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

    It("returns board type template", func() {
      SetMockInput(&reader, "2\n1\n")
      console.SelectBoardChoice(boardTypes)
      Expect(writer.String()).To(ContainSubstring("Whoops, that choice is invalid"))
    })

  })

  AfterEach(func(){
    writer.Reset()
    reader.Reset()
  })

})


