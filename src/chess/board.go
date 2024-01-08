package chess

import (
	"github.com/ggeorgiev/instant-chess/src/board"
	"github.com/ggeorgiev/instant-chess/src/peace"
	"github.com/ggeorgiev/instant-chess/src/square"
)

type Board [square.Number]peace.Figure

func (brd Board) Moves() Moves {
	var moves Moves

	whiteKing := board.Matrix(brd).FindSinglePeace(peace.WhiteKing)
	blackKing := board.Matrix(brd).FindSinglePeace(peace.BlackKing)

	whiteFromTos := board.Matrix(brd).WhiteTos(whiteKing)
	for _, fromTos := range whiteFromTos {
		var toAnswers []ToAnswer
		for _, to := range fromTos.Tos {
			original := brd[to]
			brd[to] = brd[fromTos.From]
			brd[fromTos.From] = peace.NoFigure

			blackFromTos := board.Matrix(brd).BlackTos(blackKing)
			if blackFromTos != nil {
				toAnswers = append(toAnswers, ToAnswer{
					WhiteTo:      to,
					BlackAnswers: blackFromTos,
				})
			}

			brd[fromTos.From] = brd[to]
			brd[to] = original

		}

		moves = append(moves, Move{WhiteForm: fromTos.From, ToAnswers: toAnswers})
	}

	return moves
}
