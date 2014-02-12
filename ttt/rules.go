package ttt

type Rules interface {
  Winner(Board) string
  IsWinner(Board) bool
  Score(Board) string
  GameOver(Board) bool
}
