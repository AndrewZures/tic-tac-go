package ttt

import "bytes"

type ComputerPlayer struct {
  symbol string;
}


func (h *ComputerPlayer) MakeMove(board Board) (int) {
  return len(board.Array())
}

func (h *ComputerPlayer) miniMax(board Board) (int) {
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
