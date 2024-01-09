package board

import (
	"github.com/ggeorgiev/instant-chess/src/bitboard"
	"github.com/ggeorgiev/instant-chess/src/peace"
	"github.com/ggeorgiev/instant-chess/src/peaceattacks"
	"github.com/ggeorgiev/instant-chess/src/square"
)

func (m Matrix) AttackBitboardMaskFromBlack() bitboard.Mask {
	attackerOccupiedMask := bitboard.Empty
	attackMask := bitboard.Empty

	fallShadowMask := bitboard.Empty
	fallLeftShadowMask := bitboard.Empty
	for s := square.ZeroIndex; s <= square.LastIndex; s++ {
		figure := m[s]
		if figure == peace.NoFigure {
			continue
		}

		if figure.Color() == peace.BlackColor {
			attackerOccupiedMask |= square.IndexMask[s]
		}

		if figure == peace.BlackPawn {
			attackMask |= peaceattacks.BlackPawnBitboardMasks[s]
		} else if figure == peace.BlackKnight {
			attackMask |= peaceattacks.KnightBitboardMasks[s]
		} else if figure == peace.BlackBishop {
			attackMask |= peaceattacks.DiagonalsFallBitboardMasks[s] & ^fallShadowMask
		} else if figure == peace.BlackRook {
			attackMask |= peaceattacks.LinearsFallLeftBitboardMasks[s] & ^fallLeftShadowMask
		} else if figure == peace.BlackQueen {
			attackMask |= (peaceattacks.LinearsFallLeftBitboardMasks[s] & ^fallLeftShadowMask) |
				(peaceattacks.DiagonalsFallBitboardMasks[s] & ^fallShadowMask)
		} else if figure == peace.BlackKing {
			attackMask |= peaceattacks.KingBitboardMasks[s]
		}

		fallShadowMask |= peaceattacks.DiagonalsFallBitboardMasks[s]
		fallLeftShadowMask |= peaceattacks.LinearsFallLeftBitboardMasks[s]
	}

	riseShadowMask := bitboard.Empty
	riseRightShadowMask := bitboard.Empty
	for s := square.LastIndex; s >= square.ZeroIndex; s-- {
		figure := m[s]
		if figure == peace.NoFigure {
			continue
		}

		if figure.Color() == peace.BlackColor {
			if figure == peace.BlackBishop {
				attackMask |= peaceattacks.DiagonalsRiseBitboardMasks[s] & ^riseShadowMask
			} else if figure == peace.BlackRook {
				attackMask |= peaceattacks.LinearsRiseRightBitboardMasks[s] & ^riseRightShadowMask
			} else if figure == peace.BlackQueen {
				attackMask |= (peaceattacks.LinearsRiseRightBitboardMasks[s] & ^riseRightShadowMask) |
					(peaceattacks.DiagonalsRiseBitboardMasks[s] & ^riseShadowMask)
			}
		}
		riseShadowMask |= peaceattacks.DiagonalsRiseBitboardMasks[s]
		riseRightShadowMask |= peaceattacks.LinearsRiseRightBitboardMasks[s]
	}

	return attackMask & ^attackerOccupiedMask
}
