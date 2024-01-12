package board

import (
	"github.com/ggeorgiev/instant-chess/src/peace"
	"github.com/ggeorgiev/instant-chess/src/peacemoves"
	"github.com/ggeorgiev/instant-chess/src/peaceplaces"
	"github.com/ggeorgiev/instant-chess/src/square"
)

type Snapshot struct {
	Captured peace.Figure
}

func (s *State) Do(from square.Index, to square.Index) Snapshot {
	if from == peaceplaces.WhiteKingStartingPlace && s.Matrix[from] == peace.WhiteKing {
		if to == peaceplaces.WhiteKingKingsideCastlingPlace {
			s.Matrix[peaceplaces.WhiteRookKingsideStartingPlace] = peace.NoFigure
			s.Matrix[peaceplaces.WhiteRookKingsideCastlingPlace] = peace.WhiteRook
		} else if to == peaceplaces.WhiteKingQueensideCastlingPlace {
			s.Matrix[peaceplaces.WhiteRookQueensideStartingPlace] = peace.NoFigure
			s.Matrix[peaceplaces.WhiteRookQueensideCastlingPlace] = peace.WhiteRook
		}
	}

	if from == peaceplaces.BlackKingStartingPlace && s.Matrix[from] == peace.BlackKing {
		if to == peaceplaces.BlackKingKingsideCastlingPlace {
			s.Matrix[peaceplaces.BlackRookKingsideStartingPlace] = peace.NoFigure
			s.Matrix[peaceplaces.BlackRookKingsideCastlingPlace] = peace.BlackRook
		} else if to == peaceplaces.BlackKingQueensideCastlingPlace {
			s.Matrix[peaceplaces.BlackRookQueensideStartingPlace] = peace.NoFigure
			s.Matrix[peaceplaces.BlackRookQueensideCastlingPlace] = peace.BlackRook
		}
	}

	snapshot := Snapshot{
		Captured: s.Matrix[to],
	}
	s.Matrix[to] = s.Matrix[from]
	s.Matrix[from] = peace.NoFigure

	return snapshot
}

func (s *State) Undo(snapshot Snapshot, from square.Index, to square.Index) {
	if from == peaceplaces.WhiteKingStartingPlace && s.Matrix[to] == peace.WhiteKing {
		if to == peaceplaces.WhiteKingKingsideCastlingPlace {
			s.Matrix[peaceplaces.WhiteRookKingsideStartingPlace] = peace.WhiteRook
			s.Matrix[peaceplaces.WhiteRookKingsideCastlingPlace] = peace.NoFigure
		} else if to == peaceplaces.WhiteKingQueensideCastlingPlace {
			s.Matrix[peaceplaces.WhiteRookQueensideStartingPlace] = peace.WhiteRook
			s.Matrix[peaceplaces.WhiteRookQueensideCastlingPlace] = peace.NoFigure
		}
	}

	if from == peaceplaces.BlackKingStartingPlace && s.Matrix[to] == peace.BlackKing {
		if to == peaceplaces.BlackKingKingsideCastlingPlace {
			s.Matrix[peaceplaces.BlackRookKingsideStartingPlace] = peace.BlackRook
			s.Matrix[peaceplaces.BlackRookKingsideCastlingPlace] = peace.NoFigure
		} else if to == peaceplaces.BlackKingQueensideCastlingPlace {
			s.Matrix[peaceplaces.BlackRookQueensideStartingPlace] = peace.BlackRook
			s.Matrix[peaceplaces.BlackRookQueensideCastlingPlace] = peace.NoFigure
		}
	}

	s.Matrix[from] = s.Matrix[to]
	s.Matrix[to] = snapshot.Captured
}

func (s *State) Moves() peacemoves.Fulls {
	var moves peacemoves.Fulls

	whiteKing := s.Matrix.FindSinglePeace(peace.WhiteKing)
	blackKing := s.Matrix.FindSinglePeace(peace.BlackKing)

	whiteFromTos := s.Matrix.WhiteTos(whiteKing)
	for _, fromTos := range whiteFromTos {
		var answers peacemoves.Answers
		for _, to := range fromTos.Tos {
			snapshot := s.Do(fromTos.From, to)

			blackFromTos := s.Matrix.BlackTos(blackKing)
			answers = append(answers, peacemoves.Answer{
				WhiteTo:      to,
				BlackAnswers: blackFromTos,
			})

			s.Undo(snapshot, fromTos.From, to)
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
			snapshot := s.Do(fromTos.From, to)

			mate := s.Matrix.IsSquareUnderAttackFromWhite(blackKing) && s.Matrix.BlackTos(blackKing) == nil

			s.Undo(snapshot, fromTos.From, to)

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
			snapshot := s.Do(fromTos.From, to)

			if s.Matrix.IsSquareUnderAttackFromWhite(blackKing) && s.Matrix.BlackTos(blackKing) == nil {
				mates++
			}

			s.Undo(snapshot, fromTos.From, to)
		}
	}

	return mates
}
