package ttt

func GetGameStateAndOffset (board Board) ([]string, int) {
  gameState := board.BoardState()
  offset := board.Offset()
  return gameState, offset
}
