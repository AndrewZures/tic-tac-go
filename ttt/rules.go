package ttt

type Rules interface {
  Winner(Board) string
  IsWinner(Board) bool
  IsPlayerTurn(Player) bool
  Score(Board) string
  GameOver(Board) bool
}
