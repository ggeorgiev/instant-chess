package chess

import (
	"github.com/ggeorgiev/instant-chess/src/bitboard"
	"github.com/ggeorgiev/instant-chess/src/board"
	"github.com/ggeorgiev/instant-chess/src/peace"
	"github.com/ggeorgiev/instant-chess/src/peaceattacks"
	"github.com/ggeorgiev/instant-chess/src/peacemoves"
	"github.com/ggeorgiev/instant-chess/src/square"
)

func (brd Board) BlackKingTos(s square.Index) square.Indexes {
	var tos square.Indexes

	original := brd[s]
	brd[s] = peace.NoFigure

	kingMoves := peacemoves.KingSquareIndexes[s]

	for _, square := range kingMoves {
		figure := brd[square]
		if figure.IsNoFigureOrWhite() && !board.Matrix(brd).IsSquareUnderAttackFromWhite(square) {
			tos = append(tos, square)
		}
	}

	brd[s] = original
	return tos
}

func (brd Board) BlackKnightTos(s square.Index, ks square.Index) square.Indexes {
	var tos square.Indexes

	check := func(square square.Index) {
		figure := brd[square]
		if figure.IsNoFigureOrWhite() {
			original := brd[square]
			brd[square] = brd[s]
			brd[s] = peace.NoFigure

			if !board.Matrix(brd).IsSquareUnderAttackFromWhite(ks) {
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

func (brd Board) BlackRookTos(s square.Index, ks square.Index) square.Indexes {
	var tos square.Indexes

	check := func(square square.Index) bool {
		figure := brd[square]
		if figure.IsNoFigureOrWhite() {
			original := brd[square]
			brd[square] = brd[s]
			brd[s] = peace.NoFigure

			if !board.Matrix(brd).IsSquareUnderAttackFromWhite(ks) {
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

func (brd Board) BlackTos(s square.Index, kingSquare square.Index) square.Indexes {
	switch brd[s] {
	case peace.BlackKing:
		return brd.BlackKingTos(s)
	case peace.BlackRook:
		return brd.BlackRookTos(s, kingSquare)
	case peace.BlackKnight:
		return brd.BlackKnightTos(s, kingSquare)
	}
	return nil
}

func (brd Board) AttackBitboardMaskFromBlack() bitboard.Mask {
	attackerOccupiedMask := bitboard.Empty
	attackMask := bitboard.Empty

	fallShadowMask := bitboard.Empty
	fallLeftShadowMask := bitboard.Empty
	for s := square.ZeroIndex; s <= square.LastIndex; s++ {
		figure := brd[s]
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
		figure := brd[s]
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
