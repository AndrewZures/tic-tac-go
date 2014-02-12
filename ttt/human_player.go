package ttt

type HumanPlayer struct {
	symbol        string
	typeTitle     string
	userInterface UserInterface
}

func (h *HumanPlayer) NewHumanPlayer(symbol string, typeTitle string, userInterface UserInterface) {
	h.symbol = symbol
	h.typeTitle = typeTitle
	h.userInterface = userInterface
}

func (h *HumanPlayer) SetSymbol(thisSymbol string) {
	h.symbol = thisSymbol
}

func (h *HumanPlayer) MakeMove(board Board) int {
	return h.userInterface.SelectMove(h, board)
}

func (h *HumanPlayer) Symbol() string {
	return h.symbol
}

func (h *HumanPlayer) Description() string {
	return h.typeTitle
}
