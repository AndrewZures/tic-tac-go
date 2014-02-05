package ttt

type Board interface {
  Array() []string; //TODO can probably get rid of this
  SetArray([]string);
  NewBoard(string) bool;
  OpenSpots() []int;
  RecordMove(int, string) bool;
  RemoveMove(int);
  Status() string;
  PlayerTurn() string;
  Score([]string) string;

  //bundle below to common location
  RowWinner([]string) string;
  ColumnWinner([]string) string;
  DiagonalWinner([]string) string;
}
