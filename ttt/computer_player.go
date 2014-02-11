package ttt

type ComputerPlayer struct {
  symbol string;
  typeTitle string;
  rules Rules
}

func (h *ComputerPlayer) SetSymbol(symbol string) {
  h.symbol = symbol
}

func (h *ComputerPlayer) NewComputerPlayer(symbol string, typeTitle string, rules Rules){
  h.symbol = symbol
  h.typeTitle = typeTitle
  h.rules = rules
}

func (h *ComputerPlayer) MakeMove(board Board) (int) {
  return h.startMiniMax(board)
}

func (h *ComputerPlayer) startMiniMax(board Board) (int) {
  var bestMove, depth int
  var score, bestScore, minAlpha, maxBeta float64
  bestScore = -100000000.0000
  minAlpha = -100000000.0000
  maxBeta = 100000000.0000
  bestMove = -1
  depth = 1

  openMoves := board.OpenSpots()

  for i := 0; i < len(openMoves); i++ {
    score = h.executeMiniMax(board, h.Symbol(), openMoves[i], depth, minAlpha, maxBeta)

    if score > bestScore {
      bestScore = score
      bestMove = openMoves[i]
    }
  }

  return bestMove
}

func (h *ComputerPlayer) miniMax(board Board, symbol string, depth int, alpha float64, beta float64) (float64) {
  var score float64
  openMoves := board.OpenSpots()


  for i := 0; i < len(openMoves); i++ {

    board.RecordMove(openMoves[i], symbol)
    score = h.GetScore(board, symbol, depth, alpha, beta)
    board.RemoveMove(openMoves[i])


    if score > alpha {
      alpha = score
    }
  }

  return alpha
}

func (h *ComputerPlayer) executeMiniMax(board Board, symbol string, index int, depth int, alpha float64, beta float64) (float64){
    var score float64

    board.RecordMove(index, symbol)
    score = h.GetScore(board, symbol, depth, alpha, beta)
    board.RemoveMove(index)

    return score
}

func (h *ComputerPlayer) GetScore(board Board, symbol string, depth int, alpha float64, beta float64) (float64){
  var score float64

  if alpha >= beta {
    return 0
  }

  gameStatus := h.rules.Score(board)

  if gameStatus == symbol{
    score = 1.0 / float64(depth)
  } else if gameStatus == "tie"{
    score = 0
  } else {
    score = -h.miniMax(board, opponent(symbol), depth + 1, -beta, -alpha)
  }

  return score
}

func opponent(symbol string) (string) {
  opponentSymbol := "other"
  if symbol == "o" {
    opponentSymbol = "x"
  } else if symbol == "x" {
    opponentSymbol = "o"
  }

  return opponentSymbol
}

func (h *ComputerPlayer) Symbol() (string) {
  return h.symbol
}

func (h *ComputerPlayer) Description() (string) {
  return h.typeTitle
}
