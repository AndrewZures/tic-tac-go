package ttt

import ("fmt"
        "io"
      )

type ConsoleUI struct {
  Writer io.Writer;
}

func (c *ConsoleUI) PrintBoard(boardArray []int, preferredBreak int) {

  //TODO Refactor, Esp for Readability
  for i := 0; i < len(boardArray); i++ {
    fmt.Fprintf(c.Writer, "%v ", boardArray[i])
    if (i+1) % preferredBreak == 0 {
        fmt.Fprintf(c.Writer, "\n")
    }
  }

}

func (c *ConsoleUI) Print (input string){
  fmt.Fprintln(c.Writer, input)
}

