package board

import (
	"github.com/ggeorgiev/instant-chess/src/move"
	"github.com/ggeorgiev/instant-chess/src/peace"
	"github.com/ggeorgiev/instant-chess/src/peaceplaces"
	"github.com/ggeorgiev/instant-chess/src/square"
)

func (s *State) CanBlackEnPassant(file square.File) bool {
	return file > 0 && s.Matrix[square.NewIndex(file-1, peaceplaces.WhitePawnsJumpRank)] == peace.BlackPawn ||
		file < square.LastFile && s.Matrix[square.NewIndex(file+1, peaceplaces.WhitePawnsJumpRank)] == peace.BlackPawn
}

func (s *State) DoWhite(from square.Index, to square.Index) Snapshot {
	snapshot := Snapshot{
		Captured:      s.Matrix[to],
		Castling:      s.Rights.WhiteCastling,
		EnPassantFile: s.Rights.EnPassantFile,
	}

	if from == peaceplaces.WhiteKingStartingPlace && s.Matrix[from] == peace.WhiteKing {
		s.Rights.WhiteCastling = move.NoCastling

		if to == peaceplaces.WhiteKingKingsideCastlingPlace {
			s.Matrix[peaceplaces.WhiteRookKingsideStartingPlace] = peace.NoFigure
			s.Matrix[peaceplaces.WhiteRookKingsideCastlingPlace] = peace.WhiteRook
		} else if to == peaceplaces.WhiteKingQueensideCastlingPlace {
			s.Matrix[peaceplaces.WhiteRookQueensideStartingPlace] = peace.NoFigure
			s.Matrix[peaceplaces.WhiteRookQueensideCastlingPlace] = peace.WhiteRook
		}
	}

	if from == peaceplaces.WhiteRookKingsideStartingPlace && s.Matrix[from] == peace.WhiteRook {
		s.Rights.WhiteCastling.Kingside = false
	}
	if from == peaceplaces.WhiteRookQueensideStartingPlace && s.Matrix[from] == peace.WhiteRook {
		s.Rights.WhiteCastling.Queenside = false
	}

	if from.Rank() == peaceplaces.WhitePawnsStartingRank &&
		to.Rank() == peaceplaces.WhitePawnsJumpRank &&
		s.Matrix[from] == peace.WhitePawn &&
		s.CanBlackEnPassant(to.File()) {
		s.Rights.EnPassantFile = from.File()
	} else {
		s.Rights.EnPassantFile = square.InvalidFile
	}

	s.Matrix[to] = s.Matrix[from]
	s.Matrix[from] = peace.NoFigure

	return snapshot
}

func (s *State) UndoWhite(snapshot Snapshot, from square.Index, to square.Index) {
	s.Rights.WhiteCastling = snapshot.Castling
	s.Rights.EnPassantFile = snapshot.EnPassantFile

	if from == peaceplaces.WhiteKingStartingPlace && s.Matrix[to] == peace.WhiteKing {
		if to == peaceplaces.WhiteKingKingsideCastlingPlace {
			s.Matrix[peaceplaces.WhiteRookKingsideStartingPlace] = peace.WhiteRook
			s.Matrix[peaceplaces.WhiteRookKingsideCastlingPlace] = peace.NoFigure
		} else if to == peaceplaces.WhiteKingQueensideCastlingPlace {
			s.Matrix[peaceplaces.WhiteRookQueensideStartingPlace] = peace.WhiteRook
			s.Matrix[peaceplaces.WhiteRookQueensideCastlingPlace] = peace.NoFigure
		}
	}

	s.Matrix[from] = s.Matrix[to]
	s.Matrix[to] = snapshot.Captured
}
