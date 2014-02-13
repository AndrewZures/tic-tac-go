package ttt

import (
	"strconv"
	"strings"
)

type ConsoleUI struct {
	ConsoleIO      InOut
	Messages       Messages
	BoardFormatter BoardFormatter
}

func (c ConsoleUI) DisplayBoard(board Board) {
	formattedBoard := c.BoardFormatter.FormatBoard(board, c.Messages)
	c.ConsoleIO.Println(formattedBoard)
}

func (c ConsoleUI) QueryMove(player Player, board Board) int {

		c.displayQueryMoveText(player)
    rawMove := c.GetIntegerFromUser()
    return c.shiftToZerosBasedIndex(rawMove)
}

func (c ConsoleUI) QueryPlayerChoice(playerList []Player, description string) Player {
	c.PrintPlayerTypeQuestion(description)
	c.DisplayPlayerTypes(playerList)
	playerChoice := c.PlayerChoice(playerList)
	return playerChoice
}

func (c ConsoleUI) DisplayPlayerTypes(playerList []Player) {

	for index, playerType := range playerList {
		c.ConsoleIO.Printf(c.Messages.PlayerTypesResponse(), (index + 1), playerType.Description())
	}

}

func (c ConsoleUI) PrintPlayerTypeQuestion(playerDescription string) {
	c.ConsoleIO.Printf(c.Messages.PlayerTypePrompt(), playerDescription)
}

func (c ConsoleUI) PrintBoardTypeQuestion() {
	c.ConsoleIO.Println(c.Messages.BoardTypePrompt())
}

func (c *ConsoleUI) PlayerChoice(playerList []Player) Player {

	for {
		userChoice := c.GetIntegerFromUser()

		if c.ChoiceValid(userChoice, len(playerList)) {
			userChoice = c.shiftToZerosBasedIndex(userChoice)
			return playerList[userChoice]
		} else {
			c.DisplayChoiceInvalid()
		}
	}
}

func (c ConsoleUI) DisplayBoardTypes(boardList []Board) {

	for index, boardType := range boardList {
		c.ConsoleIO.Printf(c.Messages.BoardTypesResponse(), (index + 1), boardType.Description())
	}

}

func (c ConsoleUI) QueryBoardChoice(boardList []Board) Board {
	c.PrintBoardTypeQuestion()
	c.DisplayBoardTypes(boardList)
	boardChoice := c.BoardChoice(boardList)
	return boardChoice
}

func (c ConsoleUI) BoardChoice(boardList []Board) Board {
	for {
		userChoice := c.GetIntegerFromUser()

		if c.ChoiceValid(userChoice, len(boardList)) {
			userChoice = c.shiftToZerosBasedIndex(userChoice)
			return boardList[userChoice]
		} else {
			c.DisplayChoiceInvalid()
		}
	}
}

func (c ConsoleUI) GetIntegerFromUser() int {
	var userInput string
	var copiedInput string

	for {
		userInput = c.ConsoleIO.Read()
		copiedInput = strings.Repeat(userInput, 1)
		value, err := strconv.ParseInt(copiedInput, 0, 0)

		if err == nil {
			return int(value)
		} else {
			c.DisplayChoiceInvalid()
		}
	}

}


func (c ConsoleUI) QueryNewGame() bool {
	c.displayNewGameQuery()
	return c.getNewGameDecision()
}

func (c ConsoleUI) displayNewGameQuery() {
	c.ConsoleIO.Println(c.Messages.NewGamePrompt())
}

func (c ConsoleUI) getNewGameDecision() bool {
	response := strings.ToLower(c.ConsoleIO.Read())
	return response == c.Messages.YesOption()
}

func (c *ConsoleUI) shiftToZerosBasedIndex(onesBasedIndexChoice int) int {
	return onesBasedIndexChoice - 1
}

func (c *ConsoleUI) ChoiceValid(choice int, numChoices int) bool {
	return choice > 0 && choice <= numChoices
}

func (c ConsoleUI) DisplayWinner(winner string) {
	if winner == "tie" {
		c.ConsoleIO.Printf(c.Messages.GameTieResponse())
	} else {
		c.ConsoleIO.Printf(c.Messages.GameWinnerResponse(), c.Messages.WinnerSymbol(winner))
	}
}

func (c ConsoleUI) DisplayChoiceInvalid() {
	c.ConsoleIO.Println(c.Messages.InvalidChoiceResponse())
}

func (c ConsoleUI) DisplayExitMessage() {
	c.ConsoleIO.Println(c.Messages.ExitMessage())
}

func (c ConsoleUI) DisplayIntroMessage() {
	c.ConsoleIO.Println(c.Messages.IntroMessage())
}

func (c ConsoleUI) displayQueryMoveText(player Player) {
	c.ConsoleIO.Printf(c.Messages.ChooseMovePrompt(), player.Description())
}
