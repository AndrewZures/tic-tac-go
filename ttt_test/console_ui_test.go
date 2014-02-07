package ttt_test

import (
  . "github.com/andrewzures/tictactoe/ttt"
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  "bytes"
  "fmt"
  "bufio"
)

var _ = Describe("Console UI", func() {
  var console ConsoleUI
  var factory Factory
  var writer bytes.Buffer
  var reader bytes.Buffer

  BeforeEach(func(){
    console = ConsoleUI{&writer, *bufio.NewReader(&reader)}
    factory = Factory(new(TTTFactory))
  })

  It("implements UserInterface", func() {
    playerTypes := factory.PlayerTypes()
    interfacedConsole := UserInterface(console)
    setMockInput(&reader, "1")
    selectedPlayer := interfacedConsole.SelectPlayerChoice(playerTypes, "player 1")
    Expect(selectedPlayer.Description()).To(Equal("Human"))
  })

  Context("when receiving input from user", func() {

    It("reads from console", func() {
      setMockInput(&reader, "Test Console Input")
      result := console.ReadConsole()
      Expect(result).To(Equal("Test Console Input"))
    })

    It("validates input from use is integer", func() {
      setMockInput(&reader, "3")
      Expect(console.GetIntegerFromUser()).To(Equal(3))
    })

    It("requeries user until integer is provided", func() {
      setMockInput(&reader, "firsstring\nsecond123string2\n1\n")
      Expect(console.GetIntegerFromUser()).To(Equal(1))

    })
  })

  Context("when displaying information to user", func() {

    It("indicates invalid choice to user and asks for another choice", func() {
      console.PrintChoiceInvalid()
      Expect(writer.String()).To(ContainSubstring("Whoops, that choice is invalid"))
      Expect(writer.String()).To(ContainSubstring("Try Again"))
    })

    It("indicates invalid choice to user if integer expected and not received", func() {
      setMockInput(&reader, "firsstring\n1\n")
      console.GetIntegerFromUser()
      Expect(writer.String()).To(ContainSubstring("Whoops, that choice is invalid"))
      Expect(writer.String()).To(ContainSubstring("Try Again"))

    })

    It("lists available player types", func() {
      console.DisplayPlayerTypes(factory.PlayerTypes())
      Expect(writer.String()).To(ContainSubstring("Human"))
      Expect(writer.String()).To(ContainSubstring("Computer"))
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

    BeforeEach(func() {
      playerTypes = factory.PlayerTypes()
    })

    It("returns selected player choice", func() {
      setMockInput(&reader, "1")
      selectedPlayer := console.SelectPlayerChoice(playerTypes, "player 1")
      Expect(selectedPlayer.Description()).To(Equal("Human"))

      setMockInput(&reader, "2")
      selectedPlayer = console.SelectPlayerChoice(playerTypes, "player 1")
      Expect(selectedPlayer.Description()).To(Equal("Computer"))
    })

    It("only returns valid player choice", func() {
      setMockInput(&reader, "-1\n100\na\n2\n")
      selectedPlayer := console.SelectPlayerChoice(playerTypes, "player 1")
      Expect(selectedPlayer.Description()).To(Equal("Computer"))
    })

  })

  AfterEach(func(){
    writer.Reset()
    reader.Reset()
  })

})

func setMockInput(reader *bytes.Buffer, input string){
    fmt.Fprintf(reader, input)
}

