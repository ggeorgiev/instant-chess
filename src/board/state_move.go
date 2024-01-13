package board

import (
	"github.com/ggeorgiev/instant-chess/src/move"
	"github.com/ggeorgiev/instant-chess/src/peace"
	"github.com/ggeorgiev/instant-chess/src/peacemoves"
	"github.com/ggeorgiev/instant-chess/src/square"
)

type Snapshot struct {
	Captured      peace.Figure
	Castling      move.CastlingRights
	EnPassantFile square.File
}

func (s *State) Moves() peacemoves.Fulls {
	var moves peacemoves.Fulls

	whiteKing := s.Matrix.FindSinglePeace(peace.WhiteKing)
	blackKing := s.Matrix.FindSinglePeace(peace.BlackKing)

	whiteFromTos := s.Matrix.WhiteTos(whiteKing)
	for _, fromTos := range whiteFromTos {
		var answers peacemoves.Answers
		for _, to := range fromTos.Tos {
			snapshot := s.DoWhite(fromTos.From, to)

			blackFromTos := s.Matrix.BlackTos(blackKing)
			answers = append(answers, peacemoves.Answer{
				WhiteTo:      to,
				BlackAnswers: blackFromTos,
			})

			s.UndoWhite(snapshot, fromTos.From, to)
		}

		moves = append(moves, peacemoves.Full{
			WhiteForm: fromTos.From,
			Answers:   answers,
		})
	}

	return moves
}

func (s *State) M1() bool {
	whiteKing := s.Matrix.FindSinglePeace(peace.WhiteKing)
	blackKing := s.Matrix.FindSinglePeace(peace.BlackKing)

	whiteFromTos := s.Matrix.WhiteTos(whiteKing)
	for _, fromTos := range whiteFromTos {
		for _, to := range fromTos.Tos {
			snapshot := s.DoWhite(fromTos.From, to)

			mate := s.Matrix.IsSquareUnderAttackFromWhite(blackKing) && s.Matrix.BlackTos(blackKing) == nil

			s.UndoWhite(snapshot, fromTos.From, to)

			if mate {
				return true
			}
		}
	}

	return false
}

func (s *State) M1s() int {
	whiteKing := s.Matrix.FindSinglePeace(peace.WhiteKing)
	blackKing := s.Matrix.FindSinglePeace(peace.BlackKing)

	mates := 0

	whiteFromTos := s.Matrix.WhiteTos(whiteKing)
	for _, fromTos := range whiteFromTos {
		for _, to := range fromTos.Tos {
			snapshot := s.DoWhite(fromTos.From, to)

			if s.Matrix.IsSquareUnderAttackFromWhite(blackKing) && s.Matrix.BlackTos(blackKing) == nil {
				mates++
			}

			s.UndoWhite(snapshot, fromTos.From, to)
		}
	}

	return mates
}
