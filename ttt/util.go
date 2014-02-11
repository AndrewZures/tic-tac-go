package ttt

func GetGameStateAndOffset (board Board) ([]string, int) {
  gameState := board.State()
  offset := board.Offset()
  return gameState, offset
}
