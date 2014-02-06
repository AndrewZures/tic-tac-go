package ttt

import ("fmt"
        "io"
        "bufio"
        "strconv"
      )

type ConsoleUI struct {
  Writer io.Writer;
  Reader bufio.Reader;
}

func (c *ConsoleUI) DisplayBoard(boardArray []int, preferredBreak int) {

  //TODO Refactor, Esp for Readability
  for i := 0; i < len(boardArray); i++ {
    fmt.Fprintf(c.Writer, "%v ", boardArray[i])
    if (i+1) % preferredBreak == 0 {
        fmt.Fprintf(c.Writer, "\n")
    }
  }
}

func (c *ConsoleUI) SelectPlayerChoice(playerList []Player) (Player) {
  c.DisplayPlayerTypes(playerList)
  return c.PlayerChoice(playerList)
  return playerList[0]
}

func (c *ConsoleUI) PlayerChoice(playerList []Player) (Player) {
  var userChoice int

  for {
    userChoice = c.GetIntegerFromUser()
    if c.ChoiceValid(userChoice, len(playerList)){
      userChoice = c.shiftOnesBasedToZerosBased(userChoice)
      return playerList[userChoice]
    } else {
      c.PrintChoiceInvalid()
      userChoice = c.GetIntegerFromUser()
    }
  }
}

func (c *ConsoleUI) shiftOnesBasedToZerosBased(onesBasedIndexChoice int) (int) {
  zerosBasedIndexChoice := onesBasedIndexChoice - 1
  return zerosBasedIndexChoice
}

func (c *ConsoleUI) ChoiceValid(choice int, numChoices int) (bool) {
  return choice > 0 && choice <= numChoices
}

func (c *ConsoleUI) GetIntegerFromUser() (int) {
  userInput := c.ReadConsole()
  for {
    value, err := strconv.ParseInt(userInput,0,0)

    if err == nil {
      return int(value)
    } else {
      c.PrintChoiceInvalid()
      userInput = c.ReadConsole()
    }
  }
}

func (c *ConsoleUI) PrintChoiceInvalid(){
  fmt.Fprintln(c.Writer, "Whoops, that choice is invalid! Try Again")
}

func (c *ConsoleUI) DisplayPlayerTypes(playerList []Player){

  for i := 0; i < len(playerList); i++ {
    fmt.Fprintln(c.Writer, playerList[i].Description())
  }

}

func (c *ConsoleUI) ReadConsole() (string) {
  line, _, _ := c.Reader.ReadLine()
    return string(line)
}

func (c *ConsoleUI) Print (input string){
  fmt.Fprintln(c.Writer, input)
}

