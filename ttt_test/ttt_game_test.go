package ttt_test

import (
	"bytes"
	. "github.com/andrewzures/tictactoe/ttt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Game Test", func() {
	var userInterface UserInterface
	var messages Messages
	var writer bytes.Buffer
	var reader bytes.Buffer
	var factory Factory
	var console UserInterface
	var game Game

	BeforeEach(func() {
		boardFormatter := BoardFormatter(new(ConsoleBoardFormatter))
		inOut := InOut(ConsoleIO{&writer, &reader})
		userInterface = UserInterface(console)

		consoleMessage := new(ConsoleMessages)
		messages = Messages(consoleMessage)
		messages.BuildMessages()
		consoleui := ConsoleUI{inOut, messages, boardFormatter}
		console = UserInterface(consoleui)

		factory = Factory(new(TTTFactory))

		game = new(TTTGame)
	})

	It("sets up a game", func() {
		SetMockInput(&reader, "1\n2\n1\n")
		emptyBoardArray := []string{"", "", "", "", "", "", "", "", ""}

		board, player1, player2, _ := game.SetupNewGame(console, factory)

		Expect(player1.Description()).To(Equal("Human"))
		Expect(player2.Description()).To(Equal("Computer"))
		Expect(board.Description()).To(Equal("3x3 Board"))
		Expect(board.State()).To(Equal(emptyBoardArray))
	})

	It("has choose move response", func() {
		completeHumanGame := "1\n1\n1\n1\n2\n3\n4\n5\n6\n7\nn\n"
		SetMockInput(&reader, completeHumanGame)
		game.Run(console, factory)
		Expect(writer.String()).To(ContainSubstring("Winner is X"))
	})

	It("has choose move response", func() {
		completeComputerGame := "2\n2\n1\nn\n"
		SetMockInput(&reader, completeComputerGame)
		game.Run(console, factory)
		Expect(writer.String()).To(ContainSubstring(messages.GameTieResponse()))
	})

	It("displays empty board for human/human game", func() {
		completeHumanGame := "1\n1\n1\n1\n2\n3\n4\n5\n6\n7\nn\n"
		SetMockInput(&reader, completeHumanGame)
		game.Run(console, factory)
		Expect(writer.String()).To(ContainSubstring("  |   |  \n---------\n  |   |  \n---------\n  |   |  \n"))
	})

	AfterEach(func() {
		writer.Reset()
		reader.Reset()
	})

})
