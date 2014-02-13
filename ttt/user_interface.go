package ttt

type UserInterface interface {
	QueryBoardChoice([]Board) Board
	QueryPlayerChoice([]Player, string) Player
	QueryMove(Player, Board) int
	QueryNewGame() bool

	DisplayPlayerTypes([]Player)
	DisplayBoard(Board)
	DisplayWinner(string)
	DisplayChoiceInvalid()
	DisplayIntroMessage()
	DisplayExitMessage()
}
