package ttt

type Rules interface {
  Winner(Board) string
  IsPlayerTurn(Player) bool
  Score(Board) string
  RowWinner(Board) string
  ColumnWinner(Board) string
  DiagonalWinner(Board) string
}
