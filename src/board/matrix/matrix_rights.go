package matrix

import (
	"github.com/ggeorgiev/instant-chess/src/move"
	"github.com/ggeorgiev/instant-chess/src/peace"
	"github.com/ggeorgiev/instant-chess/src/place"
	"github.com/ggeorgiev/instant-chess/src/square"
)

func (m *Matrix) MoveRights() move.RightsList {
	whiteCasting := move.RightsList{move.NoRights}
	if m[place.WhiteKingStarting] == peace.WhiteKing {
		if m[place.WhiteRookKingsideStarting] == peace.WhiteRook {
			whiteCasting = append(whiteCasting, move.WhiteKingsideCastlingRights)
		}
		if m[place.WhiteRookQueensideStarting] == peace.WhiteRook {
			whiteCasting = append(whiteCasting, move.WhiteQueensideCastlingRights)
		}
		if len(whiteCasting) == 3 {
			whiteCasting = append(whiteCasting, move.WhiteBothCastlingRights)
		}
	}

	blackCasting := move.RightsList{move.NoRights}
	if m[place.BlackKingStarting] == peace.BlackKing {
		if m[place.BlackRookKingsideStarting] == peace.BlackRook {
			blackCasting = append(blackCasting, move.BlackKingsideCastlingRights)
		}
		if m[place.BlackRookQueensideStarting] == peace.BlackRook {
			blackCasting = append(blackCasting, move.BlackQueensideCastlingRights)
		}
		if len(blackCasting) == 3 {
			blackCasting = append(blackCasting, move.BlackBothCastlingRights)
		}
	}

	enPassantFiles := move.RightsList{move.NoEnPassantFile}
	for f := square.ZeroFile; f <= square.LastFile; f++ {
		if m[square.NewIndex(f, place.BlackPawnsJumpRank)] == peace.BlackPawn {
			if f > square.ZeroFile && m[square.NewIndex(f-1, place.BlackPawnsJumpRank)] == peace.WhitePawn ||
				f < square.LastFile && m[square.NewIndex(f+1, place.BlackPawnsJumpRank)] == peace.WhitePawn {
				enPassantFiles = append(enPassantFiles, move.NoRights.SetEnPassantFile(f))
			}
		}
	}

	return move.CombineRights(whiteCasting, blackCasting, enPassantFiles)
}
