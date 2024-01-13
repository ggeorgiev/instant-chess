package board

import (
	"github.com/ggeorgiev/instant-chess/src/move"
	"github.com/ggeorgiev/instant-chess/src/peace"
	"github.com/ggeorgiev/instant-chess/src/peaceplaces"
	"github.com/ggeorgiev/instant-chess/src/square"
)

func (s *State) CanWhiteEnPassant(file square.File) bool {
	return file > 0 && s.Matrix[square.NewIndex(file-1, peaceplaces.BlackPawnsJumpRank)] == peace.WhitePawn ||
		file < square.LastFile && s.Matrix[square.NewIndex(file+1, peaceplaces.BlackPawnsJumpRank)] == peace.WhitePawn
}

func (s *State) DoBlack(from square.Index, to square.Index) Snapshot {
	snapshot := Snapshot{
		Captured:      s.Matrix[to],
		Castling:      s.Rights.BlackCastling,
		EnPassantFile: s.Rights.EnPassantFile,
	}

	if from == peaceplaces.BlackKingStartingPlace && s.Matrix[from] == peace.BlackKing {
		s.Rights.BlackCastling = move.NoCastling

		if to == peaceplaces.BlackKingKingsideCastlingPlace {
			s.Matrix[peaceplaces.BlackRookKingsideStartingPlace] = peace.NoFigure
			s.Matrix[peaceplaces.BlackRookKingsideCastlingPlace] = peace.BlackRook
		} else if to == peaceplaces.BlackKingQueensideCastlingPlace {
			s.Matrix[peaceplaces.BlackRookQueensideStartingPlace] = peace.NoFigure
			s.Matrix[peaceplaces.BlackRookQueensideCastlingPlace] = peace.BlackRook
		}
	}

	if from == peaceplaces.BlackRookKingsideStartingPlace && s.Matrix[from] == peace.BlackRook {
		s.Rights.BlackCastling.Kingside = false
	}
	if from == peaceplaces.BlackRookQueensideStartingPlace && s.Matrix[from] == peace.BlackRook {
		s.Rights.BlackCastling.Queenside = false
	}

	if from.Rank() == peaceplaces.BlackPawnsStartingRank &&
		to.Rank() == peaceplaces.BlackPawnsJumpRank &&
		s.Matrix[from] == peace.BlackPawn &&
		s.CanWhiteEnPassant(to.File()) {
		s.Rights.EnPassantFile = from.File()
	} else {
		s.Rights.EnPassantFile = square.InvalidFile
	}

	s.Matrix[to] = s.Matrix[from]
	s.Matrix[from] = peace.NoFigure

	return snapshot
}

func (s *State) UndoBlack(snapshot Snapshot, from square.Index, to square.Index) {
	s.Rights.BlackCastling = snapshot.Castling
	s.Rights.EnPassantFile = snapshot.EnPassantFile

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