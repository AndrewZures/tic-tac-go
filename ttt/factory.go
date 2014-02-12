package ttt

type Factory interface {
	PlayerTypes(UserInterface, Rules) []Player
	Player(Player, UserInterface, Rules) Player
	BoardTypes() []Board
	Board(Board) Board
}
