package ttt

type UserInterface interface {
  DisplayBoard(Board)
  SelectBoardChoice([]Board) (Board)
  DisplayPlayerTypes([]Player)
  SelectPlayerChoice([]Player, string) (Player)
  PrintChoiceInvalid()
  AskForNewGame() bool
  DisplayWinner(string)
  SelectMove(Player, Board) (int)
}
