package ttt

type Factory interface {
  PlayerTypes() ([]Player);
  Player(Player) (Player);
  BoardTypes() ([]Board);
}
