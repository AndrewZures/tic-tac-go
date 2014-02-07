package ttt

import ("fmt"
        "io"
        "bufio"
        "strconv"
        "strings"
      )

type ConsoleUI struct {
  Writer io.Writer;
  Reader bufio.Reader;
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

func (c ConsoleUI) DisplayWinner(winner string) {
  fmt.Fprintf(c.Writer, "The Winner is ", winner)
}

func (c *ConsoleUI) EndOfRow(board Board, index int) (bool) {
  return (index+1) % board.Offset() == 0
}

func (c *ConsoleUI) PrintSymbol(symbol string) {
  fmt.Fprintf(c.Writer, "%v", symbol)
}

func (c *ConsoleUI) PrintHorizontalDivider(dividerLength int) {
  divider := []byte("\n")

  //TODO this may not work for larger board sizes
  for i := 0; i < dividerLength; i++ {
    divider = append(divider, []byte("-")...)
  }

  divider = append(divider, []byte("\n")...)
  fmt.Fprintf(c.Writer, string(divider))
}

func (c *ConsoleUI) PrintVerticalDivider() {
  fmt.Fprintf(c.Writer, " | ")
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
  c.DisplayPlayerTypes(playerList)
  return c.PlayerChoice(playerList)
}

func (c ConsoleUI) DisplayPlayerTypes(playerList []Player){

  for i := 0; i < len(playerList); i++ {
    fmt.Fprintln(c.Writer, playerList[i].Description())
  }

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

func (c *ConsoleUI) GetIntegerFromUser() (int) {

  for i := 0; i < 4; i++ {
    userInput := c.ReadConsole()
    value, err := strconv.ParseInt(userInput,0,0)

    if err == nil {
      return int(value)
    } else {
      c.PrintChoiceInvalid()
    }
  }

  return -1
}

func (c *ConsoleUI) shiftToZerosBasedIndex(onesBasedIndexChoice int) (int) {
  return onesBasedIndexChoice - 1
}

func (c *ConsoleUI) ChoiceValid(choice int, numChoices int) (bool) {
  return choice > 0 && choice <= numChoices
}

func (c *ConsoleUI) PrintChoiceInvalid(){
  fmt.Fprintln(c.Writer, "Whoops, that choice is invalid! Try Again")
}


func (c *ConsoleUI) ReadConsole() (string) {
  line, _, _ := c.Reader.ReadLine()
  return string(line)
}

func (c *ConsoleUI) Print (input string){
  fmt.Fprintln(c.Writer, input)
}
