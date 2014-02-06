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
    var buffer bytes.Buffer
    var board []int

    BeforeEach(func(){
      console = ConsoleUI{&buffer, &buffer}
      factory = Factory(new(TTTFactory))
      board = make([]int, 9, 9)
    })

    It("test mock input", func() {
      fmt.Fprintf(&buffer, "Hello")
      result := console.TestIO()
      Expect(result).To(Equal("Hello"))

    })

    It("prints using variable io.buffer options ", func() {
      console.Print("hello world")
      Expect(buffer.String()).To(ContainSubstring("hello world\n"))
    })

    //TODO refactor with or without break? int to string
    It("prints a board", func() {
      console.DisplayBoard(board, 3)
      Expect(buffer.String()).To(ContainSubstring("0 0 0 \n0 0 0 \n0 0 0 \n"))
    })

    It("lists available player types", func() {
      console.DisplayPlayerTypes(factory.PlayerTypes())
      Expect(buffer.String()).To(ContainSubstring("Human"))
      Expect(buffer.String()).To(ContainSubstring("Computer"))
    })

    It("returns selected player choice", func() {
      playerTypes := factory.PlayerTypes()
      selectedPlayer := console.SelectPlayerChoice(playerTypes)
      Expect(selectedPlayer.Description()).To(Equal("Human"))

    })

    AfterEach(func(){
      buffer.Reset()
    })

})

func setBufferText (text string) {


}

