package ttt

type MessagesInterface interface {
  BuildMessages()
  ChooseMovePrompt() string
  GameTieResponse() string
  GameWinnerResponse() string
  NewGamePrompt() string
  InvalidChoiceResponse() string
  BoardTypePrompt() string
  PlayerTypePrompt() string
  PlayerTypesResponse() string
}
