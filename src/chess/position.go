package chess

type Position struct {
	HorizontalySymetric bool
	VerticalySymetric   bool

	Result Result

	Board Board

	WhitePeaces int8
	BlackPeaces int8

	Valid bool

	Moves []Move
}

func CreatePosition(board Board) *Position {
	position := &Position{
		HorizontalySymetric: true,
		VerticalySymetric:   true,

		Result: UnknownResult,

		Board: board,

		WhitePeaces: 0,
		BlackPeaces: 0,

		Valid: true,
	}

	whiteKings := 0
	blackKings := 0
	whitePawns := 0
	blackPawns := 0

	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			peace := board[x][y]
			if peace.Type == Empty {
				continue
			}

			if peace.Color == PeaceColorWhite {
				position.Moves = append(position.Moves, board.Move(x, y))
				position.WhitePeaces++
			} else {
				position.BlackPeaces++
			}

			if peace.Type == King {
				if peace.Color == PeaceColorWhite {
					whiteKings++
				} else {
					if board.SquareUnderAttack(x, y, PeaceColorWhite) {
						position.Valid = false
					}

					blackKings++
				}
			} else if peace.Type == Pawn {
				if peace.Color == PeaceColorWhite {
					whitePawns++
				} else {
					blackPawns++
				}
			}
		}
	}

	position.Valid = position.Valid &&
		(whiteKings == 1 && blackKings == 1) &&
		(whitePawns <= 8 && blackPawns <= 8)

	return position
}

func ParsePosition(text string) *Position {
	board := ParseBoard(text)
	return CreatePosition(board)
}

func (p *Position) Print() {
	p.Board.Print()
}
