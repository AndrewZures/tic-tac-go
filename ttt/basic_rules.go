package ttt

type BasicRules struct {
    playerTurn string;
}

func (r *BasicRules) NewBasicRules(startSymbol string) {
    r.playerTurn = startSymbol
}

func (r BasicRules) IsWinner(board Board) (bool) {
  if r.Score(board) != "inprogress" {
    return true
  }
  return false
}

func (r BasicRules) Winner(board Board) (string) {
    return r.Score(board)
}

func (r BasicRules) GameOver(board Board) (bool){
  return board.Status() != "inprogress"
}

func (r BasicRules) toggleTurn() {
  if r.playerTurn == "x" {
    r.playerTurn = "o"
  } else if r.playerTurn == "o" {
    r.playerTurn = "x"
  } else {
    r.playerTurn = "z"
  }
}

func (r BasicRules) PlayerTurn() (string) {
  return r.playerTurn
}

func (r BasicRules) IsPlayerTurn (player Player) (bool) {
  return r.PlayerTurn() == player.Symbol()
}

func (r BasicRules) Score(board Board) (string) {
  gameState := board.Array()
  openSpots := board.OpenSpots(gameState)

  rowStatus := r.RowWinner(board)
  if  rowStatus != "" {
    return rowStatus
  }

  columnStatus := r.ColumnWinner(board)
  if  columnStatus != "" {
    return columnStatus
  }

  diagonalStatus := r.DiagonalWinner(board)
  if  diagonalStatus != "" {
    return diagonalStatus
  }

  if len(openSpots) == 0 {
    return "tie"
  } else {
    return "inprogress"
  }

}

func (r BasicRules) RowWinner(board Board) (string) {
  var row []string
  gameState, offset := GetGameStateAndOffset(board)

  for i := 0; i < len(gameState); i += offset {
    row = r.rowElements(board, i)
    if r.AllSameSymbols(row) == true {
      return gameState[i]
    }
  }

  return ""
}

func (r BasicRules) rowElements(board Board, startIndex int) ([]string) {
  gameState := board.Array()
  return gameState[startIndex:startIndex+board.Offset()]
}

func (r BasicRules) ColumnWinner(board Board) (string) {
  var column []string
  gameState, offset := GetGameStateAndOffset(board)

  for i := 0; i < offset; i++ {
    column = r.columnElements(board, i)
    if r.AllSameSymbols(column) == true {
      return gameState[i]
    }
  }

  return ""
}

func (r *BasicRules) columnElements(board Board, startIndex int) ([]string) {
  elements := make([]string, 0)
  gameState, offset := GetGameStateAndOffset(board)

  for i := startIndex; i < len(gameState); i += offset {
    elements = append(elements,gameState[i])
  }

  return elements
}

func (r *BasicRules) DiagonalWinner(board Board) (string) {
    gameState := board.Array()
    offset := board.Offset()

    if r.AllSameSymbols(r.LRDiagonalElements(board)){
      return gameState[0]
    }

    if r.AllSameSymbols(r.RLDiagonalElements(board)){
      return gameState[offset - 1]
    }

  return ""
}

func (r *BasicRules) LRDiagonalElements(board Board) ([]string) {
  elements := make([]string, 0)
  gameState, offset := GetGameStateAndOffset(board)

  for i := 0; i < len(gameState); i += offset+1 {
    elements = append(elements, gameState[i])
  }

  return elements
}

func (r *BasicRules) RLDiagonalElements(board Board) ([]string) {
  elements := make([]string, 0)
  gameState, offset := GetGameStateAndOffset(board)
  newOffset := offset - 1

  for i := newOffset; i < len(gameState)-1; i += newOffset {
    elements = append(elements,gameState[i])
  }

  return elements
}

func (r *BasicRules) AllSameSymbols (data []string) (bool) {
  if data[0] == "" {
    return false
  }

  for i := 1; i < len(data); i++ {
      if data[0] != data[i] {
          return false
      }
  }

  return true
}

func GetGameStateAndOffset (board Board) ([]string, int) {
  gameState := board.Array()
  offset := board.Offset()
  return gameState, offset
}
