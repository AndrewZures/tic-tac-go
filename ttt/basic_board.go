package ttt

//import "fmt"

type BasicBoard struct {
    array []string;
    offset int;
    playerTurn string;
}

func (b *BasicBoard) NewBoard(startSymbol string) (bool) {
  if b.array == nil {
    b.array = make([]string, 9, 9)
    b.offset = 3
    b.playerTurn = startSymbol
    return true
  } else {
    return false
  }
}

func (b *BasicBoard) PlayerTurn() (string) {
  return b.playerTurn
}

func (b *BasicBoard) Array() ([]string) {
  return b.array
}

func (b *BasicBoard) SetArray(newArray []string ) {
  b.array = newArray
}

func (b *BasicBoard) OpenSpots(gameState []string) ([]int) {
  openSpots := make([]int, 0)

  for i := 0; i < len(gameState); i++ {
    if gameState[i] == "" {
      openSpots = append(openSpots, i)
    }
  }

  return openSpots
}

func (b *BasicBoard) RecordMove(move int, symbol string) (bool) {
  if b.validateMove(move, symbol) == true {
    b.array[move] = symbol
    b.toggleTurn()
    return true
  } else {
    return false
  }
}

func (b *BasicBoard) Status() (string) {
  return b.Score(b.array)
}

func (b *BasicBoard) toggleTurn() {
  if b.playerTurn == "x" {
    b.playerTurn = "o"
  } else if b.playerTurn == "o" {
    b.playerTurn = "x"
  } else {
    b.playerTurn = "z"
  }
}

func (b *BasicBoard) RemoveMove(move int) {
  b.array[move] = ""
}

func (b *BasicBoard) validateMove(move int, symbol string) (bool) {
  return b.moveIsWithinBounds(move) && b.spotIsAvailable(move) && b.isPlayerTurn(symbol)
}

func (b *BasicBoard) isPlayerTurn (symbol string) (bool) {
  return b.PlayerTurn() == symbol
}

func (b *BasicBoard) moveIsWithinBounds (move int) (bool) {
  return move < len(b.array) && move > -1
}

func (b *BasicBoard) spotIsAvailable (move int) (bool) {
  return b.array[move] == ""
}

/////

func (b *BasicBoard) Score(gameState []string) (string) {
  rowStatus := b.RowWinner(gameState)
  if  rowStatus != "" {
    return rowStatus
  }

  columnStatus := b.ColumnWinner(gameState)
  if  columnStatus != "" {
    return columnStatus
  }

  diagonalStatus := b.DiagonalWinner(gameState)
  if  diagonalStatus != "" {
    return diagonalStatus
  }

  if len(b.OpenSpots(gameState)) == 0 {
    return "tie"
  } else {
    return "inprogress"
  }

}

func (b *BasicBoard) RowWinner(gameState []string) (string) {
  var row []string

  for i := 0; i < len(b.array); i += b.offset {
    row = b.rowElements(gameState, i)
    if b.AllSameSymbols(row) == true {
      return gameState[i]
    }
  }
  return ""
}

func (b *BasicBoard) rowElements(gameState []string, startIndex int) ([]string) {
  return gameState[startIndex:startIndex+b.offset]
}

func (b *BasicBoard) ColumnWinner(gameState []string) (string) {
  var column []string

  for i := 0; i < b.offset; i++ {
    column = b.columnElements(gameState, i)
    if b.AllSameSymbols(column) == true {
      return gameState[i]
    }
  }
  return ""

}

func (b *BasicBoard) columnElements(gameState []string, startIndex int) ([]string) {
  elements := make([]string, 0)

  for i := startIndex; i < len(gameState); i += b.offset {
    elements = append(elements,gameState[i])
  }
  return elements

}

func (b *BasicBoard) DiagonalWinner(gameState []string) (string) {
    if b.AllSameSymbols(b.LRDiagonalElements(gameState)){
      return gameState[0]
    }

    if b.AllSameSymbols(b.RLDiagonalElements(gameState)){
      return gameState[b.offset - 1]
    }

  return ""
}

func (b *BasicBoard) LRDiagonalElements(gameState []string) ([]string) {
  elements := make([]string, 0)

  for i := 0; i < len(gameState); i += b.offset+1 {
    elements = append(elements, gameState[i])
  }
  return elements

}

//TODO refactor
func (b *BasicBoard) RLDiagonalElements(gameState []string) ([]string) {
  elements := make([]string, 0)
  newOffset := b.offset - 1

  for i := newOffset; i < len(gameState)-1; i += newOffset {
    elements = append(elements,gameState[i])
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

