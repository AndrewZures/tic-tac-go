package ttt_test

import (
  . "github.com/andrewzures/tictactoe/ttt"
  "fmt"
  "bytes"
)


func GenerateBoard(boardSize int, offset int, description string, gameState []string) (Board) {
  newBoard := Board(new(BasicBoard))
  newBoard.NewBoard(boardSize, offset, description)
  newBoard.SetState(gameState)
  return newBoard
}

func GenerateEmpty3x3Board() (Board) {
  gameState := []string{"","","","","","","","",""}
  return Generate3x3Board(gameState)
}

func Generate3x3Board(gameState []string) (Board) {
  return GenerateBoard(9, 3,"3x3 Board", gameState)
}

func GenerateEmpty4x4Board() (Board) {
  gameState := make([]string, 16, 16)
  return Generate3x3Board(gameState)
}

func Generate4x4Board(gameState []string) (Board) {
  return GenerateBoard(16, 4,"4x4 Board", gameState)
}

func SetMockInput(reader *bytes.Buffer, input string){
    fmt.Fprintf(reader, input)
}

func BuildConsoleUI(inOut InOut) ConsoleUI {
  messages := BuildConsoleMessages()
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
