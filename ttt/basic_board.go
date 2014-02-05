package ttt

//import "fmt"

type BasicBoard struct {
    array []string;
    offset int;
}

func (b *BasicBoard) NewBoard() (bool) {
  if b.array == nil {
    b.array = make([]string, 9, 9)
    b.offset = 3
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
  rowStatus := b.RowWinner()
  if  rowStatus != "" {
    return rowStatus
  }

  columnStatus := b.ColumnWinner()
  if  columnStatus != "" {
    return columnStatus
  }

  diagonalStatus := b.DiagonalWinner()
  if  diagonalStatus != "" {
    return diagonalStatus
  }

  if len(b.OpenSpots()) == 0 {
    return "tie"
  } else {
    return "inprogress"
  }

}

func (b *BasicBoard) RowWinner() (string) {
  var row []string

  for i := 0; i < len(b.array); i += b.offset {
    row = b.rowElements(i)
    if b.AllSameSymbols(row) == true {
      return b.array[i]
    }
  }
  return ""
}

func (b *BasicBoard) rowElements(startIndex int) ([]string) {
  return b.array[startIndex:startIndex+b.offset]
}

func (b *BasicBoard) ColumnWinner() (string) {
  var column []string

  for i := 0; i < b.offset; i++ {
    column = b.columnElements(i)
    if b.AllSameSymbols(column) == true {
      return b.array[i]
    }
  }
  return ""

}

func (b *BasicBoard) columnElements(startIndex int) ([]string) {
  elements := make([]string, 0)

  for i := startIndex; i < len(b.array); i += b.offset {
    elements = append(elements,b.array[i])
  }
  return elements

}

func (b *BasicBoard) DiagonalWinner() (string) {
    if b.AllSameSymbols(b.LRDiagonalElements()){
      return b.array[0]
    }

    if b.AllSameSymbols(b.RLDiagonalElements()){
      return b.array[b.offset - 1]
    }

  return ""
}

func (b *BasicBoard) LRDiagonalElements() ([]string) {
  elements := make([]string, 0)

  for i := 0; i < len(b.array); i += b.offset+1 {
    elements = append(elements,b.array[i])
  }
  return elements

}

//TODO refactor
func (b *BasicBoard) RLDiagonalElements() ([]string) {
  elements := make([]string, 0)
  newOffset := b.offset - 1

  for i := newOffset; i < len(b.array)-1; i += newOffset {
    elements = append(elements,b.array[i])
  }
  return elements

}


//TODO refactor
func (b *BasicBoard) AllSameSymbols (data []string) (bool) {
  if data[0] == "" {
    return false
  }

  for i := 1; i < len(data); i++ {
      if data[0] != data[i] {
          return false
      }
  }

  return true
}

