package ttt

import "math"

type ComputerPlayer struct {
  symbol string;
  typeTitle string;
  depthLimit int
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

func (h *ComputerPlayer) setupMiniMax(board Board) (float64, float64, float64, int, int) {
  h.depthLimit = h.setDepthLimit(len(board.OpenSpots()))
  bestScore, minAlpha := math.Inf(-1), math.Inf(-1)
  maxBeta := math.Inf(1)
  bestMove := -1
  depth := 1

  return bestScore, minAlpha, maxBeta, bestMove, depth
}

func (h *ComputerPlayer) startMiniMax(board Board) (int) {
  bestScore, minAlpha, maxBeta, bestMove, depth := h.setupMiniMax(board)

  for _, move := range board.OpenSpots() {

    score := h.placeAndScore(board, h.Symbol(), move, depth, minAlpha, maxBeta)

    if score > bestScore {
      bestScore = score
      bestMove = move
    }
  }

  return bestMove
}

func (h ComputerPlayer) placeAndScore(board Board, player string, move, depth int, alpha, beta float64) (float64) {

    board.RecordMove(move, player)
    score := h.Score(board, player, depth, alpha, beta)
    board.RemoveMove(move)
    return score
}

func (h *ComputerPlayer) Score(board Board, symbol string, depth int, alpha, beta float64) (float64){
  var score float64

  if h.depthLimitReached(depth) || h.pruneComplete(alpha, beta) {
    return 0
  }

  gameStatus := h.rules.Score(board)

  if h.winner(gameStatus, symbol) {
    score = h.depthScore(depth)
  } else if gameStatus == "tie"{
    score = 0
  } else {
    score = -h.miniMax(board, opponent(symbol), depth + 1, -beta, -alpha)
  }

  return score
}

func (h *ComputerPlayer) miniMax(board Board, player string, depth int, alpha float64, beta float64) (float64) {

  for _, move := range board.OpenSpots() {

    score := h.placeAndScore(board, player, move, depth, alpha, beta)

    if score > alpha {
      alpha = score
    }
  }

  return alpha
}

func (h ComputerPlayer) winner(gameStatus, symbol string) (bool) {
  return gameStatus == symbol
}

func (h ComputerPlayer) depthScore (depth int) (float64) {
    return 1.0 / float64(depth)
}

func (h ComputerPlayer) depthLimitReached(depth int) (bool) {
  return depth > h.depthLimit
}

func (h ComputerPlayer) pruneComplete(alpha, beta float64) (bool) {
  return alpha >= beta
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

func (h ComputerPlayer) setDepthLimit(numOpenSpots int) (int) {
  switch {
  case numOpenSpots > 18:
    return 4
  case numOpenSpots > 13:
    return 5
  case numOpenSpots > 9:
    return 7
  }

  return 10
}

func (h *ComputerPlayer) Symbol() (string) {
  return h.symbol
}

func (h *ComputerPlayer) Description() (string) {
  return h.typeTitle
}
