package ttt

import ("fmt"
        "io"
      )

type ConsoleUI struct {
  Writer io.Writer;
}

func (c *ConsoleUI) PrintBoard(boardArray []int) {

  for i := 0; i < len(boardArray); i++ {
    fmt.Fprintf(c.Writer, "value at %v is %v\n", i, boardArray[i])
  }

}

func (c *ConsoleUI) Print (input string){
  fmt.Fprintln(c.Writer, input)
}

