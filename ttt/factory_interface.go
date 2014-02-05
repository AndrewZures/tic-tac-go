package ttt

type Factory interface {
  PlayerTypes() ([]Player);
}
