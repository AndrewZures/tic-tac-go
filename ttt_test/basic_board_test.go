package ttt_test

import (
  . "github.com/andrewzures/tictactoe/ttt"
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
)

var _ = Describe("Basic Board", func() {
  var basicBoard *BasicBoard
  var board Board
  var playerX, playerO Player

  BeforeEach(func(){
    playerX, playerO = getPlayers()
    basicBoard = new(BasicBoard)
    basicBoard.NewBoard()
  })

  It("can implement the Board interface", func() {
    board = Board(basicBoard)
    Expect(board.RecordMove(0,playerX.Symbol())).To(Equal(true))
  })

  Context("recording player moves", func() {

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
        basicBoard.RecordMove(0,playerX.Symbol())
        basicBoard.RecordMove(3,playerX.Symbol())
        basicBoard.RecordMove(6,playerX.Symbol())
        basicBoard.RecordMove(8,playerX.Symbol())
        Expect(basicBoard.Array()).To(Equal(expectedResult))
      })
    })
  })

  Context("holding and accessing board status", func() {
    var emptyBoardState []string

    BeforeEach(func(){
      emptyBoardState = make([]string, 9, 9)
    })

    It("initalizes with empty board", func() {
      Expect(basicBoard.Array()).To(Equal(emptyBoardState))
      Expect(len(basicBoard.Array())).To(Equal(9))
    })

    It("keeps game state after attempted reset", func() {
      basicBoard.RecordMove(1,playerX.Symbol())
      basicBoard.RecordMove(4,playerX.Symbol())
      basicBoard.NewBoard()
      Expect(basicBoard.Array()).NotTo((Equal(emptyBoardState)))
    })

    It("return list of available moves single move", func() {
        basicBoard.RecordMove(1, playerX.Symbol())
        expectedResult := []int{0,2,3,4,5,6,7,8}
        Expect(basicBoard.OpenSpots()).To(Equal(expectedResult))
    })

    It("return list of available moves after multiple moves", func() {
        basicBoard.RecordMove(1, playerX.Symbol())
        basicBoard.RecordMove(3, playerO.Symbol())
        basicBoard.RecordMove(8, playerX.Symbol())
        basicBoard.RecordMove(0, playerO.Symbol())
        expectedResult := []int{2,4,5,6,7}
        Expect(basicBoard.OpenSpots()).To(Equal(expectedResult))
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
