package ttt_test

import (
	"bytes"
	. "github.com/andrewzures/tictactoe/ttt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Human Player", func() {
	var player Player
	var board Board
	var writer bytes.Buffer
	var reader bytes.Buffer

	BeforeEach(func() {
		inOut := InOut(ConsoleIO{&writer, &reader})
		console := BuildConsoleUI(inOut)
		userInterface := UserInterface(console)

		human := new(HumanPlayer)
		human.NewHumanPlayer("x", "human", userInterface)
		player = Player(human)
		board = Generate3x3Board(make([]string, 9, 9))
	})

	It("makes a move", func() {
		SetMockInput(&reader, "1\n8")
		Expect(player.MakeMove(board)).To(Equal(0))

		Expect(player.MakeMove(board)).To(Equal(7))
	})

	It("has a symbol", func() {
		Expect(player.Symbol()).To(Equal("x"))
	})

})
