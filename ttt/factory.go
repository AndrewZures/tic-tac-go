package ttt

type Factory interface {
  PlayerTypes(UserInterface) ([]Player);
  Player(Player, UserInterface) (Player);
  BoardTypes() ([]Board);
  Board(Board) (Board);
}
