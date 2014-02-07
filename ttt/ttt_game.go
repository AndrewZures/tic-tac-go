package ttt

type TTTGame struct {
  status string;
  userInterface UserInterface
  factory Factory
}

//TODO complete
func (g *TTTGame) SetupNewGame(console UserInterface, factory Factory) (Board, Player, Player) {
  playerTypes := factory.PlayerTypes()
  player1Template := console.SelectPlayerChoice(playerTypes, "Player 1")
  console.DisplayPlayerTypes(playerTypes)
  player2Template := console.SelectPlayerChoice(playerTypes, "Player 2")

  player1 := factory.Player(player1Template)
  player2 := factory.Player(player2Template)
  board := new(BasicBoard)
  board.NewBoard("X")

  return board, player1, player2

}

func (g TTTGame) Run(console UserInterface, factory Factory) {
  board, player1, player2 := g.SetupNewGame(console, factory)

  currentPlayer := player1

  for i := 0; i < 100; i++ {
    console.DisplayBoard(board)
    currentMove := currentPlayer.MakeMove(board)
    board.RecordMove(currentMove, currentPlayer.Symbol())

    if board.Status() != "inprogress" {
      console.DisplayWinner(board.Winner())
      return
    }

    if currentPlayer == player1{
      currentPlayer = player2
    } else {
      currentPlayer = player1
    }
  }
}



