package ttt

type TTTGame struct {
  status string;
  userInterface UserInterface
  factory Factory
}

func (g TTTGame) Run(console UserInterface, factory Factory) {
  newGame := true

  for newGame == true {
    newGame = g.RunGame(console, factory)
  }
}


func (g TTTGame) RunGame(console UserInterface, factory Factory) (bool) {
  board, player1, player2, rules := g.SetupNewGame(console, factory)

  currentPlayer := player1

  for {
    console.DisplayBoard(board)
    currentMove := currentPlayer.MakeMove(board)
    validMove := board.RecordMove(currentMove, currentPlayer.Symbol())

    if !validMove {
      console.PrintChoiceInvalid()
    } else {

      if rules.GameOver(board) {
        console.DisplayBoard(board)
        console.DisplayWinner(rules.Winner(board))
        newGame := console.AskForNewGame()
        return newGame
      }

      if currentPlayer == player1{
        currentPlayer = player2
      } else {
        currentPlayer = player1
      }
    }

  }
  return true
}

func (g TTTGame) SetupNewGame(userInterface UserInterface, factory Factory) (Board, Player, Player, Rules) {
  playerTypes := factory.PlayerTypes(userInterface)
  player1Template := userInterface.SelectPlayerChoice(playerTypes, "Player 1")
  player2Template := userInterface.SelectPlayerChoice(playerTypes, "Player 2")

  player1 := factory.Player(player1Template, userInterface)
  player2 := factory.Player(player2Template, userInterface)

  rules := Rules(new(BasicRules))


  player1.SetSymbol("x")
  player2.SetSymbol("o")

  boardTypes := factory.BoardTypes()
  boardTemplate := userInterface.SelectBoardChoice(boardTypes)
  board := factory.Board(boardTemplate)

  return board, player1, player2, rules
}
