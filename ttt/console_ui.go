package ttt

import ("fmt"
        "io"
        "bufio"
        "strconv"
      )

type ConsoleUI struct {
  Writer io.Writer;
  //Reader io.Reader;
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
  return playerList[0]
}

func (c *ConsoleUI) GetIntegerFromUser() (int) {
  userInput := c.ReadConsole()
  for i := 0; i < 4; i++ {
    value, err := strconv.ParseInt(userInput,0,0)
    if err == nil {
      return int(value)
    } else {
      //c.PrintChoiceInvalid()
      fmt.Println(err)
      userInput = c.ReadConsole()
    }
  }

  return -1
}

func (c *ConsoleUI) PrintChoiceInvalid(){
  fmt.Println("Whoops, that choice is invalid! Try Again")
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

