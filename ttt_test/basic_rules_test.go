package ttt_test

import (
	. "github.com/andrewzures/tictactoe/ttt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Basic Rules", func() {
	var rules Rules

	BeforeEach(func() {
		rules = Rules(new(BasicRules))
	})

	It("determines if winner exists", func() {
		board := Generate3x3Board([]string{"x", "x", "x", "", "", "", "", "", ""})
		Expect(rules.IsWinner(board)).To(Equal(true))

		board = Generate3x3Board([]string{"x", "o", "x", "", "", "", "", "", ""})
		Expect(rules.IsWinner(board)).To(Equal(false))
	})

	It("returns winner", func() {

		board := Generate3x3Board([]string{"x", "x", "x", "", "", "", "", "", ""})
		Expect(rules.Winner(board)).To(Equal("x"))

		board = Generate3x3Board([]string{"x", "o", "x", "", "", "", "", "", ""})
		Expect(rules.Winner(board)).To(Equal("inprogress"))
	})

	Context("when scoring a board", func() {

		It("returns winner if there is a row winner", func() {
			board := Generate3x3Board([]string{"x", "x", "x", "", "", "", "", "", ""})
			Expect(rules.Score(board)).To(Equal("x"))
		})

		It("returns winner if there is a column winner", func() {
			board := Generate3x3Board([]string{"o", "", "", "o", "", "", "o", "", ""})
			Expect(rules.Score(board)).To(Equal("o"))
		})

		It("returns winner if there is a diagonal winner", func() {
			board := Generate3x3Board([]string{"", "", "x", "", "x", "", "x", "", ""})
			Expect(rules.Score(board)).To(Equal("x"))
		})

		It("returns tie if no winner and no more avialable moves", func() {
			board := Generate3x3Board([]string{"o", "x", "o", "o", "x", "o", "x", "o", "x"})
			Expect(rules.Score(board)).To(Equal("tie"))
		})

		It("returns inprogress if no winner but more moves available", func() {
			board := Generate3x3Board([]string{"o", "", "", "", "x", "", "", "", ""})
			Expect(rules.Score(board)).To(Equal("inprogress"))

			board = Generate3x3Board([]string{"o", "x", "x", "", "x", "o", "o", "", ""})
			Expect(rules.Score(board)).To(Equal("inprogress"))
		})

		It("finds a row winner", func() {
			board := Generate3x3Board([]string{"x", "x", "x", "", "", "", "", "", ""})
			Expect(rules.Score(board)).To(Equal("x"))

			board = Generate3x3Board([]string{"", "", "", "o", "o", "o", "", "", ""})
			Expect(rules.Score(board)).To(Equal("o"))

			board = Generate3x3Board([]string{"", "", "", "", "", "", "x", "x", "x"})
			Expect(rules.Score(board)).To(Equal("x"))
		})

		It("finds a column winner", func() {
			board := Generate3x3Board([]string{"x", "", "", "x", "", "", "x", "", ""})
			Expect(rules.Score(board)).To(Equal("x"))

			board = Generate3x3Board([]string{"", "o", "", "", "o", "", "", "o", ""})
			Expect(rules.Score(board)).To(Equal("o"))

			board = Generate3x3Board([]string{"", "", "x", "", "", "x", "", "", "x"})
			Expect(rules.Score(board)).To(Equal("x"))
		})

		It("finds a diagonal winner", func() {
			board := Generate3x3Board([]string{"x", "", "", "", "x", "", "", "", "x"})
			Expect(rules.Score(board)).To(Equal("x"))

			board = Generate3x3Board([]string{"", "", "o", "", "o", "", "o", "", ""})
			Expect(rules.Score(board)).To(Equal("o"))
		})
	})

})
