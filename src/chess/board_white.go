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

	x := s.X()
	y := s.Y()

	for i := x - 1; i >= 0; i-- {
		figure := board[square.NewIndex(i, y)]
		if figure.IsWhiteLinearMover() {
			return true
		}
		if figure != peace.NoFigure {
			break
		}
	}
	for i := x + 1; i < 8; i++ {
		figure := board[square.NewIndex(i, y)]
		if figure.IsWhiteLinearMover() {
			return true
		}
		if figure != peace.NoFigure {
			break
		}
	}
	for i := y - 1; i >= 0; i-- {
		figure := board[square.NewIndex(x, i)]
		if figure.IsWhiteLinearMover() {
			return true
		}
		if figure != peace.NoFigure {
			break
		}
	}
	for i := y + 1; i < 8; i++ {
		figure := board[square.NewIndex(x, i)]
		if figure.IsWhiteLinearMover() {
			return true
		}
		if figure != peace.NoFigure {
			break
		}
	}

	i := x - 1
	j := y - 1
	for i >= 0 && j >= 0 {
		figure := board[square.NewIndex(i, j)]
		if figure.IsWhiteDiagonalMover() {
			return true
		}
		if figure != peace.NoFigure {
			break
		}
		i--
		j--
	}

	i = x - 1
	j = y + 1
	for i >= 0 && j < 8 {
		figure := board[square.NewIndex(i, j)]
		if figure.IsWhiteDiagonalMover() {
			return true
		}
		if figure != peace.NoFigure {
			break
		}
		i--
		j++
	}

	i = x + 1
	j = y - 1
	for i < 8 && j >= 0 {
		figure := board[square.NewIndex(i, j)]
		if figure.IsWhiteDiagonalMover() {
			return true
		}
		if figure != peace.NoFigure {
			break
		}
		i++
		j--
	}

	i = x + 1
	j = y + 1
	for i < 8 && j < 8 {
		figure := board[square.NewIndex(i, j)]
		if figure.IsWhiteDiagonalMover() {
			return true
		}
		if figure != peace.NoFigure {
			break
		}
		i++
		j++
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

	x := s.X()
	y := s.Y()

	for i := x - 1; i >= 0; i-- {
		if !check(square.NewIndex(i, y)) {
			break
		}
	}
	for i := x + 1; i < 8; i++ {
		if !check(square.NewIndex(i, y)) {
			break
		}
	}

	for i := y - 1; i >= 0; i-- {
		if !check(square.NewIndex(x, i)) {
			break
		}
	}
	for i := y + 1; i < 8; i++ {
		if !check(square.NewIndex(x, i)) {
			break
		}
	}
	return tos
}

func (board Board) WhiteTos(s square.Index, kingSquare square.Index) square.Indexes {
	if board[s] == peace.WhiteKing {
		return board.WhiteKingTos(s)
	}
	if board[s] == peace.WhiteRook {
		return board.WhiteRookTos(s, kingSquare)
	}
	if board[s] == peace.WhiteKnight {
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
