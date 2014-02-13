package ttt

type Board interface {
	NewBoard(int, int, string) bool
	State() []string
	SetState([]string)
	Description() string
	Offset() int
	RecordMove(int, string) bool
	RemoveMove(int)
	OpenSpots() []int
}
