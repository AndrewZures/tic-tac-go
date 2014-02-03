package ttt

import "fmt"

type ConsoleUI struct {
}

func (c *ConsoleUI) PrintBoard(boardArray []int) (bool) {

    for i := 0; i < len(boardArray); i++ {
      fmt.Printf("value at %v is %v\n", i, boardArray[i])
    }

    return true
}

