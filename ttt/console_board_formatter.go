package ttt

type ConsoleBoardFormatter struct {
}

func (c *ConsoleBoardFormatter) FormatBoard (board Board, messages Messages) (string) {
  var boardResult string

  gameState := board.Array()

  for i := 0; i < len(gameState); i++ {

    boardResult += c.FormatSymbol(gameState[i], messages)

    if c.LastIndex(board, i) {
      boardResult += "\n"
    } else if c.EndOfRow(board, i) {
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

func (c *ConsoleBoardFormatter) LastIndex(board Board, index int) (bool) {
  return index == len(board.Array()) - 1
}

func (c *ConsoleBoardFormatter) PrintSymbol(symbol string) (string) {
  return symbol
}

func (c *ConsoleBoardFormatter) EndOfRow(board Board, index int) (bool) {
  return (index+1) % board.Offset() == 0
}

func (c *ConsoleBoardFormatter) BuildHorizontalDivider(board Board, messages Messages) (string) {
  divider := []byte("\n")
  dividerLength := board.Offset() * messages.SpotWidth()

  for i := 0; i < dividerLength; i++ {
    divider = append(divider, []byte(messages.HorizontalDivider())...)
  }

  divider = append(divider, []byte("\n")...)
  return string(divider)
}
