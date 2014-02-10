package ttt

import "strings"

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


func (c *ConsoleMessages) DisplayBoard (board Board) (string) {
  var boardResult string

  gameState := board.Array()
  gameState = c.FormatGameState(gameState)

  for i := 0; i < len(gameState); i++ {

    boardResult += gameState[i]

    if c.LastIndex(board, i) {
      boardResult += "\n"
    } else if c.EndOfRow(board, i) {
      boardResult += c.PrintHorizontalDivider(len(gameState))
    } else {
      boardResult += c.PrintVerticalDivider()
    }
  }

  return boardResult
}

func (c *ConsoleMessages) LastIndex(board Board, index int) (bool) {
  return index == len(board.Array()) - 1
}

func (c *ConsoleMessages) PrintSymbol(symbol string) (string) {
  return symbol
}

func (c *ConsoleMessages) PrintVerticalDivider() (string) {
  return c.verticalDivider
}

func (c *ConsoleMessages) EndOfRow(board Board, index int) (bool) {
  return (index+1) % board.Offset() == 0
}


func (c *ConsoleMessages) PrintHorizontalDivider(dividerLength int) (string) {
  divider := []byte("\n")

  //TODO this may not work for larger board sizes
  for i := 0; i < dividerLength; i++ {
    divider = append(divider, []byte("-")...)
  }

  divider = append(divider, []byte("\n")...)
  return string(divider)
}

func (c *ConsoleMessages) FormatGameState(list []string) ([]string) {
  list = c.ConvertToUpperCase(list)
  list = c.FormatEmptySpots(list)
  return list
}

func (c *ConsoleMessages) ConvertToUpperCase(list []string) ([]string) {
  for i := 0; i < len(list); i++ {
    list[i] = strings.ToUpper(list[i])
  }
  return list
}

func (c *ConsoleMessages) FormatEmptySpots(list []string) ([]string) {

  for i := 0; i < len(list); i++ {
    if list[i] == "" {
      list[i] = " "
    }
  }
  return list
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
