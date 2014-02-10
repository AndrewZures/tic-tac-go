package ttt

type Rules interface {
  Winner(Board) string
  IsWinner(Board) bool
  IsPlayerTurn(Player) bool
  Score(Board) string
  RowWinner(Board) string
  ColumnWinner(Board) string
  DiagonalWinner(Board) string
}
