package ttt

import "bytes"
import "fmt"

type ComputerPlayer struct {
  symbol string;
}

func (h *ComputerPlayer) MakeMove(board Board) (int) {
  return h.miniMaxTop(board)
}

func (h *ComputerPlayer) miniMaxTop(board Board) (int) {
  var bestMove int
  var score float64
  var bestScore float64
  bestScore = -100000000
  bestMove = -1
  depth := 1

  openMoves := board.OpenSpots()

  for i := 0; i < len(openMoves); i++ {

    board.RecordMove(openMoves[i], h.Symbol())

    score = h.GetScore(board, h.Symbol(), depth)


    board.RemoveMove(openMoves[i])
    fmt.Printf("score = %f for option %v\n", score, openMoves[i])

    //recordBestIndexPossibility(openMoves[i], score)

    if score > bestScore {
      //fmt.Printf("replacing bestscore(%f) with new score(%f) from option %v\n", bestScore, score, openMoves[i])
      bestScore = score
      bestMove = openMoves[i]
    }
  }

  return bestMove
}

func (h *ComputerPlayer) GetScore(board Board, symbol string, depth int) (float64){
  var score float64

  if board.Status() == symbol{
    score = float64(1.0 / depth)
  } else if board.Status() == "tie"{
    score = 0
  } else {
    score = -h.miniMax(board, opponent(symbol), depth)
  }

  return score
}

func (h *ComputerPlayer) miniMax(board Board, symbol string, depth int) (float64) {
  var score float64
  bestScore := float64(-100000000)
  openMoves := board.OpenSpots()

  for i := 0; i < len(openMoves); i++ {

    board.RecordMove(openMoves[i], symbol)

    if board.Status() == symbol {
      score = float64(1.0 / depth)
    } else if board.Status() == "tie" {
      score = 0
    } else {
      score = -h.miniMax(board, opponent(symbol), depth+1)
    }

    board.RemoveMove(openMoves[i])

    if score > bestScore {
      bestScore = score
    }
  }

  return bestScore
}

func opponent(symbol string) (string) {
  if symbol == "X" {
    return "O"
  } else if symbol == "O" {
    return "X"
  } else {
    return "other"
  }
}

func (h *ComputerPlayer) CompareStrings (a, b string) (bool) {
  aInBytes := []byte(a)
  bInBytes := []byte(b)
  return 0 == bytes.Compare(aInBytes,bInBytes)
}

func (h *ComputerPlayer) Symbol() (string) {
  return h.symbol
}

func (h *ComputerPlayer) SetSymbol(newSymbol string) {
  h.symbol = newSymbol
}
