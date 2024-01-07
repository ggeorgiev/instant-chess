package chess

import (
	"github.com/ggeorgiev/instant-chess/src/bitboard"
	"github.com/ggeorgiev/instant-chess/src/board"
	"github.com/ggeorgiev/instant-chess/src/peace"
	"github.com/ggeorgiev/instant-chess/src/peaceattacks"
	"github.com/ggeorgiev/instant-chess/src/peacemoves"
	"github.com/ggeorgiev/instant-chess/src/square"
)

func (brd Board) WhiteKnightTos(s square.Index, ks square.Index) square.Indexes {
	var tos square.Indexes

	check := func(square square.Index) {
		figure := brd[square]
		if figure.IsNoFigureOrBlack() {
			original := brd[square]
			brd[square] = brd[s]
			brd[s] = peace.NoFigure

			if !board.Matrix(brd).IsSquareUnderAttackFromBlack(ks) {
				tos = append(tos, square)
			}

			brd[s] = brd[square]
			brd[square] = original
		}
	}

	knightMoves := peacemoves.KnightSquareIndexes[s]

	for _, square := range knightMoves {
		check(square)
	}

	return tos
}

func (brd Board) WhiteRookTos(s square.Index, ks square.Index) square.Indexes {
	var tos square.Indexes

	check := func(square square.Index) bool {
		figure := brd[square]
		if figure.IsNoFigureOrBlack() {
			original := brd[square]
			brd[square] = brd[s]
			brd[s] = peace.NoFigure

			if !board.Matrix(brd).IsSquareUnderAttackFromBlack(ks) {
				tos = append(tos, square)
			}

			brd[s] = brd[square]
			brd[square] = original
		}
		return figure.IsNoFigure()
	}

	file := s.File()
	rank := s.Rank()

	for f := file - 1; f >= square.ZeroFile; f-- {
		if !check(square.NewIndex(f, rank)) {
			break
		}
	}
	for f := file + 1; f <= square.LastFile; f++ {
		if !check(square.NewIndex(f, rank)) {
			break
		}
	}

	for r := rank - 1; r >= square.ZeroRank; r-- {
		if !check(square.NewIndex(file, r)) {
			break
		}
	}
	for r := rank + 1; r <= square.LastRank; r++ {
		if !check(square.NewIndex(file, r)) {
			break
		}
	}
	return tos
}

func (brd Board) WhiteTos(s square.Index, kingSquare square.Index) square.Indexes {
	switch brd[s] {
	case peace.WhiteKing:
		return board.Matrix(brd).WhiteKingTos(s)
	case peace.WhiteRook:
		return brd.WhiteRookTos(s, kingSquare)
	case peace.WhiteKnight:
		return brd.WhiteKnightTos(s, kingSquare)
	}
	return nil
}

func (brd Board) AttackBitboardMaskFromWhite() bitboard.Mask {
	attackerOccupiedMask := bitboard.Empty
	attackMask := bitboard.Empty

	fallShadowMask := bitboard.Empty
	fallLeftShadowMask := bitboard.Empty
	for s := square.ZeroIndex; s <= square.LastIndex; s++ {
		figure := brd[s]
		if figure == peace.NoFigure {
			continue
		}

		if figure.Color() == peace.WhiteColor {
			attackerOccupiedMask |= square.IndexMask[s]
		}

		if figure == peace.WhitePawn {
			attackMask |= peaceattacks.WhitePawnBitboardMasks[s]
		} else if figure == peace.WhiteKnight {
			attackMask |= peaceattacks.KnightBitboardMasks[s]
		} else if figure == peace.WhiteBishop {
			attackMask |= peaceattacks.DiagonalsFallBitboardMasks[s] & ^fallShadowMask
		} else if figure == peace.WhiteRook {
			attackMask |= peaceattacks.LinearsFallLeftBitboardMasks[s] & ^fallLeftShadowMask
		} else if figure == peace.WhiteQueen {
			attackMask |= (peaceattacks.LinearsFallLeftBitboardMasks[s] & ^fallLeftShadowMask) |
				(peaceattacks.DiagonalsFallBitboardMasks[s] & ^fallShadowMask)
		} else if figure == peace.WhiteKing {
			attackMask |= peaceattacks.KingBitboardMasks[s]
		}

		fallShadowMask |= peaceattacks.DiagonalsFallBitboardMasks[s]
		fallLeftShadowMask |= peaceattacks.LinearsFallLeftBitboardMasks[s]
	}

	riseShadowMask := bitboard.Empty
	riseRightShadowMask := bitboard.Empty
	for s := square.LastIndex; s >= square.ZeroIndex; s-- {
		figure := brd[s]
		if figure == peace.NoFigure {
			continue
		}

		if figure.Color() == peace.WhiteColor {
			if figure == peace.WhiteBishop {
				attackMask |= peaceattacks.DiagonalsRiseBitboardMasks[s] & ^riseShadowMask
			} else if figure == peace.WhiteRook {
				attackMask |= peaceattacks.LinearsRiseRightBitboardMasks[s] & ^riseRightShadowMask
			} else if figure == peace.WhiteQueen {
				attackMask |= (peaceattacks.LinearsRiseRightBitboardMasks[s] & ^riseRightShadowMask) |
					(peaceattacks.DiagonalsRiseBitboardMasks[s] & ^riseShadowMask)
			}
		}
		riseShadowMask |= peaceattacks.DiagonalsRiseBitboardMasks[s]
		riseRightShadowMask |= peaceattacks.LinearsRiseRightBitboardMasks[s]
	}

	return attackMask & ^attackerOccupiedMask
}
