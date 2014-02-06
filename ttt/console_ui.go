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

func (c *ConsoleUI) DisplayBoard(board Board) {
  gameState := board.Array()
  gameState = c.FormatGameState(gameState)


  for i := 0; i < len(gameState); i++ {

    fmt.Fprintf(c.Writer, "%v", gameState[i])

    if (i+1) % board.Offset() == 0 {
        fmt.Fprintf(c.Writer, "\n")
    } else {
        fmt.Fprintf(c.Writer, " | ")
    }
  }
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

func (c *ConsoleUI) SelectPlayerChoice(playerList []Player) (Player) {
  c.DisplayPlayerTypes(playerList)
  return c.PlayerChoice(playerList)
}

func (c *ConsoleUI) DisplayPlayerTypes(playerList []Player){

  for i := 0; i < len(playerList); i++ {
    fmt.Fprintln(c.Writer, playerList[i].Description())
  }

}

func (c *ConsoleUI) PlayerChoice(playerList []Player) (Player) {
  var userChoice int

  for {
    userChoice = c.GetIntegerFromUser()

    if c.ChoiceValid(userChoice, len(playerList)){
      userChoice = c.shiftToZerosBasedIndex(userChoice)
      return playerList[userChoice]
    } else {
      c.PrintChoiceInvalid()
    }
  }
}

func (c *ConsoleUI) GetIntegerFromUser() (int) {
  var userInput string

  for {
    userInput = c.ReadConsole()
    value, err := strconv.ParseInt(userInput,0,0)

    if err == nil {
      return int(value)
    } else {
      c.PrintChoiceInvalid()
    }
  }
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

