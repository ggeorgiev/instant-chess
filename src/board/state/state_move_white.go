package state

import (
	"github.com/ggeorgiev/instant-chess/src/peace"
	"github.com/ggeorgiev/instant-chess/src/place"
	"github.com/ggeorgiev/instant-chess/src/square"
)

func (s *State) CanBlackEnPassant(file square.File) bool {
	return file > 0 && s.Matrix[square.NewIndex(file-1, place.WhitePawnsJumpRank)] == peace.BlackPawn ||
		file < square.LastFile && s.Matrix[square.NewIndex(file+1, place.WhitePawnsJumpRank)] == peace.BlackPawn
}

func (s *State) DoWhite(from square.Index, to square.Index) Snapshot {
	snapshot := Snapshot{
		Captured: s.Matrix[to],
		Rights:   s.Rights,
	}

	if from == place.WhiteKingStarting && s.Matrix[from] == peace.WhiteKing {
		s.Rights = s.Rights.ResetWhiteBothCastling()

		if to == place.WhiteKingKingsideCastling {
			s.Matrix[place.WhiteRookKingsideStarting] = peace.Null
			s.Matrix[place.WhiteRookKingsideCastling] = peace.WhiteRook
		} else if to == place.WhiteKingQueensideCastling {
			s.Matrix[place.WhiteRookQueensideStarting] = peace.Null
			s.Matrix[place.WhiteRookQueensideCastling] = peace.WhiteRook
		}
	}

	if from == place.WhiteRookKingsideStarting && s.Matrix[from] == peace.WhiteRook {
		s.Rights = s.Rights.ResetWhiteKingsideCastling()
	}
	if from == place.WhiteRookQueensideStarting && s.Matrix[from] == peace.WhiteRook {
		s.Rights = s.Rights.ResetWhiteQueensideCastling()
	}

	if from.Rank() == place.WhitePawnsStartingRank &&
		to.Rank() == place.WhitePawnsJumpRank &&
		s.Matrix[from] == peace.WhitePawn &&
		s.CanBlackEnPassant(to.File()) {
		s.Rights = s.Rights.SetEnPassantFile(from.File())
	} else {
		s.Rights = s.Rights.SetEnPassantFile(square.InvalidFile)
	}

	s.Matrix[to] = s.Matrix[from]
	s.Matrix[from] = peace.Null

	return snapshot
}

func (s *State) UndoWhite(snapshot Snapshot, from square.Index, to square.Index) {
	s.Rights = snapshot.Rights

	if from == place.WhiteKingStarting && s.Matrix[to] == peace.WhiteKing {
		if to == place.WhiteKingKingsideCastling {
			s.Matrix[place.WhiteRookKingsideStarting] = peace.WhiteRook
			s.Matrix[place.WhiteRookKingsideCastling] = peace.Null
		} else if to == place.WhiteKingQueensideCastling {
			s.Matrix[place.WhiteRookQueensideStarting] = peace.WhiteRook
			s.Matrix[place.WhiteRookQueensideCastling] = peace.Null
		}
	}

	s.Matrix[from] = s.Matrix[to]
	s.Matrix[to] = snapshot.Captured
}
