package ttt_test

import (
  . "github.com/andrewzures/tictactoe/ttt"
  "fmt"
  "bytes"
)


func GenerateBoard(startSymbol string, gameState []string) (Board) {
  newBoard := Board(new(BasicBoard))
  newBoard.NewBoard(startSymbol)
  newBoard.SetArray(gameState)
  return newBoard
}

func SetMockInput(reader *bytes.Buffer, input string){
    fmt.Fprintf(reader, input)
}
