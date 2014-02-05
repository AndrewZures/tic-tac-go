package ttt_test

import (
  . "github.com/andrewzures/tictactoe/ttt"
)


func GenerateBoard(startSymbol string, gameState []string) (Board) {
  newBoard := Board(new(BasicBoard))
  newBoard.NewBoard(startSymbol)
  newBoard.SetArray(gameState)
  return newBoard
}
