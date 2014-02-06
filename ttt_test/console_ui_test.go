package ttt_test

import (
	. "github.com/andrewzures/tictactoe/ttt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
   "bytes"
   "fmt"
)

var _ = Describe("Console UI", func() {
    var console ConsoleUI
    var factory Factory
    var writer bytes.Buffer
    var reader bytes.Buffer
    var board []int

    BeforeEach(func(){
      console = ConsoleUI{&writer, &reader}
      factory = Factory(new(TTTFactory))
      board = make([]int, 9, 9)
    })

    It("reads in data", func() {
      fmt.Fprintf(&reader, "Test Console Input")
      result := console.ReadConsole()
      Expect(result).To(Equal("Test Console Input"))
    })

    //TODO refactor with or without break? int to string
    It("prints a board", func() {
      console.DisplayBoard(board, 3)
      Expect(writer.String()).To(ContainSubstring("0 0 0 \n0 0 0 \n0 0 0 \n"))
    })

    It("lists available player types", func() {
      console.DisplayPlayerTypes(factory.PlayerTypes())
      Expect(writer.String()).To(ContainSubstring("Human"))
      Expect(writer.String()).To(ContainSubstring("Computer"))
    })

    It("returns selected player choice", func() {
      fmt.Fprintf(&reader, "1")
      playerTypes := factory.PlayerTypes()
      selectedPlayer := console.SelectPlayerChoice(playerTypes)
      Expect(selectedPlayer.Description()).To(Equal("Human"))
    })

    AfterEach(func(){
      writer.Reset()
      reader.Reset()
    })

})
