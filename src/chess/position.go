package chess

import (
	"github.com/ggeorgiev/instant-chess/src/peace"
	"github.com/ggeorgiev/instant-chess/src/square"
)

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

	for s := square.ZeroIndex; s <= square.LastIndex; s++ {
		peace := board[s]
		if peace.IsNoFigure() {
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
				if board.SquareUnderAttackFromWhite(s) {
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
				board[move.WhiteForm] = peace.NoFigure

				bk := board.FindPeace(peace.BlackKing)
				mate := board.SquareUnderAttackFromWhite(bk)

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
