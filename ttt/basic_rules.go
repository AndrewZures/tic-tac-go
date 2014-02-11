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

  rowStatus := r.Rows(board)
  columnStatus := r.Columns(board)
  firstDiagonal := r.LRDiagonalElements(board)
  secondDiagonal := r.RLDiagonalElements(board)

  segments := [][][]string{rowStatus, columnStatus, firstDiagonal, secondDiagonal}

  for i := 0; i < len(segments); i++ {
    segmentStatus := r.EvaluateSegment(segments[i])
    if segmentStatus != "" {
      return segmentStatus
    }
  }


  if len(openSpots) == 0 {
    return "tie"
  } else {
    return "inprogress"
  }

}

func (r BasicRules) EvaluateSegment(segment [][]string) (string) {
  for i := 0; i < len(segment); i++ {
    if r.AllSameSymbols(segment[i]) {
      return segment[i][0]
    }
  }
  return ""
}


func (r BasicRules) Rows(board Board) ([][]string) {
  offset := board.Offset()
  rows := make([][]string, offset)

  for i := 0; i < offset; i++ {
    startIndex := i * offset
    rows[i] = r.rowElements(board, startIndex)
  }

  return rows
}

func (r BasicRules) rowElements(board Board, startIndex int) ([]string) {
  gameState, offset := GetGameStateAndOffset(board)
  return gameState[startIndex:startIndex+offset]
}

func (r BasicRules) Columns(board Board) ([][]string) {
  offset := board.Offset()
  columns := make([][]string, offset)

  for i := 0; i < offset; i++ {
    columns[i] = r.columnElements(board, i)
  }

  return columns
}

func (r *BasicRules) columnElements(board Board, startIndex int) ([]string) {
  elements := make([]string, 0)
  gameState, offset := GetGameStateAndOffset(board)

  for i := startIndex; i < len(gameState); i += offset {
    elements = append(elements,gameState[i])
  }

  return elements
}

func (r *BasicRules) LRDiagonalElements(board Board) ([][]string) {
  elements := make([]string, 0)
  gameState, offset := GetGameStateAndOffset(board)
  diagonal := make([][]string, offset)

  for i := 0; i < len(gameState); i += offset+1 {
    elements = append(elements, gameState[i])
  }
  diagonal[0] = elements

  return diagonal
}

func (r *BasicRules) RLDiagonalElements(board Board) ([][]string) {
  elements := make([]string, 0)
  gameState, offset := GetGameStateAndOffset(board)
  diagonal := make([][]string, offset)
  newOffset := offset - 1

  for i := newOffset; i < len(gameState)-1; i += newOffset {
    elements = append(elements,gameState[i])
  }
  diagonal[0] = elements

  return diagonal
}

func (r *BasicRules) AllSameSymbols (data []string) (bool) {
  if len(data) == 0 || data[0] == "" {
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
