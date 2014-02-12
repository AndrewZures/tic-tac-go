package ttt

type Messages interface {
  BuildMessages()

  ChooseMovePrompt() string
  BoardTypePrompt() string
  PlayerTypePrompt() string
  NewGamePrompt() string

  GameTieResponse() string
  GameWinnerResponse() string
  InvalidChoiceResponse() string
  BoardTypesResponse() string
  PlayerTypesResponse() string

  HorizontalDivider() string
  VerticalDivider() string
  EmptySpot() string
  PlayerXSymbol() string
  PlayerOSymbol() string
  SpotWidth() int

  YesOption() string

  WinnerSymbol(string) string
}
