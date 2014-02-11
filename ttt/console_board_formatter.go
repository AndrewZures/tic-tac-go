package ttt

type ConsoleBoardFormatter struct { }

func (c *ConsoleBoardFormatter) FormatBoard (board Board, messages Messages) (string) {
  var boardResult string

  gameState := board.State()

  for index, gameSpace := range gameState {

    boardResult += c.FormatSymbol(gameSpace, messages)

    if c.lastIndex(board, index) {
      boardResult += "\n"
    } else if c.endOfRow(board, index) {
      boardResult += c.BuildHorizontalDivider(board, messages)
    } else {
      boardResult += messages.VerticalDivider()
    }
  }

  return boardResult
}

func (c *ConsoleBoardFormatter) FormatSymbol(spotData string, messages Messages) string {
  switch spotData {
  case "": return messages.EmptySpot()
  case "x": return messages.PlayerXSymbol()
  case "o": return messages.PlayerOSymbol()
  }

  return ""
}

func (c ConsoleBoardFormatter) lastIndex(board Board, index int) (bool) {
  return index == len(board.State()) - 1
}

func (c *ConsoleBoardFormatter) endOfRow(board Board, index int) (bool) {
  return (index+1) % board.Offset() == 0
}

func (c *ConsoleBoardFormatter) BuildHorizontalDivider(board Board, messages Messages) (string) {
  divider := []byte("\n")
  adjustment := board.Offset() - 3
  dividerLength := board.Offset() * messages.SpotWidth() + adjustment

  for i := 0; i < dividerLength; i++ {
    divider = append(divider, []byte(messages.HorizontalDivider())...)
  }

  divider = append(divider, []byte("\n")...)
  return string(divider)
}
