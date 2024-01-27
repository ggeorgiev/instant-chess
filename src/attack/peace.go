package attack

import (
	"github.com/ggeorgiev/instant-chess/src/bitboard"
	"github.com/ggeorgiev/instant-chess/src/peace"
)

func PeaceBitboardMasks(pc peace.Code) bitboard.Masks {
	switch pc {
	case peace.WhitePawn:
		return WhitePawnBitboardMasks
	case peace.BlackPawn:
		return BlackPawnBitboardMasks
	case peace.WhiteKnight, peace.BlackKnight:
		return KnightBitboardMasks
	case peace.WhiteBishop, peace.BlackBishop:
		return DiagonalsBitboardMasks
	case peace.WhiteRook, peace.BlackRook:
		return LinearsBitboardMasks
	case peace.WhiteQueen, peace.BlackQueen:
	case peace.WhiteKing, peace.BlackKing:
		return KingBitboardMasks
	}
	return nil
}
