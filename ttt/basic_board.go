package ttt

type BasicBoard struct {
    array []string;
}

func (b *BasicBoard) NewBoard() (bool) {
  if b.array == nil {
    b.array = make([]string, 9, 9)
    return true
  } else {
    return false
  }
}

func (b *BasicBoard) Array() ([]string) {
  return b.array
}

func (b *BasicBoard) OpenSpots() ([]int) {
  openSpots := make([]int, 0)

  for i := 0; i < len(b.array); i++ {
    if b.array[i] == "" {
      openSpots = append(openSpots, i)
    }
  }

  return openSpots
}

func (b *BasicBoard) RecordMove(move int, symbol string) (bool) {
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

func (b *BasicBoard) validateMove(move int) (bool) {
  return b.moveIsWithinBounds(move) && b.spotIsAvailable(move)
}

func (b *BasicBoard) moveIsWithinBounds (move int) (bool) {
  return move < len(b.array) && move > -1
}

func (b *BasicBoard) spotIsAvailable (move int) (bool) {
  return b.array[move] == ""
}

/////

func (b *BasicBoard) Status() (string) {
  return ""
}

