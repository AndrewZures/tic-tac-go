package ttt

type Board interface {
  Array() []string; //TODO can probably get rid of this
  Description() string;
  SetArray([]string);
  Offset() int;
  NewBoard(string) bool;
  RecordMove(int, string) bool;
  RemoveMove(int);
  Status() string;
  GameOver() bool;
  Winner() string;
  PlayerTurn() string;
  Score([]string) string;

  //bundle below to common location
  OpenSpots([]string) []int;
  RowWinner([]string) string;
  ColumnWinner([]string) string;
  DiagonalWinner([]string) string;
}
