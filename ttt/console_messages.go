package ttt

type ConsoleMessages struct {
	introMessage string
	exitMessage  string

	xSymbol string
	oSymbol string

	chooseMovePrompt string
	newGamePrompt    string
	boardTypePrompt  string
	playedTypePrompt string

	gameTieResponse       string
	gameWinnerResponse    string
	invalidChoiceResponse string
	playerTypePrompt      string
	playerTypesResponse   string
	boardTypesResponse    string

	verticalDivider   string
	horizontalDivider string
	emptySpot         string
	spotWidth         int

	yesOption string
}

func (c *ConsoleMessages) BuildMessages() {
	c.introMessage = "Welcome to Tic Tac Go!"
	c.exitMessage = "So Long!"

	c.chooseMovePrompt = "%v, Choose a Move!\n"
	c.newGamePrompt = "Would you like to start a new game? Press (Y) for yes, any other key to exit"
	c.boardTypePrompt = "Choose Board Type:"
	c.playerTypePrompt = "Choose Type for: %v\n"

	c.gameTieResponse = "The Game Has Ended In A Tie\n"
	c.gameWinnerResponse = "The Winner is %v\n"
	c.invalidChoiceResponse = "Whoops, that choice is invalid! Try Again"
	c.boardTypesResponse = "%v. %v\n"
	c.playerTypesResponse = "%v. %v\n"

	c.verticalDivider = " | "
	c.horizontalDivider = "-"
	c.emptySpot = " "
	c.xSymbol = "X"
	c.oSymbol = "O"
	c.spotWidth = 3

	c.yesOption = "y"
}

func (c ConsoleMessages) WinnerSymbol(winner string) string {
	if winner == "x" {
		return c.xSymbol
	} else if winner == "o" {
		return c.oSymbol
	}

	return winner
}

func (c ConsoleMessages) IntroMessage() string {
	return c.introMessage
}

func (c ConsoleMessages) ExitMessage() string {
	return c.exitMessage
}

func (c ConsoleMessages) YesOption() string {
	return c.yesOption
}

func (c ConsoleMessages) HorizontalDivider() string {
	return c.horizontalDivider
}

func (c ConsoleMessages) VerticalDivider() string {
	return c.verticalDivider
}

func (c ConsoleMessages) EmptySpot() string {
	return c.emptySpot
}

func (c ConsoleMessages) ChooseMovePrompt() string {
	return c.chooseMovePrompt
}

func (c ConsoleMessages) GameTieResponse() string {
	return c.gameTieResponse
}

func (c ConsoleMessages) GameWinnerResponse() string {
	return c.gameWinnerResponse
}

func (c ConsoleMessages) NewGamePrompt() string {
	return c.newGamePrompt
}

func (c ConsoleMessages) InvalidChoiceResponse() string {
	return c.invalidChoiceResponse
}

func (c ConsoleMessages) BoardTypePrompt() string {
	return c.boardTypePrompt
}

func (c ConsoleMessages) BoardTypesResponse() string {
	return c.boardTypesResponse
}

func (c ConsoleMessages) PlayerTypePrompt() string {
	return c.playerTypePrompt
}

func (c ConsoleMessages) PlayerTypesResponse() string {
	return c.playerTypesResponse
}

func (c ConsoleMessages) SpotWidth() int {
	return c.spotWidth
}

func (c ConsoleMessages) SetPlayerXSymbol(newSymbol string) {
	c.xSymbol = newSymbol
}

func (c ConsoleMessages) SetPlayerOSymbol(newSymbol string) {
	c.oSymbol = newSymbol
}

func (c ConsoleMessages) PlayerXSymbol() string {
	return c.xSymbol
}

func (c ConsoleMessages) PlayerOSymbol() string {
	return c.oSymbol
}
