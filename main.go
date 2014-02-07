package main

import "ttt"

func main() {

  game := new(ttt.Game)
  console := ttt.UserInterface(new(ttt.ConsoleUI))
  factory := ttt.Factory(new(ttt.TTTFactory))
  game.Run(console, factory)

}

