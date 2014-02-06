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
  var board []int

  BeforeEach(func(){
    console = ConsoleUI{&writer, *bufio.NewReader(&reader)}
    factory = Factory(new(TTTFactory))
    board = make([]int, 9, 9)
  })

  Context("when receiving input from user", func() {

    It("reads from console", func() {
      fmt.Fprintf(&reader, "Test Console Input")
      result := console.ReadConsole()
      Expect(result).To(Equal("Test Console Input"))
    })

    It("validates input from use is integer", func() {
      fmt.Fprintf(&reader, "3")
      Expect(console.GetIntegerFromUser()).To(Equal(3))
    })

    It("requeries user until integer is provided", func() {
      fmt.Fprintf(&reader, "firsstring\nsecondstring2\n1\n")
      Expect(console.GetIntegerFromUser()).To(Equal(1))

    })
  })

  Context("when sending information to user", func() {

    It("indicates invalid choice to user and asks for another choice", func() {
      console.PrintChoiceInvalid()
      Expect(writer.String()).To(ContainSubstring("Whoops, that choice is invalid"))
      Expect(writer.String()).To(ContainSubstring("Try Again"))
    })

    It("indicates invalid choice to user if integer expected and not received", func() {
      fmt.Fprintf(&reader, "firsstring\n1\n")
      console.GetIntegerFromUser()
      Expect(writer.String()).To(ContainSubstring("Whoops, that choice is invalid"))
      Expect(writer.String()).To(ContainSubstring("Try Again"))

    })

    It("lists available player types", func() {
      console.DisplayPlayerTypes(factory.PlayerTypes())
      Expect(writer.String()).To(ContainSubstring("Human"))
      Expect(writer.String()).To(ContainSubstring("Computer"))
    })

    //TODO refactor with or without break? int to string
    It("prints a board", func() {
      console.DisplayBoard(board, 3)
      Expect(writer.String()).To(ContainSubstring("0 0 0 \n0 0 0 \n0 0 0 \n"))
    })
  })

  Context("when communicating with rest of program", func() {

    It("returns selected player choice", func() {
      fmt.Fprintf(&reader, "1")
      playerTypes := factory.PlayerTypes()
      selectedPlayer := console.SelectPlayerChoice(playerTypes)
      Expect(selectedPlayer.Description()).To(Equal("Human"))
    })

  })

  AfterEach(func(){
    writer.Reset()
    reader.Reset()
  })

})
