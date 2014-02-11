package ttt_test

import (
  . "github.com/andrewzures/tictactoe/ttt"
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
)

var _ = Describe("Basic Rules", func() {
  var rules Rules

    BeforeEach(func(){
      rules = Rules(new(BasicRules))
    })

      It("keeps track of player turn", func() {
        Expect(1).To(Equal(1))
        })

      It("determines if winner exists", func() {
        board := Generate3x3Board("", []string{"x","x","x","","","","","",""})
        Expect(rules.IsWinner(board)).To(Equal(true))

        board = Generate3x3Board("", []string{"x","o","x","","","","","",""})
        Expect(rules.IsWinner(board)).To(Equal(false))
      })

      It("returns winner", func() {

        board := Generate3x3Board("", []string{"x","x","x","","","","","",""})
        Expect(rules.Winner(board)).To(Equal("x"))

        board = Generate3x3Board("", []string{"x","o","x","","","","","",""})
        Expect(rules.Winner(board)).To(Equal("inprogress"))
      })

    Context("when scoring a board", func() {

      It("returns winner if there is a row winner", func() {
        board := Generate3x3Board("", []string{"x","x","x","","","","","",""})
        Expect(rules.Score(board)).To(Equal("x"))
      })

      It("returns winner if there is a column winner", func() {
        board := Generate3x3Board("", []string{"o","","","o","","","o","",""})
        Expect(rules.Score(board)).To(Equal("o"))
      })

      It("returns winner if there is a diagonal winner", func() {
        board := Generate3x3Board("", []string{"","","x","","x","","x","",""})
        Expect(rules.Score(board)).To(Equal("x"))
      })

      It("returns tie if no winner and no more avialable moves", func() {
        board := Generate3x3Board("", []string{"o","x","o","o","x","o","x","o","x"})
        Expect(rules.Score(board)).To(Equal("tie"))
      })

      It("returns inprogress if no winner but more moves available", func() {
        board := Generate3x3Board("", []string{"o","","","","x","","","",""})
        Expect(rules.Score(board)).To(Equal("inprogress"))

        board = Generate3x3Board("x", []string{"o","x","x","","x","o","o","",""})
        Expect(rules.Score(board)).To(Equal("inprogress"))
      })

    })


    Context("when scoring a row", func() {

      It("finds a row winner", func() {
        board := Generate3x3Board("", []string{"x","x","x","","","","","",""})
        Expect(rules.Score(board)).To(Equal("x"))

        board = Generate3x3Board("",[]string{"","","","o","o","o","","",""})
        Expect(rules.Score(board)).To(Equal("o"))

        board = Generate3x3Board("", []string{"","","","","","","x","x","x"})
        Expect(rules.Score(board)).To(Equal("x"))
      })

//      It("does not generate false positives", func() {
//
//        board := GenerateEmpty3x3Board("")
//        Expect(rules.Score(board)).To(Equal(""))
//
//        board = Generate3x3Board("", []string{"x","","","x","","","x","",""})
//        Expect(rules.Score(board)).To(Equal(""))
//
//        board = Generate3x3Board("", []string{"x","","","","x","","","","x"})
//        Expect(rules.Score(board)).To(Equal(""))
//      })

    })

    Context("when scoring a column", func() {

      It("finds a column winner", func() {
        board := Generate3x3Board("", []string{"x","","","x","","","x","",""})
        Expect(rules.Score(board)).To(Equal("x"))

        board = Generate3x3Board("", []string{"","o","","","o","","","o",""})
        Expect(rules.Score(board)).To(Equal("o"))

        board = Generate3x3Board("", []string{"","","x","","","x","","","x"})
        Expect(rules.Score(board)).To(Equal("x"))
      })

//      It("does not generate false positives", func() {
//
//        board := Generate3x3Board("", []string{"","","","","","","","",""})
//        Expect(rules.Score(board)).To(Equal(""))
//
//        board = Generate3x3Board("", []string{"","","","o","x","x","","",""})
//        Expect(rules.Score(board)).To(Equal(""))
//
//        board = Generate3x3Board("", []string{"","x","","","o","","","","x"})
//        Expect(rules.Score(board)).To(Equal(""))
//      })

    })

    Context("when scoring a diagonal", func() {

      It("finds a diagonal winner", func() {
        board := Generate3x3Board("", []string{"x","","","","x","","","","x"})
        Expect(rules.Score(board)).To(Equal("x"))

        board = Generate3x3Board("", []string{"","","o","","o","","o","",""})
        Expect(rules.Score(board)).To(Equal("o"))
      })

//      It("does not generate false positives", func() {
//
//        board := GenerateEmpty3x3Board("")
//        Expect(rules.Score(board)).To(Equal(""))
//
//        board = Generate3x3Board("", []string{"x","","","","","","","x","x"})
//        Expect(rules.Score(board)).To(Equal(""))
//
//        board = Generate3x3Board("", []string{"","","x","","","o","","","o"})
//        Expect(rules.Score(board)).To(Equal(""))
//      })

    })

//      It("keeps track of player turn", func() {
//        board := Generate3x3Board("o", []string{"","","","","","","","",""})
//        Expect(board.PlayerTurn()).To(Equal("o"))
//
//        board.RecordMove(0, "x")
//        Expect(board.PlayerTurn()).To(Equal("o"))
//
//        board.RecordMove(1, "o")
//        Expect(board.PlayerTurn()).To(Equal("x"))
//      })
//
//      It("records a player's move also", func() {
//        board := Board(GenerateEmpty3x3Board("o"))
//        moveResult := board.RecordMove(0, "o")
//        Expect(moveResult).To(Equal(true))
//
//        Expect(board.PlayerTurn()).To(Equal("x"))
//        Expect(board.RecordMove(1,"o")).To(Equal(false))
//        Expect(board.RecordMove(1,"o")).To(Equal(false))
//        Expect(board.RecordMove(1,"x")).To(Equal(true))
//      })
})
