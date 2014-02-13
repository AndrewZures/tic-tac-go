package ttt

type TTTGame struct{}

func (g TTTGame) Run(console UserInterface, factory Factory) {
	newGame := true

	console.DisplayIntroMessage()

	for newGame == true {
		g.RunGame(console, factory)
		newGame = console.QueryNewGame()
	}

	console.DisplayExitMessage()

}

func (g TTTGame) RunGame(console UserInterface, factory Factory) {
	board, player1, player2, rules := g.SetupNewGame(console, factory)

	currentPlayer := player1

	for {
		console.DisplayBoard(board)

		currentMove := currentPlayer.MakeMove(board)
		validMove := board.ValidMove(currentMove)

		if !validMove {
			console.DisplayChoiceInvalid()
		} else {
			board.RecordMove(currentMove, currentPlayer.Symbol())

			if rules.GameOver(board) {
				console.DisplayBoard(board)
				console.DisplayWinner(rules.Winner(board))
				break
			}

			currentPlayer = g.getOpponent(currentPlayer, player1, player2)
		}

	}
}

func (g TTTGame) SetupNewGame(userInterface UserInterface, factory Factory) (Board, Player, Player, Rules) {

	rules := Rules(new(BasicRules))

	player1, player2 := g.getPlayers(factory, userInterface, rules)
	board := g.getBoard(factory, userInterface)

	return board, player1, player2, rules
}

func (g TTTGame) getBoard(factory Factory, userInterface UserInterface) Board {
	boardTypes := factory.BoardTypes()
	boardTemplate := userInterface.QueryBoardChoice(boardTypes)
	return factory.Board(boardTemplate)
}

func (g TTTGame) getPlayers(factory Factory, userInterface UserInterface, rules Rules) (Player, Player) {

	player1Template, player2Template := g.getPlayerTemplates(factory, userInterface, rules)

	player1 := factory.Player(player1Template, userInterface, rules)
	player2 := factory.Player(player2Template, userInterface, rules)

	player1.SetSymbol("x")
	player2.SetSymbol("o")

	return player1, player2
}

func (g TTTGame) getPlayerTemplates(factory Factory, userInterface UserInterface, rules Rules) (Player, Player) {
	playerTypes := factory.PlayerTypes(userInterface, rules)
	player1Template := userInterface.QueryPlayerChoice(playerTypes, "Player X")
	player2Template := userInterface.QueryPlayerChoice(playerTypes, "Player O")
	return player1Template, player2Template
}

func (g TTTGame) getOpponent(currentPlayer, player1, player2 Player) Player {

	if currentPlayer == player1 {
		return player2
	}
	return player1
}
