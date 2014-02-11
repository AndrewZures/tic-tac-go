package ttt

import "math"

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
  bestScore = math.Inf(-1)
  minAlpha = math.Inf(-1)
  maxBeta = math.Inf(1)
  bestMove = -1
  depth = 1

  for _, move := range board.OpenSpots() {

    board.RecordMove(move, h.Symbol())
    score = h.Score(board, h.Symbol(), depth, minAlpha, maxBeta)
    board.RemoveMove(move)

    if score > bestScore {
      bestScore = score
      bestMove = move
    }
  }

  return bestMove
}

func (h *ComputerPlayer) miniMax(board Board, symbol string, depth int, alpha float64, beta float64) (float64) {

  for _, move := range board.OpenSpots() {

    board.RecordMove(move, symbol)
    score := h.Score(board, symbol, depth, alpha, beta)
    board.RemoveMove(move)


    if score > alpha {
      alpha = score
    }
  }

  return alpha
}

func (h *ComputerPlayer) Score(board Board, symbol string, depth int, alpha float64, beta float64) (float64){
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
