package ttt

type ConsoleMessages struct {

  chooseMovePrompt string
  newGamePrompt string
  boardTypePrompt string
  playerTypePrompt string

  gameTieResponse string
  gameWinnerResponse string
  invalidChoiceResponse string
  playerTypesResponse string

  verticalDivider string
  horizontalDivider string
  emptySpot string
}

func (c *ConsoleMessages) BuildMessages() {
  c.chooseMovePrompt = "%v, Choose a Move!\n"
  c.gameTieResponse = "The Game Has Ended In A Tie\n"
  c.gameWinnerResponse = "The Winner is %v\n"
  c.newGamePrompt = "Would you like to start a new game? Press (Y) for yes, any other key to exit"
  c.invalidChoiceResponse = "Whoops, that choice is invalid! Try Again"
  c.boardTypePrompt = "Choose Board Type:"
  c.playerTypePrompt = "Choose Type for: %v\n"
  c.playerTypesResponse = "%v. %v\n"

  c.verticalDivider = " | "
  c.horizontalDivider = "-"
  c.emptySpot = " "
}

func (c *ConsoleMessages) ChooseMovePrompt() string {
  return c.chooseMovePrompt
}

func (c *ConsoleMessages) GameTieResponse() string {
  return c.gameTieResponse
}

func (c *ConsoleMessages) GameWinnerResponse() string {
  return c.gameWinnerResponse
}

func (c *ConsoleMessages) NewGamePrompt() string {
  return c.newGamePrompt
}

func (c *ConsoleMessages) InvalidChoiceResponse() string {
  return c.invalidChoiceResponse
}

func (c *ConsoleMessages) BoardTypePrompt() string {
  return c.boardTypePrompt
}

func (c *ConsoleMessages) PlayerTypePrompt() string {
  return c.playerTypePrompt
}

func (c *ConsoleMessages) PlayerTypesResponse() string {
  return c.playerTypesResponse
}
