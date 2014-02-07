package ttt

type UserInterface interface {
  DisplayBoard(Board);
  DisplayPlayerTypes([]Player);
  SelectPlayerChoice([]Player, string) (Player);
}
