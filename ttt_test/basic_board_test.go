package ttt_test

import (
  . "github.com/andrewzures/tictactoe/ttt"
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"

)

var _ = Describe("Basic Board", func() {
  var basicBoard *BasicBoard
  var player Player
  var emptyBoardState []string

  BeforeEach(func(){
    human := new(HumanPlayer)
    human.SetSymbol("X")
    player = Player(human)
    basicBoard = new(BasicBoard)
    basicBoard.NewBoard()
    emptyBoardState = make([]string, 9, 9)
  })

  Context("recording player moves", func() {

    It("records a player's move", func() {
      expectedResult := []string{"X","","","","","","","",""}
      basicBoard.RecordMove(0,player)
      Expect(basicBoard.Array()).To(Equal(expectedResult))
    })

    It("records multiple moves", func() {
      expectedResult := []string{"X","","","X","","","X","","X"}
      basicBoard.RecordMove(0,player)
      basicBoard.RecordMove(3,player)
      basicBoard.RecordMove(6,player)
      basicBoard.RecordMove(8,player)
      Expect(basicBoard.Array()).To(Equal(expectedResult))
    })

    It("validates player move", func(){
      Expect(basicBoard.RecordMove(3,player)).To(Equal(true))
      Expect(basicBoard.RecordMove(3,player)).To(Equal(false))
      Expect(basicBoard.RecordMove(3,player)).To(Equal(false))
    })
  })

  Context("keeping and accessing board status", func() {

    It("initalizes with empty board", func() {
      Expect(basicBoard.Array()).To(Equal(emptyBoardState))
      Expect(len(basicBoard.Array())).To(Equal(9))
    })

    It("keeps game state after attempted reset", func() {
      basicBoard.RecordMove(1,player)
      basicBoard.RecordMove(4,player)
      basicBoard.NewBoard()
      Expect(basicBoard.Array()).NotTo((Equal(emptyBoardState)))
    })
  })




})

