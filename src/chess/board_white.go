package chess

import (
	"github.com/ggeorgiev/instant-chess/src/bitboard"
	"github.com/ggeorgiev/instant-chess/src/peace"
	"github.com/ggeorgiev/instant-chess/src/peaceattacks"
	"github.com/ggeorgiev/instant-chess/src/peacemoves"
	"github.com/ggeorgiev/instant-chess/src/square"
)

func (board Board) SquareUnderAttackFromWhite(s square.Index) bool {
	attackedFromKing := peaceattacks.FromKing[s]
	for _, kingSquare := range attackedFromKing {
		if board[kingSquare] == peace.WhiteKing {
			return true
		}
	}

	attackedFromKnight := peaceattacks.FromKnight[s]
	for _, knightSquare := range attackedFromKnight {
		if board[knightSquare] == peace.WhiteKnight {
			return true
		}
	}

	file := s.File()
	rank := s.Rank()

	for f := file - 1; f >= 0; f-- {
		figure := board[square.NewIndex(f, rank)]
		if figure.IsWhiteLinearMover() {
			return true
		}
		if figure != peace.NoFigure {
			break
		}
	}
	for f := file + 1; f < 8; f++ {
		figure := board[square.NewIndex(f, rank)]
		if figure.IsWhiteLinearMover() {
			return true
		}
		if figure != peace.NoFigure {
			break
		}
	}
	for r := rank - 1; r >= 0; r-- {
		figure := board[square.NewIndex(file, r)]
		if figure.IsWhiteLinearMover() {
			return true
		}
		if figure != peace.NoFigure {
			break
		}
	}
	for r := rank + 1; r < 8; r++ {
		figure := board[square.NewIndex(file, r)]
		if figure.IsWhiteLinearMover() {
			return true
		}
		if figure != peace.NoFigure {
			break
		}
	}

	f := file - 1
	r := rank - 1
	for f >= 0 && r >= 0 {
		figure := board[square.NewIndex(f, r)]
		if figure.IsWhiteDiagonalMover() {
			return true
		}
		if figure != peace.NoFigure {
			break
		}
		f--
		r--
	}

	f = file - 1
	r = rank + 1
	for f >= 0 && r < 8 {
		figure := board[square.NewIndex(f, r)]
		if figure.IsWhiteDiagonalMover() {
			return true
		}
		if figure != peace.NoFigure {
			break
		}
		f--
		r++
	}

	f = file + 1
	r = rank - 1
	for f < 8 && r >= 0 {
		figure := board[square.NewIndex(f, r)]
		if figure.IsWhiteDiagonalMover() {
			return true
		}
		if figure != peace.NoFigure {
			break
		}
		f++
		r--
	}

	f = file + 1
	r = rank + 1
	for f < 8 && r < 8 {
		figure := board[square.NewIndex(f, r)]
		if figure.IsWhiteDiagonalMover() {
			return true
		}
		if figure != peace.NoFigure {
			break
		}
		f++
		r++
	}

	return false
}

func (board Board) WhiteKingTos(s square.Index) square.Indexes {
	var tos square.Indexes

	original := board[s]
	board[s] = peace.NoFigure

	kingMoves := peacemoves.KingSquareIndexes[s]

	for _, square := range kingMoves {
		figure := board[square]
		if figure.IsNoFigureOrBlack() && !board.SquareUnderAttackFromBlack(square) {
			tos = append(tos, square)
		}
	}

	board[s] = original
	return tos
}

func (board Board) WhiteKnightTos(s square.Index, ks square.Index) square.Indexes {
	var tos square.Indexes

	check := func(square square.Index) {
		figure := board[square]
		if figure.IsNoFigureOrBlack() {
			original := board[square]
			board[square] = board[s]
			board[s] = peace.NoFigure

			if !board.SquareUnderAttackFromBlack(ks) {
				tos = append(tos, square)
			}

			board[s] = board[square]
			board[square] = original
		}
	}

	knightMoves := peacemoves.KnightSquareIndexes[s]

	for _, square := range knightMoves {
		check(square)
	}

	return tos
}

func (board Board) WhiteRookTos(s square.Index, ks square.Index) square.Indexes {
	var tos square.Indexes

	check := func(square square.Index) bool {
		figure := board[square]
		if figure.IsNoFigureOrBlack() {
			original := board[square]
			board[square] = board[s]
			board[s] = peace.NoFigure

			if !board.SquareUnderAttackFromBlack(ks) {
				tos = append(tos, square)
			}

			board[s] = board[square]
			board[square] = original
		}
		return figure.IsNoFigure()
	}

	file := s.File()
	rank := s.Rank()

	for f := file - 1; f >= 0; f-- {
		if !check(square.NewIndex(f, rank)) {
			break
		}
	}
	for f := file + 1; f < 8; f++ {
		if !check(square.NewIndex(f, rank)) {
			break
		}
	}

	for r := rank - 1; r >= 0; r-- {
		if !check(square.NewIndex(file, r)) {
			break
		}
	}
	for r := rank + 1; r < 8; r++ {
		if !check(square.NewIndex(file, r)) {
			break
		}
	}
	return tos
}

func (board Board) WhiteTos(s square.Index, kingSquare square.Index) square.Indexes {
	switch board[s] {
	case peace.WhiteKing:
		return board.WhiteKingTos(s)
	case peace.WhiteRook:
		return board.WhiteRookTos(s, kingSquare)
	case peace.WhiteKnight:
		return board.WhiteKnightTos(s, kingSquare)
	}
	return nil
}

func (board Board) AttackBitboardMaskFromWhite() bitboard.Mask {
	attackerOccupiedMask := bitboard.Empty
	attackMask := bitboard.Empty

	fallShadowMask := bitboard.Empty
	fallLeftShadowMask := bitboard.Empty
	for s := square.ZeroIndex; s <= square.LastIndex; s++ {
		figure := board[s]
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
		figure := board[s]
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
