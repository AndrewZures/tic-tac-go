package main

import "github.com/andrewzures/tictactoe/ttt"
import "os"

func main() {

  game := new(ttt.TTTGame)
  inOut := ttt.InOutInterface(ttt.ConsoleIO{os.Stdout, os.Stdin})
  consoleMessage := new(ttt.ConsoleMessages)
  messages := ttt.MessagesInterface(consoleMessage)
  messages.BuildMessages()
  boardFormatter := ttt.BoardFormatterInterface(new(ttt.ConsoleBoardFormatter))
  consoleui := ttt.ConsoleUI{inOut, messages, boardFormatter}
  console := ttt.UserInterface(consoleui)
  factory := ttt.Factory(new(ttt.TTTFactory))
  game.Run(console, factory)

}

