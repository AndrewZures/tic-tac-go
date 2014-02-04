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
 // bestMoves := make([]int, 0)
  bestScore := -100000000

  openMoves := board.OpenSpots()
  fmt.Println("openMoves = ", openMoves)

  for i := 0; i < len(openMoves); i++ {
      board.RecordMove(openMoves[i], h.Symbol())
      score := -h.miniMax(board, h.opponent(), 1)
      board.RemoveMove(openMoves[i])
      fmt.Println("score = ", score)

      if score > bestScore {
        bestScore = score
      }
  }

  return bestScore
}

func (h *ComputerPlayer) miniMax(board Board, symbol string, depth int) (int) {
  bestScore := -100000000
  openMoves := board.OpenSpots()

  for i := 0; i < len(openMoves); i++ {
    board.RecordMove(openMoves[i], h.Symbol())
    score := -h.miniMax(board, h.opponent(), 1)
    board.RemoveMove(openMoves[i])

    if score > bestScore {
      bestScore = score
    }
  }

  return bestScore
}

func (h *ComputerPlayer) opponent() (string) {
  if h.Symbol() == "X" {
    return "O"
  } else {
    return "X"
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
