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
			answers = append(answers, peacemoves.Answer{
				WhiteTo:      to,
				BlackAnswers: blackFromTos,
			})

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

func (m Matrix) M1() bool {
	whiteKing := m.FindSinglePeace(peace.WhiteKing)
	blackKing := m.FindSinglePeace(peace.BlackKing)

	whiteFromTos := m.WhiteTos(whiteKing)
	for _, fromTos := range whiteFromTos {
		for _, to := range fromTos.Tos {
			original := m[to]
			m[to] = m[fromTos.From]
			m[fromTos.From] = peace.NoFigure

			mate := m.IsSquareUnderAttackFromWhite(blackKing) && m.BlackTos(blackKing) == nil

			m[fromTos.From] = m[to]
			m[to] = original

			if mate {
				return true
			}
		}
	}

	return false
}

func (m Matrix) M1s() int {
	whiteKing := m.FindSinglePeace(peace.WhiteKing)
	blackKing := m.FindSinglePeace(peace.BlackKing)

	mates := 0

	whiteFromTos := m.WhiteTos(whiteKing)
	for _, fromTos := range whiteFromTos {
		for _, to := range fromTos.Tos {
			original := m[to]
			m[to] = m[fromTos.From]
			m[fromTos.From] = peace.NoFigure

			if m.IsSquareUnderAttackFromWhite(blackKing) && m.BlackTos(blackKing) == nil {
				mates++
			}

			m[fromTos.From] = m[to]
			m[to] = original
		}
	}

	return mates
}
