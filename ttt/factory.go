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
        human.NewHumanPlayer("X", "Human")
        return human

      case playerTemplate.Description() == "Computer":
        computer := new(ComputerPlayer)
        computer.NewComputerPlayer("O", "Computer")
        return computer
    }

    return nil
}

func (f TTTFactory) BoardTypes() ([]Board) {
  boardList := make([]Board, 0)

  board3x3 := new(BasicBoard)
  board3x3.NewBoard("")
  boardList = append(boardList, Board(board3x3))

  return boardList
}

func (f *TTTFactory) Board(boardTemplate Board) (Board) {
  var board Board

  switch {
  case boardTemplate.Description() == "3x3 Board":
    board = new(BasicBoard)
    board.NewBoard("X")
    return board
  }

  return nil
}

