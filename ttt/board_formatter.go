package ttt

type BoardFormatter interface {
	FormatBoard(Board, Messages) string
}
