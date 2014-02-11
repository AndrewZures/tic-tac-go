package ttt


type TTTFactory struct {
  playerList []Player;
}


func (f TTTFactory) PlayerTypes(userInterface UserInterface, rules Rules) ([]Player) {

  playerList := make([]Player, 0)
  playerList = append(playerList, f.getHumanPlayer("", "Human", userInterface))
  playerList = append(playerList, f.getComputerPlayer("", "Computer", rules))

  return playerList
}


func (f TTTFactory) Player(playerTemplate Player, userInterface UserInterface, rules Rules) (Player) {
    switch {

      case playerTemplate.Description() == "Human":
        return f.getHumanPlayer(playerTemplate.Symbol(), "Human", userInterface)

      case playerTemplate.Description() == "Computer":
        return f.getComputerPlayer(playerTemplate.Symbol(), playerTemplate.Description(), rules)
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

func (f TTTFactory) getHumanPlayer(symbol, description string, userInterface UserInterface) (Player) {
  human := new(HumanPlayer)
  human.NewHumanPlayer(symbol, description, userInterface)
  return human
}

func (f TTTFactory) getComputerPlayer(symbol, description string, rules Rules) (Player) {
        computer := new(ComputerPlayer)
        computer.NewComputerPlayer(symbol, description, rules)
        return computer
}
