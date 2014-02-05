package ttt_test

import (
  . "github.com/andrewzures/tictactoe/ttt"
)

func getPlayers() (Player, Player) {
  human1 := new(HumanPlayer)
  human1.SetSymbol("x")

  human2 := new(HumanPlayer)
  human2.SetSymbol("o")

  return Player(human1), Player(human2)
}

func GenerateBoard(startSymbol string, gameState []string) (Board) {
  newBoard := Board(new(BasicBoard))
  newBoard.NewBoard(startSymbol)
  newBoard.SetArray(gameState)
  return newBoard
}
