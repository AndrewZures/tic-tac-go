package ttt

type TTTFactory struct {
  playerList []Player;
}

func (f *TTTFactory) PlayerTypes() ([]Player) {
  playerList := make([]Player, 0)

  human := new(HumanPlayer)
  human.NewHumanPlayer("", "Human")
  playerList = append(playerList, Player(human))

  computer := new(ComputerPlayer)
  computer.NewComputerPlayer("", "Computer")
  playerList = append(playerList, Player(computer))

  return playerList
}

func (f *TTTFactory) Player(playerTemplate Player) (Player) {
    switch {

      case playerTemplate.Description() == "Human":
        human := new(HumanPlayer)
        human.NewHumanPlayer(playerTemplate.Symbol(), "Human")
        return human

      case playerTemplate.Description() == "Computer":
        computer := new(ComputerPlayer)
        computer.NewComputerPlayer(playerTemplate.Symbol(), "Computer")
        return computer
    }

    return nil
}
