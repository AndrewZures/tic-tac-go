package ttt

type Game interface {
  Run(UserInterface, Factory);
  SetupNewGame(UserInterface, Factory) (Board,Player,Player,Rules);
}
