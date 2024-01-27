package state

import (
	"github.com/ggeorgiev/instant-chess/src/peace"
	"github.com/ggeorgiev/instant-chess/src/place"
	"github.com/ggeorgiev/instant-chess/src/square"
)

func (s *State) CanWhiteEnPassant(file square.File) bool {
	return file > 0 && s.Matrix[square.NewIndex(file-1, place.BlackPawnsJumpRank)] == peace.WhitePawn ||
		file < square.LastFile && s.Matrix[square.NewIndex(file+1, place.BlackPawnsJumpRank)] == peace.WhitePawn
}

func (s *State) DoBlack(from square.Index, to square.Index) Snapshot {
	snapshot := Snapshot{
		Captured: s.Matrix[to],
		Rights:   s.Rights,
	}

	if from == place.BlackKingStarting && s.Matrix[from] == peace.BlackKing {
		s.Rights = s.Rights.ResetBlackBothCastling()

		if to == place.BlackKingKingsideCastling {
			s.Matrix[place.BlackRookKingsideStarting] = peace.NoFigure
			s.Matrix[place.BlackRookKingsideCastling] = peace.BlackRook
		} else if to == place.BlackKingQueensideCastling {
			s.Matrix[place.BlackRookQueensideStarting] = peace.NoFigure
			s.Matrix[place.BlackRookQueensideCastling] = peace.BlackRook
		}
	}

	if from == place.BlackRookKingsideStarting && s.Matrix[from] == peace.BlackRook {
		s.Rights = s.Rights.ResetBlackKingsideCastling()
	}
	if from == place.BlackRookQueensideStarting && s.Matrix[from] == peace.BlackRook {
		s.Rights = s.Rights.ResetBlackQueensideCastling()
	}

	if from.Rank() == place.BlackPawnsStartingRank &&
		to.Rank() == place.BlackPawnsJumpRank &&
		s.Matrix[from] == peace.BlackPawn &&
		s.CanWhiteEnPassant(to.File()) {
		s.Rights = s.Rights.SetEnPassantFile(from.File())
	} else {
		s.Rights = s.Rights.SetEnPassantFile(square.InvalidFile)
	}

	s.Matrix[to] = s.Matrix[from]
	s.Matrix[from] = peace.NoFigure

	return snapshot
}

func (s *State) UndoBlack(snapshot Snapshot, from square.Index, to square.Index) {
	s.Rights = snapshot.Rights

	if from == place.BlackKingStarting && s.Matrix[to] == peace.BlackKing {
		if to == place.BlackKingKingsideCastling {
			s.Matrix[place.BlackRookKingsideStarting] = peace.BlackRook
			s.Matrix[place.BlackRookKingsideCastling] = peace.NoFigure
		} else if to == place.BlackKingQueensideCastling {
			s.Matrix[place.BlackRookQueensideStarting] = peace.BlackRook
			s.Matrix[place.BlackRookQueensideCastling] = peace.NoFigure
		}
	}

	s.Matrix[from] = s.Matrix[to]
	s.Matrix[to] = snapshot.Captured
}
