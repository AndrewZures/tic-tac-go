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
      computer.SetSymbol("X")
      board = Board(new(BasicBoard))
      board.NewBoard()
    })

    Context("Basic Player Attributes", func() {

      It("has a symbol", func() {
        Expect(computer.Symbol()).To(Equal("X"))
      })

      It("meets player interface requirements", func() {
        player = Player(computer)
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

    Context("Minimax Implementation", func() {

      Context("takes win if available", func() {

//        It("takes index 7 when it's a winner", func() {
//          boardContents := []string{"x","o","x","","o","","","","x"}
//          board := generateBoard(boardContents)
//
//          //TODO remove this later
//          Expect(board.RecordMove(0, computer.Symbol())).To(Equal(false))
//
//          Expect(computer.MakeMove(board)).To(Equal(7))
//        })
//
//        XIt("takes index 5 when it's a winner", func() {
//          boardContents := []string{"x","","x","o","o","","","","x"}
//          board := generateBoard(boardContents)
//          Expect(computer.MakeMove(board)).To(Equal(5))
//        })
//
//        XIt("takes index 3 when it's a winner", func() {
//          boardContents := []string{"x","","x","","o","o","x","",""}
//          board := generateBoard(boardContents)
//          Expect(computer.MakeMove(board)).To(Equal(3))
//        })
//
//        XIt("takes index 8 when it's a winner", func() {
//          boardContents := []string{"o","x","x","x","o","","","",""}
//          board := generateBoard(boardContents)
//          Expect(player.MakeMove(board)).To(Equal(3))
//        })
      })

    })

  })

func GenerateBoard(gameState []string) (Board) {
  newBoard := Board(new(BasicBoard))
  newBoard.NewBoard()
  playerX, playerO := getPlayers()

  for i := 0; i < len(gameState); i++ {
    if gameState[i] == "x" || gameState[i] == "X" {
      newBoard.RecordMove(i, playerX.Symbol())
    }

    if gameState[i] == "o" || gameState[i] == "O" {
      newBoard.RecordMove(i, playerO.Symbol())
    }
  }

  return newBoard
}
