package ttt

type Game struct {
  status string;
}

//TODO complete
func setupNewGame(console ConsoleUI, factory Factory){
    playerTypes := factory.PlayerTypes()
    console.DisplayPlayerTypes(playerTypes)

//  player1Type := console.QueryUserForPlayerType(playerTypes, "Player 1")
//
//  console.DisplayPlayerTypes(playerTypes)
//  player2Type := console.QueryForPlayerType(playerTypes, "Player 2")






  //playerTypes = factory.PlayerTypes()
  //console.ShowPlayerTypes(playerTypes)
  //player1Type := console.QueryForPlayerType(playerTypes)
  //player2Type := console.QueryForPlayerType(playerTypes)
  //player1 := factory.NewPlayer(player1type)
  //player2 := factory.NewPlayer(player2type)

  //boardTypes = factory.BoardTypes()
  //console.ShowBoardTypes(boardTypes)
  //board = factory.Board(boardType)

  //game = Game{board, player1, player2}
  //game.Run()





  //playerTypes := []Player{Player(new(HumanPlayer)), Player(new(ComputerPlayer))}
  ////player1Type := console.QueryPlayerType(playerTypes)
  ////player2Type := console.QueryPlayerType(playerTypes)
  //if player1Type == player2Type {

  //}
}

func runGame(){

}
