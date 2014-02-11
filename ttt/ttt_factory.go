package ttt


type TTTFactory struct {
  playerList []Player;
}


func (f TTTFactory) PlayerTypes(userInterface UserInterface) ([]Player) {
  playerList := make([]Player, 0)

  human := new(HumanPlayer)
  human.NewHumanPlayer("", "Human", userInterface)
  playerList = append(playerList, Player(human))

  computer := new(ComputerPlayer)
  rules := Rules(new(BasicRules))
  computer.NewComputerPlayer("", "Computer", rules)
  playerList = append(playerList, Player(computer))

  return playerList
}

func (f TTTFactory) Player(playerTemplate Player, userInterface UserInterface) (Player) {
    switch {

      case playerTemplate.Description() == "Human":
        human := new(HumanPlayer)
        human.NewHumanPlayer(playerTemplate.Symbol(), "Human", userInterface)
        return human

      case playerTemplate.Description() == "Computer":
        computer := new(ComputerPlayer)
        rules := Rules(new(BasicRules))
        computer.NewComputerPlayer(playerTemplate.Symbol(), "Computer", rules)
        return computer
    }

    return nil
}

func (f TTTFactory) boardMaps() (map[string]map[string]int) {
  boardMap := map[string]map[string]int{ "3x3 Board": {"size": 9, "offset": 3},
                                         "4x4 Board": {"size": 16, "offset": 4},
                                         "5x5 Board": {"size": 25, "offset": 5}}
  return boardMap
}

func (f TTTFactory) BoardTypes() ([]Board) {
  boardList := make([]Board, 0)

  for description, structure := range f.boardMaps() {
      boardList = append(boardList, f.generateBoard(description, structure))
  }

  return boardList
}

func (f TTTFactory) Board(boardTemplate Board) (Board) {
  boardMaps := f.boardMaps()
  description := boardTemplate.Description()

  return f.generateBoard(description, boardMaps[description])
}

func (f TTTFactory) generateBoard(boardName string, boardMap map[string]int) (Board) {
  board := new(BasicBoard)
  board.NewBoard(boardMap["size"], boardMap["offset"], boardName)
  return board
}
