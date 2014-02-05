package ttt

type HumanPlayer struct {
  symbol string
  typeTitle string;
}

func (h *HumanPlayer) NewHumanPlayer(symbol string, typeTitle string){
  h.symbol = symbol
  h.typeTitle = typeTitle
}

func (h *HumanPlayer) MakeMove(board Board) (int) {
  return 1
}

func (h *HumanPlayer) Symbol() (string) {
  return h.symbol
}

func (h *HumanPlayer) Description() (string) {
  return h.typeTitle
}
