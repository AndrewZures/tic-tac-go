package ttt

type HumanPlayer struct {
  symbol string
}

func (h *HumanPlayer) MakeMove(board []string) (int) {
  return 1
}

func (h *HumanPlayer) Symbol() (string) {
  return h.symbol
}

func (h *HumanPlayer) SetSymbol(newSymbol string) {
  h.symbol = newSymbol
}