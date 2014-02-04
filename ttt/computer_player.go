package ttt

import "bytes"

type ComputerPlayer struct {
  symbol string;
}

func (h *ComputerPlayer) Symbol() (string) {
  return h.symbol
}

func (h *ComputerPlayer) SetSymbol(newSymbol string) {
  h.symbol = newSymbol
}

func (h *ComputerPlayer) MakeMove(board []string) (int) {
  return 1
}

func (h *ComputerPlayer) CompareStrings (a, b string) (bool) {
    aInBytes := []byte(a)
    bInBytes := []byte(b)
    return 0 == bytes.Compare(aInBytes,bInBytes)
}


