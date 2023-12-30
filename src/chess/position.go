package chess

type Position struct {
	HorizontalySymetric bool
	VerticalySymetric   bool

	Result Result

	Board Board

	WhitePeaces int8
	BlackPeaces int8

	Valid bool

	Moves Moves
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

	for s := Square(0); s < 64; s++ {
		peace := board[s]
		if peace.IsEmpty() {
			continue
		}

		if peace.IsWhite() {
			position.Moves = append(position.Moves, board.Move(s))
			position.WhitePeaces++
		} else {
			position.BlackPeaces++
		}

		if peace.IsKing() {
			if peace.IsWhite() {
				whiteKings++
			} else {
				if board.SquareUnderAttack(s, PeaceColorWhite) {
					position.Valid = false
				}

				blackKings++
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

func (p *Position) M1() bool {
	for _, move := range p.Moves {
		for _, toAnswer := range move.ToAnswers {
			if len(toAnswer.Answers) == 0 {
				board := p.Board
				original := board[toAnswer.WhiteTo]
				board[toAnswer.WhiteTo] = board[move.WhiteForm]
				board[move.WhiteForm] = Empty

				bk := board.FindPeace(BlackKing)
				mate := board.SquareUnderAttack(bk, PeaceColorWhite)

				board[move.WhiteForm] = board[toAnswer.WhiteTo]
				board[toAnswer.WhiteTo] = original
				if mate {
					return true
				}
			}
		}
	}
	return false
}
