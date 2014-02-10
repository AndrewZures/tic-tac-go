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
  OpenSpots([]string) []int;

  //rules
  Winner() string;
  Score([]string) string;
  RowWinner([]string) string;
  ColumnWinner([]string) string;
  DiagonalWinner([]string) string;
  PlayerTurn() string;
  Status() string;

  //game status
  GameOver() bool;
}
