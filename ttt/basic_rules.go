package ttt

type BasicRules struct {
    playerTurn string;
}

func (b *BasicRules) NewBasicRules(startSymbol string) {
    b.playerTurn = startSymbol
}

func (b *BasicRules) Winner(board Board) (string) {
  if board.Status() == "inprogress" {
    return "no winner yet"
  } else {
    return board.Status()
  }
}

func (b BasicRules) GameOver(board Board) (bool){
  return board.Status() != "inprogress"
}

func (b *BasicRules) toggleTurn() {
  if b.playerTurn == "x" {
    b.playerTurn = "o"
  } else if b.playerTurn == "o" {
    b.playerTurn = "x"
  } else {
    b.playerTurn = "z"
  }
}

func (b *BasicRules) PlayerTurn() (string) {
  return b.playerTurn
}

func (b *BasicRules) IsPlayerTurn (player Player) (bool) {
  return b.PlayerTurn() == player.Symbol()
}

func (b *BasicRules) Score(board Board) (string) {
  gameState := board.Array()
  openSpots := board.OpenSpots(gameState)

  rowStatus := b.RowWinner(board)
  if  rowStatus != "" {
    return rowStatus
  }

  columnStatus := b.ColumnWinner(board)
  if  columnStatus != "" {
    return columnStatus
  }

  diagonalStatus := b.DiagonalWinner(board)
  if  diagonalStatus != "" {
    return diagonalStatus
  }

  if len(openSpots) == 0 {
    return "tie"
  } else {
    return "inprogress"
  }

}

func (b *BasicRules) RowWinner(board Board) (string) {
  var row []string
  gameState, offset := GetGameStateAndOffset(board)

  for i := 0; i < len(gameState); i += offset {
    row = b.rowElements(board, i)
    if b.AllSameSymbols(row) == true {
      return gameState[i]
    }
  }

  return ""
}

func (b *BasicRules) rowElements(board Board, startIndex int) ([]string) {
  gameState := board.Array()
  return gameState[startIndex:startIndex+board.Offset()]
}

func (b *BasicRules) ColumnWinner(board Board) (string) {
  var column []string
  gameState, offset := GetGameStateAndOffset(board)

  for i := 0; i < offset; i++ {
    column = b.columnElements(board, i)
    if b.AllSameSymbols(column) == true {
      return gameState[i]
    }
  }

  return ""
}

func (b *BasicRules) columnElements(board Board, startIndex int) ([]string) {
  elements := make([]string, 0)
  gameState, offset := GetGameStateAndOffset(board)

  for i := startIndex; i < len(gameState); i += offset {
    elements = append(elements,gameState[i])
  }

  return elements
}

func (b *BasicRules) DiagonalWinner(board Board) (string) {
    gameState := board.Array()
    offset := board.Offset()

    if b.AllSameSymbols(b.LRDiagonalElements(board)){
      return gameState[0]
    }

    if b.AllSameSymbols(b.RLDiagonalElements(board)){
      return gameState[offset - 1]
    }

  return ""
}

func (b *BasicRules) LRDiagonalElements(board Board) ([]string) {
  elements := make([]string, 0)
  gameState, offset := GetGameStateAndOffset(board)

  for i := 0; i < len(gameState); i += offset+1 {
    elements = append(elements, gameState[i])
  }

  return elements
}

func (b *BasicRules) RLDiagonalElements(board Board) ([]string) {
  elements := make([]string, 0)
  gameState, offset := GetGameStateAndOffset(board)
  newOffset := offset - 1

  for i := newOffset; i < len(gameState)-1; i += newOffset {
    elements = append(elements,gameState[i])
  }

  return elements
}

func (b *BasicRules) AllSameSymbols (data []string) (bool) {
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
