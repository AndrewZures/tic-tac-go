package main

import "github.com/andrewzures/tictactoe/ttt"
import "os"

func main() {

  game := new(ttt.TTTGame)
  consoleui := ttt.ConsoleUI{os.Stdout, os.Stdin}
  console := ttt.UserInterface(consoleui)
  factory := ttt.Factory(new(ttt.TTTFactory))
  game.Run(console, factory)

}

