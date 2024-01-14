package board

import (
	"github.com/ggeorgiev/instant-chess/src/move"
	"github.com/ggeorgiev/instant-chess/src/peace"
	"github.com/ggeorgiev/instant-chess/src/peaceplaces"
	"github.com/ggeorgiev/instant-chess/src/square"
)

func (m *Matrix) MoveRights() move.RightsList {
	whiteCasting := move.RightsList{move.NoRights}
	if m[peaceplaces.WhiteKingStartingPlace] == peace.WhiteKing {
		if m[peaceplaces.WhiteRookKingsideStartingPlace] == peace.WhiteRook {
			whiteCasting = append(whiteCasting, move.WhiteKingsideCastlingRights)
		}
		if m[peaceplaces.WhiteRookQueensideStartingPlace] == peace.WhiteRook {
			whiteCasting = append(whiteCasting, move.WhiteQueensideCastlingRights)
		}
		if len(whiteCasting) == 3 {
			whiteCasting = append(whiteCasting, move.WhiteBothCastlingRights)
		}
	}

	blackCaasting := move.RightsList{move.NoRights}
	if m[peaceplaces.BlackKingStartingPlace] == peace.BlackKing {
		if m[peaceplaces.BlackRookKingsideStartingPlace] == peace.BlackRook {
			blackCaasting = append(blackCaasting, move.BlackKingsideCastlingRights)
		}
		if m[peaceplaces.BlackRookQueensideStartingPlace] == peace.BlackRook {
			blackCaasting = append(blackCaasting, move.BlackQueensideCastlingRights)
		}
		if len(blackCaasting) == 3 {
			blackCaasting = append(blackCaasting, move.BlackBothCastlingRights)
		}
	}

	enPassantFiles := move.RightsList{move.NoEnPassantFile}
	for f := square.ZeroFile; f <= square.LastFile; f++ {
		if m[square.NewIndex(f, peaceplaces.BlackPawnsJumpRank)] == peace.BlackPawn {
			if f > square.ZeroFile && m[square.NewIndex(f-1, peaceplaces.BlackPawnsJumpRank)] == peace.WhitePawn ||
				f < square.LastFile && m[square.NewIndex(f+1, peaceplaces.BlackPawnsJumpRank)] == peace.WhitePawn {
				enPassantFiles = append(enPassantFiles, move.NoRights.SetEnPassantFile(f))
			}
		}
	}

	return move.CombineRights(whiteCasting, blackCaasting, enPassantFiles)
}
