package ttt

type Board interface {

  State() []string;
  Description() string;
  SetState([]string);
  Offset() int;
  NewBoard(int, int, string) bool;
  RecordMove(int, string) bool;
  RemoveMove(int);
  OpenSpots() []int;
}
