package chess

import (
	"github.com/ggeorgiev/instant-chess/src/bitboard"
	"github.com/ggeorgiev/instant-chess/src/peace"
	"github.com/ggeorgiev/instant-chess/src/peaceattacks"
	"github.com/ggeorgiev/instant-chess/src/peacemoves"
	"github.com/ggeorgiev/instant-chess/src/square"
)

func (board Board) SquareUnderAttackFromBlack(s square.Index, countTo int) int {
	count := 0
	attackedFromKing := peaceattacks.FromKing[s]
	for _, kingSquare := range attackedFromKing {
		if board[kingSquare] == peace.BlackKing {
			count++
			if count >= countTo {
				return count
			}
		}
	}

	attackedFromKnight := peaceattacks.FromKnight[s]
	for _, knightSquare := range attackedFromKnight {
		if board[knightSquare] == peace.BlackKnight {
			count++
			if count >= countTo {
				return count
			}
		}
	}

	file := s.File()
	rank := s.Rank()

	for f := file - 1; f >= 0; f-- {
		figure := board[square.NewIndex(f, rank)]
		if figure.IsBlackLinearMover() {
			count++
			if count >= countTo {
				return count
			}
		}
		if figure != peace.NoFigure {
			break
		}
	}
	for f := file + 1; f < 8; f++ {
		figure := board[square.NewIndex(f, rank)]
		if figure.IsBlackLinearMover() {
			count++
			if count >= countTo {
				return count
			}
		}
		if figure != peace.NoFigure {
			break
		}
	}
	for r := rank - 1; r >= 0; r-- {
		figure := board[square.NewIndex(file, r)]
		if figure.IsBlackLinearMover() {
			count++
			if count >= countTo {
				return count
			}
		}
		if figure != peace.NoFigure {
			break
		}
	}
	for r := rank + 1; r < 8; r++ {
		figure := board[square.NewIndex(file, r)]
		if figure.IsBlackLinearMover() {
			count++
			if count >= countTo {
				return count
			}
		}
		if figure != peace.NoFigure {
			break
		}
	}

	f := file - 1
	r := rank - 1
	for f >= 0 && r >= 0 {
		figure := board[square.NewIndex(f, r)]
		if figure.IsBlackDiagonalMover() {
			count++
			if count >= countTo {
				return count
			}
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
		if figure.IsBlackDiagonalMover() {
			count++
			if count >= countTo {
				return count
			}
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
		if figure.IsBlackDiagonalMover() {
			count++
			if count >= countTo {
				return count
			}
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
		if figure.IsBlackDiagonalMover() {
			count++
			if count >= countTo {
				return count
			}
		}
		if figure != peace.NoFigure {
			break
		}
		f++
		r++
	}

	return count
}

func (board Board) BlackKingTos(s square.Index) square.Indexes {
	var tos square.Indexes

	original := board[s]
	board[s] = peace.NoFigure

	kingMoves := peacemoves.KingSquareIndexes[s]

	for _, square := range kingMoves {
		figure := board[square]
		if figure.IsNoFigureOrWhite() && board.SquareUnderAttackFromWhite(square, 1) == 0 {
			tos = append(tos, square)
		}
	}

	board[s] = original
	return tos
}

func (board Board) BlackKnightTos(s square.Index, ks square.Index) square.Indexes {
	var tos square.Indexes

	check := func(square square.Index) {
		figure := board[square]
		if figure.IsNoFigureOrWhite() {
			original := board[square]
			board[square] = board[s]
			board[s] = peace.NoFigure

			if board.SquareUnderAttackFromWhite(ks, 1) == 0 {
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

func (board Board) BlackRookTos(s square.Index, ks square.Index) square.Indexes {
	var tos square.Indexes

	check := func(square square.Index) bool {
		figure := board[square]
		if figure.IsNoFigureOrWhite() {
			original := board[square]
			board[square] = board[s]
			board[s] = peace.NoFigure

			if board.SquareUnderAttackFromWhite(ks, 1) == 0 {
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

func (board Board) BlackTos(s square.Index, kingSquare square.Index) square.Indexes {
	switch board[s] {
	case peace.BlackKing:
		return board.BlackKingTos(s)
	case peace.BlackRook:
		return board.BlackRookTos(s, kingSquare)
	case peace.BlackKnight:
		return board.BlackKnightTos(s, kingSquare)
	}
	return nil
}

func (board Board) AttackBitboardMaskFromBlack() bitboard.Mask {
	attackerOccupiedMask := bitboard.Empty
	attackMask := bitboard.Empty

	fallShadowMask := bitboard.Empty
	fallLeftShadowMask := bitboard.Empty
	for s := square.ZeroIndex; s <= square.LastIndex; s++ {
		figure := board[s]
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
		figure := board[s]
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
