package ttt_test

import (
  . "github.com/andrewzures/tictactoe/ttt"
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
)

var _ = Describe("Basic Board", func() {
  var playerX, playerO Player

  BeforeEach(func(){
    playerX, playerO = getPlayers()
  })

  It("can implement the Board interface", func() {
    board := Board(new(BasicBoard))
    board.NewBoard()
    Expect(board.RecordMove(0,playerX.Symbol())).To(Equal(true))
  })

  Context("recording player moves", func() {
    var basicBoard *BasicBoard

    BeforeEach(func(){
      basicBoard = new(BasicBoard)
      basicBoard.NewBoard()
    })

    Context("validates moves", func() {

      It("determines if move choice is within board bounds", func(){
        Expect(basicBoard.RecordMove(-1,playerX.Symbol())).To(Equal(false))
        Expect(basicBoard.RecordMove(0,playerX.Symbol())).To(Equal(true))

        Expect(basicBoard.RecordMove(9,playerX.Symbol())).To(Equal(false))
        Expect(basicBoard.RecordMove(8,playerX.Symbol())).To(Equal(true))
      })

      It("determines if move choice has already been selected", func(){
        expectedResult := []string{"","","","X","","","","",""}

        Expect(basicBoard.RecordMove(3,playerX.Symbol())).To(Equal(true))
        Expect(basicBoard.Array()).To(Equal(expectedResult))

        Expect(basicBoard.RecordMove(3,playerO.Symbol())).To(Equal(false))
        Expect(basicBoard.Array()).To(Equal(expectedResult))
      })
    })

    Context("adds state to internal board array", func() {

      It("records a player's move", func() {
        expectedResult := []string{"X","","","","","","","",""}
        basicBoard.RecordMove(0,playerX.Symbol())
        Expect(basicBoard.Array()).To(Equal(expectedResult))
      })

      It("records multiple moves", func() {
        expectedResult := []string{"X","","","X","","","X","","X"}
        board := GenerateBoard([]string{"x","","","x","","","x","","x"})
        Expect(board.Array()).To(Equal(expectedResult))
      })
    })
  })

  Context("holding and accessing board status", func() {
    var emptyBoardState []string

    BeforeEach(func(){
      emptyBoardState = make([]string, 9, 9)
    })

    It("initalizes with empty board", func() {
      board := GenerateBoard([]string{"","","","","","","","",""})
      Expect(board.Array()).To(Equal(emptyBoardState))
      Expect(len(board.Array())).To(Equal(9))
    })

    It("keeps game state after attempted reset", func() {
      board := GenerateBoard([]string{"","x","","","x","","","",""})
      board.NewBoard()
      Expect(board.Array()).NotTo((Equal(emptyBoardState)))
    })

    It("returns list of available moves after single move", func() {
      board := GenerateBoard([]string{"","x","","","","","","",""})
      expectedResult := []int{0,2,3,4,5,6,7,8}
      Expect(board.OpenSpots()).To(Equal(expectedResult))
    })

    It("returns list of available moves after multiple moves", func() {
      board := GenerateBoard([]string{"o","x","","o","","","","","x"})
      expectedResult := []int{2,4,5,6,7}
      Expect(board.OpenSpots()).To(Equal(expectedResult))
    })
  })

  Context("when scoring a game", func() {

    Context("when getting game status", func() {

      It("returns winner if there is a row winner", func() {
        boardContents := []string{"x","x","x","","","","","",""}
        board := GenerateBoard(boardContents)
        Expect(board.Status()).To(Equal("X"))
      })

      It("returns winner if there is a column winner", func() {
        boardContents := []string{"o","","","o","","","o","",""}
        board := GenerateBoard(boardContents)
        Expect(board.Status()).To(Equal("O"))
      })

      It("returns winner if there is a diagonal winner", func() {
        boardContents := []string{"","","x","","x","","x","",""}
        board := GenerateBoard(boardContents)
        Expect(board.Status()).To(Equal("X"))
      })

      It("returns tie if no winner and no more avialable moves", func() {
        boardContents := []string{"o","x","o","o","x","o","x","o","x"}
        board := GenerateBoard(boardContents)
        Expect(board.Status()).To(Equal("tie"))
      })
    })

    Context("when scoring a row", func() {

      It("finds a row winner", func() {
        boardContents := []string{"x","x","x","","","","","",""}
        board := GenerateBoard(boardContents)
        Expect(board.RowWinner()).To(Equal("X"))

        boardContents = []string{"","","","o","o","o","","",""}
        board = GenerateBoard(boardContents)
        Expect(board.RowWinner()).To(Equal("O"))

        boardContents = []string{"","","","","","","x","x","x"}
        board = GenerateBoard(boardContents)
        Expect(board.RowWinner()).To(Equal("X"))
      })

      It("does not generate false positives", func() {

        boardContents := []string{"","","","","","","","",""}
        board := GenerateBoard(boardContents)
        Expect(board.RowWinner()).To(Equal(""))

        boardContents = []string{"x","","","x","","","x","",""}
        board = GenerateBoard(boardContents)
        Expect(board.RowWinner()).To(Equal(""))

        boardContents = []string{"x","","","","x","","","","x"}
        board = GenerateBoard(boardContents)
        Expect(board.RowWinner()).To(Equal(""))
      })

    })

    Context("when scoring a column", func() {

      It("finds a column winner", func() {
        boardContents := []string{"x","","","x","","","x","",""}
        board := GenerateBoard(boardContents)
        Expect(board.ColumnWinner()).To(Equal("X"))

        boardContents = []string{"","o","","","o","","","o",""}
        board = GenerateBoard(boardContents)
        Expect(board.ColumnWinner()).To(Equal("O"))

        boardContents = []string{"","","x","","","x","","","x"}
        board = GenerateBoard(boardContents)
        Expect(board.ColumnWinner()).To(Equal("X"))
      })

      It("does not generate false positives", func() {

        boardContents := []string{"","","","","","","","",""}
        board := GenerateBoard(boardContents)
        Expect(board.ColumnWinner()).To(Equal(""))

        boardContents = []string{"","","","x","x","x","","",""}
        board = GenerateBoard(boardContents)
        Expect(board.ColumnWinner()).To(Equal(""))

        boardContents = []string{"x","","","","x","","","","x"}
        board = GenerateBoard(boardContents)
        Expect(board.ColumnWinner()).To(Equal(""))
      })

    })

    Context("when scoring a diagonal", func() {

      It("finds a diagonal winner", func() {
        boardContents := []string{"x","","","","x","","","","x"}
        board := GenerateBoard(boardContents)
        Expect(board.DiagonalWinner()).To(Equal("X"))

        boardContents = []string{"","","o","","o","","o","",""}
        board = GenerateBoard(boardContents)
        Expect(board.DiagonalWinner()).To(Equal("O"))
      })

      It("does not generate false positives", func() {

        boardContents := []string{"","","","","","","","",""}
        board := GenerateBoard(boardContents)
        Expect(board.DiagonalWinner()).To(Equal(""))

        boardContents = []string{"","","","","","","x","x","x"}
        board = GenerateBoard(boardContents)
        Expect(board.DiagonalWinner()).To(Equal(""))

        boardContents = []string{"","","o","","","o","","","o"}
        board = GenerateBoard(boardContents)
        Expect(board.DiagonalWinner()).To(Equal(""))
      })

    })

  })

})

func getPlayers() (Player, Player) {
  human1 := new(HumanPlayer)
  human1.SetSymbol("X")

  human2 := new(HumanPlayer)
  human2.SetSymbol("O")

  return Player(human1), Player(human2)
}
