package ttt_test

import (
	"bytes"
	. "github.com/andrewzures/tictactoe/ttt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Console UI", func() {
	var console ConsoleUI
	var factory Factory
	var messages Messages
	var rules Rules
	var writer bytes.Buffer
	var reader bytes.Buffer

	BeforeEach(func() {
		boardFormatter := BoardFormatter(new(ConsoleBoardFormatter))
		factory = Factory(new(TTTFactory))
		rules = new(BasicRules)

		inOut := InOut(ConsoleIO{&writer, &reader})
		consoleMessages := new(ConsoleMessages)
		consoleMessages.BuildMessages()
		messages = Messages(consoleMessages)

		console = ConsoleUI{inOut, messages, boardFormatter}
	})

	Context("when displaying information to user", func() {

		It("displays game winner", func() {
			console.DisplayWinner("Player 1")
			Expect(writer.String()).To(ContainSubstring("Player 1"))
		})

		It("displays tie game", func() {
			console.DisplayWinner("tie")
			Expect(writer.String()).To(ContainSubstring("Tie"))
		})

		It("displays introduction message", func() {
			console.DisplayIntroMessage()
			Expect(writer.String()).To(ContainSubstring("Welcome to Tic Tac Go"))
		})

		It("displays exit message", func() {
			console.DisplayExitMessage()
			Expect(writer.String()).To(ContainSubstring("So Long!"))
		})

		It("queries player for move", func() {
			SetMockInput(&reader, "1\n")
			player := new(HumanPlayer)
			player.NewHumanPlayer("X", "Player 1", console)
			board := GenerateEmpty3x3Board()
			console.QueryMove(player, board)
			Expect(writer.String()).To(ContainSubstring("Choose a Move"))
		})

		It("indicates invalid choice to user", func() {
			console.DisplayChoiceInvalid()
			Expect(writer.String()).To(ContainSubstring("Whoops, that choice is invalid"))
			Expect(writer.String()).To(ContainSubstring("Try Again"))
		})

		It("indicates invalid choice to user if integer expected and not received", func() {
			SetMockInput(&reader, "firsstring\n1\n")
			console.GetIntegerFromUser()
			Expect(writer.String()).To(ContainSubstring("Whoops, that choice is invalid"))
			Expect(writer.String()).To(ContainSubstring("Try Again"))
		})

		It("queries user for board type", func() {
			console.PrintBoardTypeQuestion()
			Expect(writer.String()).To(ContainSubstring("Choose Board Type:"))
		})

		It("queries user for player type", func() {
			console.PrintPlayerTypeQuestion("Player 1")
			Expect(writer.String()).To(ContainSubstring("Choose Type for: Player 1"))
		})

		It("queries user for new game", func() {
			SetMockInput(&reader, "y\nn\n")
			result := console.QueryNewGame()
			Expect(writer.String()).To(ContainSubstring("Would you like to start a new game"))
			Expect(result).To(Equal(true))

			result = console.QueryNewGame()
			Expect(result).To(Equal(false))
		})

		It("lists available player types", func() {
			console.DisplayPlayerTypes(factory.PlayerTypes(console, rules))
			Expect(writer.String()).To(ContainSubstring("Human"))
			Expect(writer.String()).To(ContainSubstring("Computer"))
		})

		It("lists available board types", func() {
			console.DisplayBoardTypes(factory.BoardTypes())
			Expect(writer.String()).To(ContainSubstring("3x3 Board"))
			Expect(writer.String()).To(ContainSubstring("4x4 Board"))
			Expect(writer.String()).To(ContainSubstring("5x5 Board"))
		})

	})

	Context("when receiving input from user", func() {

		It("validates input from use is integer", func() {
			SetMockInput(&reader, "3")
			Expect(console.GetIntegerFromUser()).To(Equal(3))
		})

		It("requeries user until integer is provided", func() {
			SetMockInput(&reader, "firsstring\nsecond123string2\n1\n")
			Expect(console.GetIntegerFromUser()).To(Equal(1))
		})

		It("reads multiple integer inputs from user", func() {
			SetMockInput(&reader, "firsstring\nsecond123string2\n3\n6\n")
			Expect(console.GetIntegerFromUser()).To(Equal(3))
			Expect(console.GetIntegerFromUser()).To(Equal(6))
		})

	})

	Context("when communicating with program", func() {
		var playerTypes []Player
		var boardTypes []Board

		BeforeEach(func() {
			playerTypes = factory.PlayerTypes(console, rules)
			boardTypes = factory.BoardTypes()
		})

		It("returns selected player template", func() {
			SetMockInput(&reader, "1")
			selectedPlayer := console.QueryPlayerChoice(playerTypes, "player 1")
			Expect(selectedPlayer.Description()).To(Equal("Human"))

			SetMockInput(&reader, "2")
			selectedPlayer = console.QueryPlayerChoice(playerTypes, "player 1")
			Expect(selectedPlayer.Description()).To(Equal("Computer"))
		})

		It("only returns valid player template", func() {
			SetMockInput(&reader, "-1\n100\na\n2\n1\n")
			selectedPlayer := console.QueryPlayerChoice(playerTypes, "player 1")
			Expect(selectedPlayer.Description()).To(Equal("Computer"))
		})

		It("returns multiptle players templates", func() {
			SetMockInput(&reader, "a\n1\n2\n2\n")
			selectedPlayer := console.QueryPlayerChoice(playerTypes, "player 1")
			Expect(selectedPlayer.Description()).To(Equal("Human"))

			anotherPlayer := console.QueryPlayerChoice(playerTypes, "player 2")
			Expect(anotherPlayer.Description()).To(Equal("Computer"))
		})

		It("returns board type template", func() {
			SetMockInput(&reader, "1\n")
			console.QueryBoardChoice(boardTypes)
			Expect(writer.String()).To(ContainSubstring("3x3 Board"))
		})

		It("re-queries user if choice not within range", func() {
			SetMockInput(&reader, "5\n1\n")
			console.QueryBoardChoice(boardTypes)
			Expect(writer.String()).To(ContainSubstring("Whoops, that choice is invalid"))
		})

	})

	AfterEach(func() {
		writer.Reset()
		reader.Reset()
	})

})
