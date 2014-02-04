package ttt

type Board interface {
  Array() []string;
  NewBoard() bool;
  RecordMove(int, Player) bool;
}
