package ttt

type Factory interface {
  PlayerTypes() ([]Player);
  Player(Player) (Player);
  BoardTypes() ([]Board);
  Board(Board, string) (Board);
}
