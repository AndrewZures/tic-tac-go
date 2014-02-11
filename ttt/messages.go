package ttt

type Messages interface {
  BuildMessages()
  ChooseMovePrompt() string
  GameTieResponse() string
  GameWinnerResponse() string
  NewGamePrompt() string
  InvalidChoiceResponse() string
  BoardTypePrompt() string
  BoardTypesResponse() string
  PlayerTypePrompt() string
  PlayerTypesResponse() string
  HorizontalDivider() string
  VerticalDivider() string
  EmptySpot() string
  PlayerXSymbol() string
  PlayerOSymbol() string
  SpotWidth() int
  YesResponse() string
}
