package ttt

//import "fmt"

type ComputerPlayer struct {
  symbol string;
  typeTitle string;
  board Board;
}

func (h *ComputerPlayer) NewComputerPlayer(symbol string, typeTitle string){
  h.symbol = symbol
  h.typeTitle = typeTitle
}

func (h *ComputerPlayer) MakeMove(board Board) (int) {
  h.board = board
  return h.startMiniMax()

}

func (h *ComputerPlayer) startMiniMax() (int) {
  var bestMove int
  var score float64
  var bestScore float64
  bestScore = -100000000.0000
  bestMove = -1
  depth := 1

  gameState := h.board.Array()
  openMoves := h.board.OpenSpots(gameState)

  for i := 0; i < len(openMoves); i++ {
    score = h.executeMiniMax(gameState, h.Symbol(), openMoves[i], depth)

    if score > bestScore {
      bestScore = score
      bestMove = openMoves[i]
    }
  }

  return bestMove
}

func (h *ComputerPlayer) miniMax(gameState []string, symbol string, depth int) (float64) {
  var score float64
  bestScore := -10000000.0000
  openMoves := h.board.OpenSpots(gameState)

  for i := 0; i < len(openMoves); i++ {

    score = h.executeMiniMax(gameState, symbol, openMoves[i], depth)

    if score > bestScore {
      bestScore = score
    }
  }

  return bestScore
}

func (h *ComputerPlayer) executeMiniMax(gameState []string, symbol string, index int, depth int) (float64){
    var score float64

    gameState = h.recordTempMove(gameState, index, symbol)
    score = h.GetScore(gameState, symbol, depth)
    gameState = h.removeTempMove(gameState, index)

    return score
}

func (h *ComputerPlayer) recordTempMove(gameState []string, index int, symbol string) ([]string) {
  gameState[index] = symbol
  return gameState
}

func (h *ComputerPlayer) removeTempMove(gameState []string, index int) ([]string) {
  gameState[index] = ""
  return gameState
}


func (h *ComputerPlayer) GetScore(gameState []string, symbol string, depth int) (float64){
  var score float64

  gameStatus := h.board.Score(gameState)

  if gameStatus == symbol{
    score = 1.0 / float64(depth)
  } else if gameStatus == "tie"{
    score = 0
  } else {
    score = -h.miniMax(gameState, opponent(symbol), depth + 1)
  }

  return score
}

func opponent(symbol string) (string) {
  opponentSymbol := "other"
  if symbol == "o" {
    opponentSymbol = "x"
  } else if symbol == "x" {
    opponentSymbol = "o"
  }

  return opponentSymbol
}

func (h *ComputerPlayer) Symbol() (string) {
  return h.symbol
}

func (h *ComputerPlayer) Description() (string) {
  return h.typeTitle
}
