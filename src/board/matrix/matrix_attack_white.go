package matrix

import (
	"github.com/ggeorgiev/instant-chess/src/attack"
	"github.com/ggeorgiev/instant-chess/src/bitboard"
	"github.com/ggeorgiev/instant-chess/src/peace"
	"github.com/ggeorgiev/instant-chess/src/square"
)

func (m *Matrix) AttackBitboardMaskFromWhite() bitboard.Mask {
	attackerOccupiedMask := bitboard.Empty
	attackMask := bitboard.Empty

	fallShadowMask := bitboard.Empty
	fallLeftShadowMask := bitboard.Empty
	for s := square.ZeroIndex; s <= square.LastIndex; s++ {
		pc := m[s]
		if pc == peace.Null {
			continue
		}

		if pc.Color() == peace.WhiteColor {
			attackerOccupiedMask |= square.IndexMask[s]
		}

		if pc == peace.WhitePawn {
			attackMask |= attack.WhitePawnBitboardMasks[s]
		} else if pc == peace.WhiteKnight {
			attackMask |= attack.KnightBitboardMasks[s]
		} else if pc == peace.WhiteBishop {
			attackMask |= attack.DiagonalsFallBitboardMasks[s] & ^fallShadowMask
		} else if pc == peace.WhiteRook {
			attackMask |= attack.LinearsFallLeftBitboardMasks[s] & ^fallLeftShadowMask
		} else if pc == peace.WhiteQueen {
			attackMask |= (attack.LinearsFallLeftBitboardMasks[s] & ^fallLeftShadowMask) |
				(attack.DiagonalsFallBitboardMasks[s] & ^fallShadowMask)
		} else if pc == peace.WhiteKing {
			attackMask |= attack.KingBitboardMasks[s]
		}

		fallShadowMask |= attack.DiagonalsFallBitboardMasks[s]
		fallLeftShadowMask |= attack.LinearsFallLeftBitboardMasks[s]
	}

	riseShadowMask := bitboard.Empty
	riseRightShadowMask := bitboard.Empty
	for s := square.LastIndex; s >= square.ZeroIndex; s-- {
		pc := m[s]
		if pc == peace.Null {
			continue
		}

		if pc.Color() == peace.WhiteColor {
			if pc == peace.WhiteBishop {
				attackMask |= attack.DiagonalsRiseBitboardMasks[s] & ^riseShadowMask
			} else if pc == peace.WhiteRook {
				attackMask |= attack.LinearsRiseRightBitboardMasks[s] & ^riseRightShadowMask
			} else if pc == peace.WhiteQueen {
				attackMask |= (attack.LinearsRiseRightBitboardMasks[s] & ^riseRightShadowMask) |
					(attack.DiagonalsRiseBitboardMasks[s] & ^riseShadowMask)
			}
		}
		riseShadowMask |= attack.DiagonalsRiseBitboardMasks[s]
		riseRightShadowMask |= attack.LinearsRiseRightBitboardMasks[s]
	}

	return attackMask & ^attackerOccupiedMask
}
