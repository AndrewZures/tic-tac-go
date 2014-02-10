package ttt

type BoardFormatterInterface interface {
  FormatBoard(Board, MessagesInterface) string
}
