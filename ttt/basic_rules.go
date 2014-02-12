package ttt

type BasicRules struct{}

func (r BasicRules) Score(board Board) string {
	numOpenSpots := len(board.OpenSpots())
	segments := r.buildSegments(board)
	segmentsStatus := r.scoreSegments(segments)
	gameStatus := r.gameStatus(segmentsStatus, numOpenSpots)
	return gameStatus
}

func (r BasicRules) buildSegments(board Board) [][][]string {
	segments := [][][]string{r.rows(board),
		r.columns(board),
		r.leftDiagonal(board),
		r.rightDiagonal(board)}

	return segments
}

func (r BasicRules) scoreSegments(segments [][][]string) string {

	for _, segment := range segments {
		segmentStatus := r.scoreSegment(segment)
		if segmentStatus != "" {
			return segmentStatus
		}
	}

	return ""
}

func (r BasicRules) scoreSegment(segment [][]string) string {

	for _, subSegment := range segment {
		if r.AllSameSymbol(subSegment) {
			return subSegment[0]
		}
	}
	return ""
}

func (r BasicRules) gameStatus(segmentsStatus string, openSpots int) string {

	if segmentsStatus != "" {
		return segmentsStatus
	} else if openSpots == 0 {
		return "tie"
	} else {
		return "inprogress"
	}

}

func (r *BasicRules) AllSameSymbol(data []string) bool {
	if len(data) == 0 || data[0] == "" {
		return false
	}

	for i := 0; i < len(data); i++ {
		if data[0] != data[i] {
			return false
		}
	}

	return true
}

func (r BasicRules) rows(board Board) [][]string {
	offset := board.Offset()
	rows := make([][]string, offset)

	for i := 0; i < offset; i++ {
		startIndex := i * offset
		rows[i] = r.row(board, startIndex)
	}

	return rows
}

func (r BasicRules) row(board Board, startIndex int) []string {
	gameState, offset := r.GetGameStateAndOffset(board)
	return gameState[startIndex : startIndex+offset]
}

func (r BasicRules) columns(board Board) [][]string {
	offset := board.Offset()
	columns := make([][]string, offset)

	for i := 0; i < offset; i++ {
		columns[i] = r.column(board, i)
	}

	return columns
}

func (r *BasicRules) column(board Board, startIndex int) []string {
	elements := make([]string, 0)
	gameState, offset := r.GetGameStateAndOffset(board)

	for i := startIndex; i < len(gameState); i += offset {
		elements = append(elements, gameState[i])
	}

	return elements
}

func (r *BasicRules) leftDiagonal(board Board) [][]string {
	elements := make([]string, 0)
	gameState, offset := r.GetGameStateAndOffset(board)
	diagonal := make([][]string, offset)

	for i := 0; i < len(gameState); i += offset + 1 {
		elements = append(elements, gameState[i])
	}
	diagonal[0] = elements

	return diagonal
}

func (r *BasicRules) rightDiagonal(board Board) [][]string {
	elements := make([]string, 0)
	gameState, offset := r.GetGameStateAndOffset(board)
	diagonal := make([][]string, offset)
	newOffset := offset - 1

	for i := newOffset; i < len(gameState)-1; i += newOffset {
		elements = append(elements, gameState[i])
	}
	diagonal[0] = elements

	return diagonal
}

func (r BasicRules) IsWinner(board Board) bool {
	if r.Score(board) != "inprogress" {
		return true
	}
	return false
}

func (r BasicRules) Winner(board Board) string {
	return r.Score(board)
}

func (r BasicRules) GameOver(board Board) bool {
	return r.Score(board) != "inprogress"
}

func (r BasicRules) GetGameStateAndOffset(board Board) ([]string, int) {
	gameState := board.State()
	offset := board.Offset()
	return gameState, offset
}
