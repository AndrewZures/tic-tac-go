package ttt

type Board interface {
  Array() []string; //TODO can probably get rid of this
  NewBoard() bool;
  OpenSpots() []int;
  RecordMove(int, string) bool;
  RemoveMove(int);
  Status() string;
  //bundle below to common location
  RowWinner() string;
  ColumnWinner() string;
  DiagonalWinner() string;
}
