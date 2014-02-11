package ttt

import ( "strconv"
        "strings"
      )

type ConsoleUI struct {
  ConsoleIO InOut
  Messages Messages
  BoardFormatter BoardFormatter
}

func (c ConsoleUI) DisplayBoard(board Board) {
  formattedBoard := c.BoardFormatter.FormatBoard(board, c.Messages)
  c.ConsoleIO.Println(formattedBoard)
}

func (c ConsoleUI) SelectMove(player Player, board Board) (int){

  for {
    c.AskUserForMove(player)
    move := c.GetIntegerFromUser()
    if c.ValidateMove(move, board) {
      zeroIndexMove := c.shiftToZerosBasedIndex(move)
      return zeroIndexMove
    } else {
      c.PrintChoiceInvalid()
    }
  }
}

func (c ConsoleUI) AskUserForMove(player Player) {
  c.ConsoleIO.Printf(c.Messages.ChooseMovePrompt(), player.Description())
}

func (c ConsoleUI) ValidateMove(move int, board Board) (bool) {
  availableMoves := board.OpenSpots()
  status := false

  for _, availableMove := range availableMoves {
    if move == availableMove {
      status = true
    }
  }

  return status
}

func (c ConsoleUI) DisplayWinner(winner string) {
  if winner == "tie" {
    c.ConsoleIO.Printf(c.Messages.GameTieResponse())
  } else {
    c.ConsoleIO.Printf(c.Messages.GameWinnerResponse(), winner)
  }
}

func (c ConsoleUI) SelectPlayerChoice(playerList []Player, description string) (Player) {
  c.PrintPlayerTypeQuestion(description)
  c.DisplayPlayerTypes(playerList)
  playerChoice := c.PlayerChoice(playerList)
  return playerChoice
}

func (c ConsoleUI) DisplayPlayerTypes(playerList []Player){

  for index, playerType := range playerList {
    c.ConsoleIO.Printf(c.Messages.PlayerTypesResponse(), (index+1), playerType.Description())
  }

}

func (c ConsoleUI) PrintPlayerTypeQuestion(playerDescription string) {
  c.ConsoleIO.Printf(c.Messages.PlayerTypePrompt(), playerDescription)
}

func (c ConsoleUI) PrintBoardTypeQuestion() {
  c.ConsoleIO.Println(c.Messages.BoardTypePrompt())
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

  for index, boardType := range boardList {
    c.ConsoleIO.Printf(c.Messages.BoardTypesResponse(), (index+1), boardType.Description())
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
    userInput = c.ConsoleIO.Read()
    copiedInput = strings.Repeat(userInput, 1)
    value, err := strconv.ParseInt(copiedInput,0,0)

    if err == nil {
      return int(value)
    } else {
      c.PrintChoiceInvalid()
    }
  }

}

func (c ConsoleUI) AskForNewGame() (bool) {
  c.displayNewGameQuery()
  return c.getNewGameDecision()
}

func (c ConsoleUI) displayNewGameQuery() {
  c.ConsoleIO.Println(c.Messages.NewGamePrompt())
}

func (c ConsoleUI) getNewGameDecision() (bool) {
  response := strings.ToLower(c.ConsoleIO.Read())
  return response == c.Messages.YesResponse()
}

func (c *ConsoleUI) shiftToZerosBasedIndex(onesBasedIndexChoice int) (int) {
  return onesBasedIndexChoice - 1
}

func (c *ConsoleUI) ChoiceValid(choice int, numChoices int) (bool) {
  return choice > 0 && choice <= numChoices
}

func (c ConsoleUI) PrintChoiceInvalid(){
  c.ConsoleIO.Println(c.Messages.InvalidChoiceResponse())
}
