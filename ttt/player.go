package ttt

type Player interface {
	MakeMove(Board) int
	Symbol() string
	SetSymbol(string)
	Description() string
}
