package ttt

//import "fmt"

type TTTGame struct {
  status string;
  userInterface UserInterface
  factory Factory
}


func (g TTTGame) Run(console UserInterface, factory Factory) {
  board, player1, player2 := g.SetupNewGame(console, factory)

  currentPlayer := player1

  for {
    console.DisplayBoard(board)
    currentMove := currentPlayer.MakeMove(board)
    validMove := board.RecordMove(currentMove, currentPlayer.Symbol())

    if !validMove {
      console.PrintChoiceInvalid()
    } else {

      if board.GameOver() {
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
}

func (g TTTGame) SetupNewGame(console UserInterface, factory Factory) (Board, Player, Player) {
  playerTypes := factory.PlayerTypes()
  player1Template := console.SelectPlayerChoice(playerTypes, "Player 1")
  player2Template := console.SelectPlayerChoice(playerTypes, "Player 2")



  player1 := factory.Player(player1Template)
  player2 := factory.Player(player2Template)


  player1.SetSymbol("x")
  player2.SetSymbol("o")

  boardTypes := factory.BoardTypes()
  boardTemplate := console.SelectBoardChoice(boardTypes)
  board := factory.Board(boardTemplate, player1.Symbol())

  return board, player1, player2
}
