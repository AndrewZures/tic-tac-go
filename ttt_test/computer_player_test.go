package ttt_test

import (
	. "github.com/andrewzures/tictactoe/ttt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

)

var _ = Describe("Computer Player", func() {
    var computer *ComputerPlayer
    var player Player
    var board []string

    BeforeEach(func(){
      computer = new(ComputerPlayer)
      computer.SetSymbol("X")
      player = Player(computer)
      board = make([]string, 9, 9)
    })

    Context("Basic Player Attributes", func() {

      It("meets player interface requirements", func() {
        Expect(player.MakeMove(board)).To(Equal(1))
      })

      It("has a symbol", func() {
        Expect(player.Symbol()).To(Equal("X"))
      })
    })

    Context("Utility Functions", func() {

      It("compares two different strings", func() {
        string1 := "x"
        string2 := "o"
        Expect(computer.CompareStrings(string1, string2)).To(Equal(false))
      })

      It("compares two equal strings", func() {
        string1 := "x"
        string2 := "x"
        Expect(computer.CompareStrings(string1, string2)).To(Equal(true))
      })

      It("compares two empty strings", func() {
        string1 := ""
        string2 := ""
        Expect(computer.CompareStrings(string1, string2)).To(Equal(true))
      })

    })



//    Context("Minimax Implementation", func() {
//
//      Context("takes win if available", func() {
//
//        XIt("takes index 7 when it's a winner", func() {
//          board := []string{"x","o","x","","o","","","","x"}
//          Expect(player.MakeMove(board)).To(Equal(7))
//        })
//
//        XIt("takes index 5 when it's a winner", func() {
//          board := []string{"x","","x","o","o","","","","x"}
//          Expect(player.MakeMove(board)).To(Equal(5))
//        })
//
//        XIt("takes index 3 when it's a winner", func() {
//          board := []string{"x","","x","","o","o","x","",""}
//          Expect(player.MakeMove(board)).To(Equal(3))
//        })
//
//        XIt("takes index 8 when it's a winner", func() {
//          board := []string{"o","x","x","x","o","","","",""}
//          Expect(player.MakeMove(board)).To(Equal(3))
//        })
//      })
//
//    })

  })

