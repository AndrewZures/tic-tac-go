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
    if board.Status() == h.Symbol(){
      score = float64(1.0 / depth)
    } else if board.Status() == "tie"{
      score = 0
    } else {
      score = -h.miniMax(board, h.opponent("O"), depth)
    }

    board.RemoveMove(openMoves[i])
    fmt.Printf("score = %f for option %v\n", score, openMoves[i])

    if score > bestScore {
      fmt.Printf("replacing bestscore(%f) with new score(%f) from option %v\n", bestScore, score, openMoves[i])
      bestScore = score
      bestMove = openMoves[i]
    }
  }

  return bestMove
}

func (h *ComputerPlayer) miniMax(board Board, symbol string, depth int) (float64) {

  //fmt.Printf("return value is %f, depth is %v\n", depthScore, depth)

  var score float64
  bestScore := float64(-100000000)
  openMoves := board.OpenSpots()

  for i := 0; i < len(openMoves); i++ {
    board.RecordMove(openMoves[i], symbol)

    //fmt.Printf("board status is %v, while symbol is %v", board.Status(), symbol)

      fmt.Printf("status for %v after recorded move is %v", openMoves[i], board.Status())
      fmt.Println(board.Array())
    if board.Status() == symbol {
      score = float64(1.0 / depth)
    } else if board.Status() == h.tempOpponent(symbol) {
      score = -float64(1.0 / depth)
    } else if board.Status() == "tie" {
      score = 0
    } else {
      score = -h.miniMax(board, h.opponent(symbol), depth+1)
    }
    board.RemoveMove(openMoves[i])

    if score > bestScore {
      bestScore = score
    }
  }

  return bestScore
}

func (h *ComputerPlayer) tempOpponent(symbol string) (string) {
  if symbol == "X" {
    return "O"
  } else if symbol == "O" {
    return "X"
  } else {
    return "Z"
  }

}

func (h *ComputerPlayer) opponent(symbol string) (string) {
  if symbol == "X" {
    return "O"
  } else if symbol == "O" {
    return "X"
  } else {
    return "Z"
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
