package chess

import (
	"github.com/ggeorgiev/instant-chess/src/peace"
	"github.com/ggeorgiev/instant-chess/src/peaceattacks"
	"github.com/ggeorgiev/instant-chess/src/peacemoves"
	"github.com/ggeorgiev/instant-chess/src/square"
)

func (board Board) SquareUnderAttackFromBlack(s square.Index) bool {
	attackedFromKing := AttackedFromKing[s]
	for _, kingSquare := range attackedFromKing {
		if board[kingSquare] == peace.BlackKing {
			return true
		}
	}

	attackedFromKnight := peaceattacks.FromKnight[s]
	for _, knightSquare := range attackedFromKnight {
		if board[knightSquare] == peace.BlackKnight {
			return true
		}
	}

	x := s.X()
	y := s.Y()

	for i := x - 1; i >= 0; i-- {
		figure := board[square.NewIndex(i, y)]
		if figure.IsBlackLinearMover() {
			return true
		}
		if figure != peace.NoFigure {
			break
		}
	}
	for i := x + 1; i < 8; i++ {
		figure := board[square.NewIndex(i, y)]
		if figure.IsBlackLinearMover() {
			return true
		}
		if figure != peace.NoFigure {
			break
		}
	}
	for i := y - 1; i >= 0; i-- {
		figure := board[square.NewIndex(x, i)]
		if figure.IsBlackLinearMover() {
			return true
		}
		if figure != peace.NoFigure {
			break
		}
	}
	for i := y + 1; i < 8; i++ {
		figure := board[square.NewIndex(x, i)]
		if figure.IsBlackLinearMover() {
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
		if figure.IsBlackDiagonalMover() {
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
		if figure.IsBlackDiagonalMover() {
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
		if figure.IsBlackDiagonalMover() {
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
		if figure.IsBlackDiagonalMover() {
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

func (board Board) BlackKingTos(s square.Index) square.Indexes {
	var tos square.Indexes

	original := board[s]
	board[s] = peace.NoFigure

	kingMoves := KingMoves[s]

	for _, square := range kingMoves {
		figure := board[square]
		if figure.IsNoFigureOrWhite() && !board.SquareUnderAttackFromWhite(square) {
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

			if !board.SquareUnderAttackFromWhite(ks) {
				tos = append(tos, square)
			}

			board[s] = board[square]
			board[square] = original
		}
	}

	knightMoves := peacemoves.KnightMoves[s]

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

			if !board.SquareUnderAttackFromWhite(ks) {
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

func (board Board) BlackTos(s square.Index, kingSquare square.Index) square.Indexes {
	if board[s] == peace.BlackKing {
		return board.BlackKingTos(s)
	}
	if board[s] == peace.BlackRook {
		return board.BlackRookTos(s, kingSquare)
	}
	if board[s] == peace.BlackKnight {
		return board.BlackKnightTos(s, kingSquare)
	}
	return nil
}
