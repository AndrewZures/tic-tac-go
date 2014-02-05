package ttt

import "fmt"

type ComputerPlayer struct {
  symbol string;
}

func (h *ComputerPlayer) MakeMove(board Board) (int) {
  return h.startMiniMax(board)
}

func (h *ComputerPlayer) startMiniMax(board Board) (int) {
  var bestMove int
  var score float64
  var bestScore float64
  bestScore = -100000000.0000
  bestMove = -1
  depth := 1

  openMoves := board.OpenSpots()

  for i := 0; i < len(openMoves); i++ {
    score = h.executeMiniMax(board, h.Symbol(), openMoves[i], depth)

    if score > bestScore {
      bestScore = score
      bestMove = openMoves[i]
    }
  }

  return bestMove
}

func (h *ComputerPlayer) miniMax(board Board, symbol string, depth int) (float64) {
  var score float64
  bestScore := -10000000.0000
  openMoves := board.OpenSpots()

  for i := 0; i < len(openMoves); i++ {

    score = h.executeMiniMax(board, symbol, openMoves[i], depth)

    if score > bestScore {
      bestScore = score
    }
  }

  return bestScore
}

func (h *ComputerPlayer) executeMiniMax(board Board, symbol string, index int, depth int) (float64){
    var score float64

    result := board.RecordMove(index, symbol)
    //board.RecordMove(index, symbol)
    fmt.Printf("player move is initially = %v\n", board.PlayerTurn())
    fmt.Printf("board recording index %v with symbol %v at depth %v, with result %v\n", index, symbol, depth, result)
    fmt.Printf("player move is later = %v\n", board.PlayerTurn())
    fmt.Println(board.Array())
    score = h.GetScore(board, symbol, depth)
    board.RemoveMove(index)

    return score
}


func (h *ComputerPlayer) GetScore(board Board, symbol string, depth int) (float64){
  var score float64

  if board.Status() == symbol{
    score = 1.0 / float64(depth)
  } else if board.Status() == "tie"{
    score = 0
  } else if depth > 7 {
    score = 0
  } else {
    score = -h.miniMax(board, opponent(symbol), depth + 1)
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

func (h *ComputerPlayer) SetSymbol(newSymbol string) {
  h.symbol = newSymbol
}
