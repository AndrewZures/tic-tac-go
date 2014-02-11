package ttt_test

import (
  . "github.com/andrewzures/tictactoe/ttt"
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
)

var _ = Describe("Basic Board", func() {
  var board Board

  BeforeEach(func(){
    board = GenerateEmpty3x3Board("x")
  })

  Context("when keeping state", func() {
    var emptyBoardState []string

    BeforeEach(func(){
      emptyBoardState = make([]string, 9, 9)
    })

    It("initalizes with empty board", func() {
      board := Generate3x3Board("x", []string{"","","","","","","","",""})
      Expect(board.State()).To(Equal(emptyBoardState))
      Expect(len(board.State())).To(Equal(9))
    })

    It("keeps game state after attempted reset", func() {
      board := Generate3x3Board("x", []string{"","x","","","x","","","",""})
      board.NewBoard(9,3,"x", "another board")
      Expect(board.State()).NotTo((Equal(emptyBoardState)))
    })
  })

  Context("recording player moves", func() {

    Context("adds state to internal board array", func() {

      It("records a player's move", func() {
        expectedResult := []string{"o","","","","","","","",""}
        board.RecordMove(0,"o")
        Expect(board.State()).To(Equal(expectedResult))
      })

      It("records multiple moves", func() {
        expectedResult := []string{"x","","","x","","","x","","x"}
        board := Generate3x3Board("o", []string{"x","","","x","","","x","","x"})
        Expect(board.State()).To(Equal(expectedResult))
      })
    })

    Context("validates moves", func() {

      It("determines if move choice is within board bounds", func(){
        Expect(board.RecordMove(-1,"o")).To(Equal(false))
        Expect(board.RecordMove(0,"o")).To(Equal(true))

        Expect(board.RecordMove(9,"x")).To(Equal(false))
        Expect(board.RecordMove(8,"x")).To(Equal(true))
      })

      It("determines if move choice has already been selected", func(){
        expectedResult := []string{"","","","o","","","","",""}

        Expect(board.RecordMove(3,"o")).To(Equal(true))
        Expect(board.State()).To(Equal(expectedResult))

        Expect(board.RecordMove(3,"x")).To(Equal(false))
        Expect(board.State()).To(Equal(expectedResult))
      })
    })

    It("returns list of available moves after single move", func() {
      board := Generate3x3Board("", []string{"","x","","","","","","",""})
      expectedResult := []int{0,2,3,4,5,6,7,8}
      Expect(board.OpenSpots()).To(Equal(expectedResult))
    })

    It("returns list of available moves after multiple moves", func() {
      board := Generate3x3Board("", []string{"o","x","","o","","","","","x"})
      expectedResult := []int{2,4,5,6,7}
      Expect(board.OpenSpots()).To(Equal(expectedResult))
    })
  })

})

