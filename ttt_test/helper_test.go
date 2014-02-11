package ttt_test

import (
  . "github.com/andrewzures/tictactoe/ttt"
  "fmt"
  "bytes"
)


func GenerateBoard(boardSize int, offset int, startSymbol string, description string, gameState []string) (Board) {
  newBoard := Board(new(BasicBoard))
  newBoard.NewBoard(boardSize, offset, startSymbol, description)
  newBoard.SetState(gameState)
  return newBoard
}

func GenerateEmpty3x3Board(startSymbol string) (Board) {
  gameState := []string{"","","","","","","","",""}
  return Generate3x3Board(startSymbol, gameState)
}

func Generate3x3Board(startSymbol string, gameState []string) (Board) {
  return GenerateBoard(9, 3, startSymbol, "3x3 Board", gameState)
}

func SetMockInput(reader *bytes.Buffer, input string){
    fmt.Fprintf(reader, input)
}

func BuildConsoleUI(inOut InOut) ConsoleUI {
  messages := buildConsoleMessages()
  boardFormatter := BoardFormatter(new(ConsoleBoardFormatter))
  console := ConsoleUI{inOut, messages, boardFormatter}
  return console
}

func BuildConsoleMessages() Messages {
  consoleMessages := new(ConsoleMessages)
  consoleMessages.BuildMessages()
  messages := Messages(consoleMessages)
  return messages
}
