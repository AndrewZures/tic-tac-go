package ttt

func GetGameStateAndOffset (board Board) ([]string, int) {
  gameState := board.Array()
  offset := board.Offset()
  return gameState, offset
}

