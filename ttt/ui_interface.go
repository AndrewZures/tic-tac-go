package ttt

type UserInterface interface {
  DisplayBoard();
  DisplayPlayerTypes([]Player);
  SelectPlayerChoice([]Player) (Player);
}
