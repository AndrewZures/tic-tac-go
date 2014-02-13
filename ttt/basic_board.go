package ttt

type BasicBoard struct {
	array       []string
	offset      int
	description string
}

func (b *BasicBoard) NewBoard(boardLength int, offset int, description string) bool {
	if b.array == nil {
		b.array = make([]string, boardLength, boardLength)
		b.offset = offset
		b.description = description
		return true
	} else {
		return false
	}
}

func (b BasicBoard) ValidMove(move int) bool {
	return b.validateMove(move)
}

func (b *BasicBoard) RecordMove(move int, symbol string) bool {
	if b.validateMove(move) == true {
		b.array[move] = symbol
		return true
	} else {
		return false
	}
}

func (b *BasicBoard) RemoveMove(move int) {
	b.array[move] = ""
}

func (b BasicBoard) validateMove(move int) bool {
	return b.moveIsWithinBounds(move) && b.spotIsAvailable(move)
}

func (b BasicBoard) moveIsWithinBounds(move int) bool {
	return move < len(b.array) && move > -1
}

func (b BasicBoard) spotIsAvailable(move int) bool {
	return b.array[move] == ""
}

func (b BasicBoard) OpenSpots() []int {
	gameState := b.State()
	openSpots := make([]int, 0)

	for i := 0; i < len(gameState); i++ {
		if gameState[i] == "" {
			openSpots = append(openSpots, i)
		}
	}

	return openSpots
}

func (b BasicBoard) State() []string {
	return b.copyStringArray(b.array)
}

func (b *BasicBoard) SetState(newArray []string) {
	b.array = newArray
}

func (b BasicBoard) Offset() int {
	return b.offset
}

func (b BasicBoard) Description() string {
	return b.description
}

func (b BasicBoard) copyStringArray(stringToCopy []string) []string {
	newString := make([]string, len(stringToCopy))
	for i := 0; i < len(stringToCopy); i++ {
		newString[i] = stringToCopy[i]
	}

	return newString
}
