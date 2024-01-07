package chess

import (
	"github.com/ggeorgiev/instant-chess/src/board"
	"github.com/ggeorgiev/instant-chess/src/peace"
	"github.com/ggeorgiev/instant-chess/src/square"
)

type Board [square.Number]peace.Figure

func (brd Board) Moves() Moves {
	var moves Moves

	wks := board.Matrix(brd).FindSinglePeace(peace.WhiteKing)
	bks := board.Matrix(brd).FindSinglePeace(peace.BlackKing)

	for s := square.ZeroIndex; s <= square.LastIndex; s++ {
		tos := board.Matrix(brd).SquareWhiteTos(s, wks)
		if len(tos) == 0 {
			continue
		}

		var toAnswers []ToAnswer
		for _, to := range tos {
			original := brd[to]
			brd[to] = brd[s]
			brd[s] = peace.NoFigure

			var answers []Answer

			for s := square.ZeroIndex; s <= square.LastIndex; s++ {
				if !brd[s].IsBlack() {
					continue
				}
				blackTos := board.Matrix(brd).SquareBlackTos(s, bks)
				if len(blackTos) > 0 {
					answers = append(answers, Answer{
						BlackFrom: s,
						BlackTos:  blackTos,
					})
				}
			}
			toAnswers = append(toAnswers, ToAnswer{
				WhiteTo: to,
				Answers: answers,
			})

			brd[s] = brd[to]
			brd[to] = original
		}

		moves = append(moves, Move{WhiteForm: s, ToAnswers: toAnswers})
	}

	return moves
}
