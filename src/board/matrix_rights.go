package board

import (
	"github.com/ggeorgiev/instant-chess/src/move"
	"github.com/ggeorgiev/instant-chess/src/peace"
	"github.com/ggeorgiev/instant-chess/src/peaceplaces"
	"github.com/ggeorgiev/instant-chess/src/square"
)

func (m Matrix) MoveRights() move.RightsList {
	whiteCasting := []move.Castling{move.NoCastling}
	if m[peaceplaces.WhiteKingStartingPlace] == peace.WhiteKing {
		if m[peaceplaces.WhiteRookKingsideStartingPlace] == peace.WhiteRook {
			whiteCasting = append(whiteCasting, move.KingsideCastling)
		}
		if m[peaceplaces.WhiteRookQueensideStartingPlace] == peace.WhiteRook {
			whiteCasting = append(whiteCasting, move.QueensideCastling)
		}
		if len(whiteCasting) == 3 {
			whiteCasting = append(whiteCasting, move.BothCastlings)
		}
	}

	blackCaasting := []move.Castling{move.NoCastling}
	if m[peaceplaces.BlackKingStartingPlace] == peace.BlackKing {
		if m[peaceplaces.BlackRookKingsideStartingPlace] == peace.BlackRook {
			blackCaasting = append(blackCaasting, move.KingsideCastling)
		}
		if m[peaceplaces.BlackRookQueensideStartingPlace] == peace.BlackRook {
			blackCaasting = append(blackCaasting, move.QueensideCastling)
		}
		if len(blackCaasting) == 3 {
			blackCaasting = append(blackCaasting, move.BothCastlings)
		}
	}

	enPassant := square.Files{square.InvalidFile}

	for f := square.ZeroFile; f <= square.LastFile; f++ {
		if m[square.NewIndex(f, peaceplaces.BLackJumpRank)] == peace.BlackPawn {
			if f > square.ZeroFile && m[square.NewIndex(f-1, peaceplaces.BLackJumpRank)] == peace.WhitePawn ||
				f < square.LastFile && m[square.NewIndex(f+1, peaceplaces.BLackJumpRank)] == peace.WhitePawn {
				enPassant = append(enPassant, f)
			}
		}
	}

	return move.CombineRights(whiteCasting, blackCaasting, enPassant)
}
