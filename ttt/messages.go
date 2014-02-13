package ttt

type Messages interface {
	BuildMessages()

	IntroMessage() string
	ExitMessage() string

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

	PlayerSymbol(string) string
}
