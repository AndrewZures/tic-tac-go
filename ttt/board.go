package ttt

type Board interface {

  State() []string;
  Description() string;
  SetArray([]string);
  Offset() int;
  NewBoard(int, int, string, string) bool;
  RecordMove(int, string) bool;
  RemoveMove(int);
  OpenSpots() []int;

}
