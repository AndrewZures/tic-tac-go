package ttt

import ( "strconv"
        "strings"
      )

type ConsoleUI struct {
  InOut InOutInterface
}

func (c ConsoleUI) DisplayBoard(board Board) {
  gameState := board.Array()
  gameState = c.FormatGameState(gameState)


  for i := 0; i < len(gameState); i++ {

    c.PrintSymbol(gameState[i])

    if c.EndOfRow(board, i) {
      c.PrintHorizontalDivider(len(gameState))
    } else {
      c.PrintVerticalDivider()
    }
  }
}

func (c ConsoleUI) SelectMove(player Player, board Board) (int){

  for {
    c.AskUserForMove(player)
    move := c.GetIntegerFromUser()
    if c.ValidateMove(move, board) {
      return move
    } else {
      c.PrintChoiceInvalid()
    }
  }
}

func (c ConsoleUI) AskUserForMove(player Player) {
  c.InOut.Printf("%v, Choose a Move!\n", player.Description())
}

func (c ConsoleUI) ValidateMove(move int, board Board) (bool) {
  availableMoves := board.OpenSpots(board.Array())
  status := false

  for i := 0; i < len(availableMoves); i++ {
    if move == availableMoves[i] {
      status = true
    }
  }

  return status
}

func (c ConsoleUI) DisplayWinner(winner string) {
  if winner == "tie"{
    c.InOut.Printf("The Game Has Ended In A Tie\n")
  } else {
    c.InOut.Printf("The Winner is %v\n", winner)
  }
}

func (c *ConsoleUI) EndOfRow(board Board, index int) (bool) {
  return (index+1) % board.Offset() == 0
}

func (c *ConsoleUI) PrintSymbol(symbol string) {
  c.InOut.Printf("%v", symbol)
}

func (c *ConsoleUI) PrintHorizontalDivider(dividerLength int) {
  divider := []byte("\n")

  //TODO this may not work for larger board sizes
  for i := 0; i < dividerLength; i++ {
    divider = append(divider, []byte("-")...)
  }

  divider = append(divider, []byte("\n")...)
  c.InOut.Printf(string(divider))
}

func (c *ConsoleUI) PrintVerticalDivider() {
  c.InOut.Printf(" | ")
}

func (c *ConsoleUI) FormatGameState(list []string) ([]string) {
  list = c.ConvertToUpperCase(list)
  list = c.FormatEmptySpots(list)
  return list
}

func (c *ConsoleUI) ConvertToUpperCase(list []string) ([]string) {
  for i := 0; i < len(list); i++ {
    list[i] = strings.ToUpper(list[i])
  }
  return list
}

func (c *ConsoleUI) FormatEmptySpots(list []string) ([]string) {

  for i := 0; i < len(list); i++ {
    if list[i] == "" {
      list[i] = " "
    }
  }
  return list
}

func (c ConsoleUI) SelectPlayerChoice(playerList []Player, description string) (Player) {
  c.PrintPlayerTypeQuestion(description)
  c.DisplayPlayerTypes(playerList)
  playerChoice := c.PlayerChoice(playerList)
  return playerChoice
}

func (c ConsoleUI) DisplayPlayerTypes(playerList []Player){

  for i := 0; i < len(playerList); i++ {
    c.InOut.Printf("%v. %v\n", (i+1), playerList[i].Description())
  }

}

func (c ConsoleUI) PrintPlayerTypeQuestion(playerDescription string) {
  c.InOut.Printf("Choose Type for: %v\n", playerDescription)
}

func (c ConsoleUI) PrintBoardTypeQuestion() {
  c.InOut.Println("Choose Board Type:")
}


func (c *ConsoleUI) PlayerChoice(playerList []Player) (Player) {

  for {
    userChoice := c.GetIntegerFromUser()

    if c.ChoiceValid(userChoice, len(playerList)){
      userChoice = c.shiftToZerosBasedIndex(userChoice)
      return playerList[userChoice]
    } else {
      c.PrintChoiceInvalid()
    }
  }
}

func (c ConsoleUI) DisplayBoardTypes(boardList []Board){

  for i := 0; i < len(boardList); i++ {
    c.InOut.Printf("%v. %v\n", (i+1), boardList[i].Description())
  }

}

func (c ConsoleUI) SelectBoardChoice(boardList []Board) (Board) {
  c.PrintBoardTypeQuestion()
  c.DisplayBoardTypes(boardList)
  boardChoice := c.BoardChoice(boardList)
  return boardChoice
}

func (c ConsoleUI) BoardChoice(boardList []Board) (Board) {
  for {
    userChoice := c.GetIntegerFromUser()

    if c.ChoiceValid(userChoice, len(boardList)) {
      userChoice = c.shiftToZerosBasedIndex(userChoice)
      return boardList[userChoice]
    } else {
      c.PrintChoiceInvalid()
    }
  }
}

func (c ConsoleUI) GetIntegerFromUser() (int) {
  var userInput string
  var copiedInput string

  for {
    userInput = c.InOut.Read()
    copiedInput = strings.Repeat(userInput, 1)
    value, err := strconv.ParseInt(copiedInput,0,0)

    if err == nil {
      return int(value)
    } else {
      c.PrintChoiceInvalid()
    }
  }

  return -1
}

func (c ConsoleUI) AskForNewGame() (bool) {
  c.displayNewGameQuery()
  response := c.InOut.Read()
  if response == "y" {
    return true
  } else {
    return false
  }
}

func (c ConsoleUI) displayNewGameQuery() {
  c.InOut.Println("Would you like to start a new game? Press (Y) for yes, any other key to exit")
}

func (c *ConsoleUI) shiftToZerosBasedIndex(onesBasedIndexChoice int) (int) {
  return onesBasedIndexChoice - 1
}

func (c *ConsoleUI) ChoiceValid(choice int, numChoices int) (bool) {
  return choice > 0 && choice <= numChoices
}

func (c ConsoleUI) PrintChoiceInvalid(){
  c.InOut.Println("Whoops, that choice is invalid! Try Again")
}


