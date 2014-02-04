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

func (b *BasicBoard) validateMove(move int) (bool) {
  return b.moveIsWithinBounds(move) && b.arraySpotAvailable(move)
}

func (b *BasicBoard) moveIsWithinBounds (move int) (bool) {
  return move < len(b.array) && move > -1
}

func (b *BasicBoard) arraySpotAvailable (move int) (bool) {
  return b.array[move] == ""
}

func (b *BasicBoard) RecordMove(move int, player Player) (bool) {
  if b.validateMove(move) == true {
    b.array[move] = player.Symbol()
    return true
  } else {
    return false
  }

}
