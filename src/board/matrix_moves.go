package board

import (
	"github.com/ggeorgiev/instant-chess/src/peace"
	"github.com/ggeorgiev/instant-chess/src/peacemoves"
)

func (m Matrix) Moves() peacemoves.Fulls {
	var moves peacemoves.Fulls

	whiteKing := m.FindSinglePeace(peace.WhiteKing)
	blackKing := m.FindSinglePeace(peace.BlackKing)

	whiteFromTos := m.WhiteTos(whiteKing)
	for _, fromTos := range whiteFromTos {
		var answers peacemoves.Answers
		for _, to := range fromTos.Tos {
			original := m[to]
			m[to] = m[fromTos.From]
			m[fromTos.From] = peace.NoFigure

			blackFromTos := m.BlackTos(blackKing)
			if blackFromTos != nil {
				answers = append(answers, peacemoves.Answer{
					WhiteTo:      to,
					BlackAnswers: blackFromTos,
				})
			}

			m[fromTos.From] = m[to]
			m[to] = original

		}

		moves = append(moves, peacemoves.Full{
			WhiteForm: fromTos.From,
			Answers:   answers,
		})
	}

	return moves
}
