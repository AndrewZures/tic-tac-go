package ttt

type Board interface {

  //board
  Array() []string; //TODO can probably get rid of this
  Description() string;
  SetArray([]string);
  Offset() int;
  NewBoard(int, int, string, string) bool;
  RecordMove(int, string) bool;
  RemoveMove(int);
  OpenSpots() []int;

}
