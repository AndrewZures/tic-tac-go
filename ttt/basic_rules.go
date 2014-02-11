package ttt

type BasicRules struct {
    playerTurn string;
}

func (r *BasicRules) NewBasicRules(startSymbol string) {
    r.playerTurn = startSymbol
}

func (r BasicRules) Score(board Board) (string) {
  numOpenSpots := len(board.OpenSpots())
  segments := r.BuildSegments(board)
  segmentsStatus := r.ScoreSegments(segments)
  gameStatus := r.gameStatus(segmentsStatus, numOpenSpots)
  return gameStatus
}

func (r BasicRules) BuildSegments(board Board) ([][][]string) {
  segments := [][][]string{r.Rows(board),
                           r.Columns(board),
                           r.LeftDiagonal(board),
                           r.RightDiagonal(board)}

  return segments
}

func (r BasicRules) ScoreSegments(segments [][][]string) (string) {

  for i := 0; i < len(segments); i++ {
    segmentStatus := r.ScoreSegment(segments[i])
    if segmentStatus != "" {
      return segmentStatus
    }
  }

  return ""
}

func (r BasicRules) ScoreSegment(segment [][]string) (string) {

  for i := 0; i < len(segment); i++ {
    if r.AllSameSymbol(segment[i]) {
      return segment[i][0]
    }
  }
  return ""
}

func (r BasicRules) gameStatus (segmentsStatus string, openSpots int) (string) {

  if segmentsStatus != "" {
    return segmentsStatus
  } else if openSpots == 0 {
    return "tie"
  } else {
    return "inprogress"
  }

}

func (r *BasicRules) AllSameSymbol (data []string) (bool) {
  if len(data) == 0 || data[0] == "" {
    return false
  }

  for i := 0; i < len(data); i++ {
      if data[0] != data[i] {
          return false
      }
  }

  return true
}

func (r BasicRules) Rows(board Board) ([][]string) {
  offset := board.Offset()
  rows := make([][]string, offset)

  for i := 0; i < offset; i++ {
    startIndex := i * offset
    rows[i] = r.row(board, startIndex)
  }

  return rows
}

func (r BasicRules) row(board Board, startIndex int) ([]string) {
  gameState, offset := GetGameStateAndOffset(board)
  return gameState[startIndex:startIndex+offset]
}

func (r BasicRules) Columns(board Board) ([][]string) {
  offset := board.Offset()
  columns := make([][]string, offset)

  for i := 0; i < offset; i++ {
    columns[i] = r.column(board, i)
  }

  return columns
}

func (r *BasicRules) column(board Board, startIndex int) ([]string) {
  elements := make([]string, 0)
  gameState, offset := GetGameStateAndOffset(board)

  for i := startIndex; i < len(gameState); i += offset {
    elements = append(elements,gameState[i])
  }

  return elements
}

func (r *BasicRules) LeftDiagonal(board Board) ([][]string) {
  elements := make([]string, 0)
  gameState, offset := GetGameStateAndOffset(board)
  diagonal := make([][]string, offset)

  for i := 0; i < len(gameState); i += offset+1 {
    elements = append(elements, gameState[i])
  }
  diagonal[0] = elements

  return diagonal
}

func (r *BasicRules) RightDiagonal(board Board) ([][]string) {
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
  return r.Score(board) != "inprogress"
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

