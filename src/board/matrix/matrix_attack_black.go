package matrix

import (
	"github.com/ggeorgiev/instant-chess/src/attack"
	"github.com/ggeorgiev/instant-chess/src/bitboard"
	"github.com/ggeorgiev/instant-chess/src/peace"
	"github.com/ggeorgiev/instant-chess/src/square"
)

func (m *Matrix) AttackBitboardMaskFromBlack() bitboard.Mask {
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
			attackMask |= attack.BlackPawnBitboardMasks[s]
		} else if figure == peace.BlackKnight {
			attackMask |= attack.KnightBitboardMasks[s]
		} else if figure == peace.BlackBishop {
			attackMask |= attack.DiagonalsFallBitboardMasks[s] & ^fallShadowMask
		} else if figure == peace.BlackRook {
			attackMask |= attack.LinearsFallLeftBitboardMasks[s] & ^fallLeftShadowMask
		} else if figure == peace.BlackQueen {
			attackMask |= (attack.LinearsFallLeftBitboardMasks[s] & ^fallLeftShadowMask) |
				(attack.DiagonalsFallBitboardMasks[s] & ^fallShadowMask)
		} else if figure == peace.BlackKing {
			attackMask |= attack.KingBitboardMasks[s]
		}

		fallShadowMask |= attack.DiagonalsFallBitboardMasks[s]
		fallLeftShadowMask |= attack.LinearsFallLeftBitboardMasks[s]
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
				attackMask |= attack.DiagonalsRiseBitboardMasks[s] & ^riseShadowMask
			} else if figure == peace.BlackRook {
				attackMask |= attack.LinearsRiseRightBitboardMasks[s] & ^riseRightShadowMask
			} else if figure == peace.BlackQueen {
				attackMask |= (attack.LinearsRiseRightBitboardMasks[s] & ^riseRightShadowMask) |
					(attack.DiagonalsRiseBitboardMasks[s] & ^riseShadowMask)
			}
		}
		riseShadowMask |= attack.DiagonalsRiseBitboardMasks[s]
		riseRightShadowMask |= attack.LinearsRiseRightBitboardMasks[s]
	}

	return attackMask & ^attackerOccupiedMask
}
