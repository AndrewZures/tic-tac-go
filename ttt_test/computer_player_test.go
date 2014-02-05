package ttt_test

import (
	. "github.com/andrewzures/tictactoe/ttt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

)

var _ = Describe("Computer Player", func() {
    var computer *ComputerPlayer
    var player Player
    var board Board

    BeforeEach(func(){
      computer = new(ComputerPlayer)
      computer.NewComputerPlayer("o", "Computer")
      board = Board(new(BasicBoard))
      board.NewBoard("o")
    })

    Context("Basic Player Attributes", func() {

      It("has a symbol", func() {
        Expect(computer.Symbol()).To(Equal("o"))
      })

      It("meets player interface requirements", func() {
        player = Player(computer)
        Expect(player.Symbol()).To(Equal("o"))
      })

    })

    Context("Minimax Implementation", func() {

      Context("takes win if available", func() {

        It("takes index 7 when it's a winner", func() {
          board := GenerateBoard("o", []string{"x","o","x","","o","","","","x"})
          Expect(board.PlayerTurn()).To(Equal("o"))
          Expect(computer.MakeMove(board)).To(Equal(7))
        })

        It("takes index 5 when it's a winner", func() {
          board := GenerateBoard("o", []string{"x","","x","o","o","","","","x"})
          Expect(computer.MakeMove(board)).To(Equal(5))
        })

        It("takes index 3 when it's a winner", func() {
          board := GenerateBoard("o", []string{"x","","x","","o","o","x","",""})
          Expect(computer.MakeMove(board)).To(Equal(3))
        })

        It("takes index 8 when it's a winner", func() {
          board := GenerateBoard("o", []string{"o","x","x","x","o","","","",""})
          Expect(player.MakeMove(board)).To(Equal(8))
        })
      })

      Context("defends against immediate loss", func() {

        It("takes index 8 when it blocks opponent win", func() {
          board := GenerateBoard("o", []string{"x","o","x","x","x","o","o","",""})
          Expect(computer.MakeMove(board)).To(Equal(8))
        })

        It("takes index 3 when it blocks opponent win", func() {
          board := GenerateBoard("o", []string{"x","","","","","o","x","",""})
          Expect(computer.MakeMove(board)).To(Equal(3))
        })

        It("takes index 1 when it blocks opponent win", func() {
          board := GenerateBoard("o", []string{"x","","x","o","","","","",""})
          Expect(computer.MakeMove(board)).To(Equal(1))
        })

      })

    })

  })

