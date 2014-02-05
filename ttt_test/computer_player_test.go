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
      computer.SetSymbol("O")
      board = Board(new(BasicBoard))
      board.NewBoard()
    })

    Context("Basic Player Attributes", func() {

      It("has a symbol", func() {
        Expect(computer.Symbol()).To(Equal("O"))
      })

      It("meets player interface requirements", func() {
        player = Player(computer)
        Expect(player.Symbol()).To(Equal("O"))
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

        It("takes index 7 when it's a winner", func() {
          boardContents := []string{"x","o","x","","o","","","","x"}
          board := GenerateBoard(boardContents)
          Expect(computer.MakeMove(board)).To(Equal(7))
        })

        It("takes index 5 when it's a winner", func() {
          boardContents := []string{"x","","x","o","o","","","","x"}
          board := GenerateBoard(boardContents)
          Expect(computer.MakeMove(board)).To(Equal(5))
        })

        It("takes index 3 when it's a winner", func() {
          boardContents := []string{"x","","x","","o","o","x","",""}
          board := GenerateBoard(boardContents)
          Expect(computer.MakeMove(board)).To(Equal(3))
        })

        It("takes index 8 when it's a winner", func() {
          boardContents := []string{"o","x","x","x","o","","","",""}
          board := GenerateBoard(boardContents)
          Expect(player.MakeMove(board)).To(Equal(8))
        })
      })

      Context("defends against immediate loss", func() {

        It("takes index 3 when it blocks opponent win", func() {
          boardContents := []string{"x","o","x","x","x","o","o","",""}
          board := GenerateBoard(boardContents)
          Expect(computer.MakeMove(board)).To(Equal(8))
        })

        It("takes index 3 when it blocks opponent win", func() {
          boardContents := []string{"x","","","","","o","x","",""}
          board := GenerateBoard(boardContents)
          Expect(computer.MakeMove(board)).To(Equal(3))
        })

        It("takes index 1 when it blocks opponent win", func() {
          boardContents := []string{"x","","x","o","","","","",""}
          board := GenerateBoard(boardContents)
          Expect(computer.MakeMove(board)).To(Equal(1))
        })

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
