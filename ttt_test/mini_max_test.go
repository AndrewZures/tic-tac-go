package ttt_test

import (
	. "github.com/andrewzures/tictactoe/ttt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Computer Player", func() {
    var computer *ComputerPlayer

    BeforeEach(func(){
      rules := Rules(new(BasicRules))
      computer = new(ComputerPlayer)
      computer.NewComputerPlayer("o", "Computer", rules)
    })

      It("has a symbol", func() {
        Expect(computer.Symbol()).To(Equal("o"))
      })

    Context("when scoring a 4x4 board", func() {

        It("takes index 13 when it's a winner", func() {
          board := Generate4x4Board([]string{"x","o","x","",
                                             "","o","","",
                                             "x","o","","x",
                                             "","","",""})
          Expect(computer.MakeMove(board)).To(Equal(13))
        })

        It("takes index 3 when it's a winner", func() {
          board := Generate4x4Board([]string{"o","o","o","",
                                             "x","x","x","",
                                             "","","","",
                                             "","","",""})
          Expect(computer.MakeMove(board)).To(Equal(3))
        })

        It("takes index 7 when it's a winner", func() {
          board := Generate4x4Board([]string{"x","x","x","",
                                             "o","o","o","",
                                             "","","","",
                                             "","","",""})
          Expect(computer.MakeMove(board)).To(Equal(7))
        })

        It("takes index 7 when it's a winner", func() {
          board := Generate4x4Board([]string{"o","x","x","",
                                             "x","o","o","",
                                             "","","o","",
                                             "","","",""})
          Expect(computer.MakeMove(board)).To(Equal(15))
        })

        It("takes index 3 to defend against immediate loss", func() {
          board := Generate4x4Board([]string{"x","x","x","",
                                             "o","o","","",
                                             "o","","","",
                                             "","","",""})
          Expect(computer.MakeMove(board)).To(Equal(3))
        })

        It("takes index 3 to defend against immediate loss", func() {
          board := Generate4x4Board([]string{"","","","",
                                             "o","o","x","",
                                             "o","x","","",
                                             "x","","",""})
          Expect(computer.MakeMove(board)).To(Equal(3))
        })

      })

    })

