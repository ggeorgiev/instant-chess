package chess

import "github.com/ggeorgiev/instant-chess/src/square"

func (board Board) SquareUnderAttackFromWhite(s square.Index) bool {
	attackedFromKing := AttackedFromKing[s]
	for _, kingSquare := range attackedFromKing {
		if board[kingSquare] == WhiteKing {
			return true
		}
	}

	attackedFromKnight := AttackedFromKnight[s]
	for _, knightSquare := range attackedFromKnight {
		if board[knightSquare] == WhiteKnight {
			return true
		}
	}

	x := s.X()
	y := s.Y()

	for i := x - 1; i >= 0; i-- {
		peace := board[square.NewIndex(i, y)]
		if peace.IsWhiteLinearMover() {
			return true
		}
		if peace != Empty {
			break
		}
	}
	for i := x + 1; i < 8; i++ {
		peace := board[square.NewIndex(i, y)]
		if peace.IsWhiteLinearMover() {
			return true
		}
		if peace != Empty {
			break
		}
	}
	for i := y - 1; i >= 0; i-- {
		peace := board[square.NewIndex(x, i)]
		if peace.IsWhiteLinearMover() {
			return true
		}
		if peace != Empty {
			break
		}
	}
	for i := y + 1; i < 8; i++ {
		peace := board[square.NewIndex(x, i)]
		if peace.IsWhiteLinearMover() {
			return true
		}
		if peace != Empty {
			break
		}
	}

	i := x - 1
	j := y - 1
	for i >= 0 && j >= 0 {
		peace := board[square.NewIndex(i, j)]
		if peace.IsWhiteDiagonalMover() {
			return true
		}
		if peace != Empty {
			break
		}
		i--
		j--
	}

	i = x - 1
	j = y + 1
	for i >= 0 && j < 8 {
		peace := board[square.NewIndex(i, j)]
		if peace.IsWhiteDiagonalMover() {
			return true
		}
		if peace != Empty {
			break
		}
		i--
		j++
	}

	i = x + 1
	j = y - 1
	for i < 8 && j >= 0 {
		peace := board[square.NewIndex(i, j)]
		if peace.IsWhiteDiagonalMover() {
			return true
		}
		if peace != Empty {
			break
		}
		i++
		j--
	}

	i = x + 1
	j = y + 1
	for i < 8 && j < 8 {
		peace := board[square.NewIndex(i, j)]
		if peace.IsWhiteDiagonalMover() {
			return true
		}
		if peace != Empty {
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
	board[s] = Empty

	kingMoves := KingMoves[s]

	for _, square := range kingMoves {
		peace := board[square]
		if peace.IsEmptyOrBlack() && !board.SquareUnderAttackFromBlack(square) {
			tos = append(tos, square)
		}
	}

	board[s] = original
	return tos
}

func (board Board) WhiteKnightTos(s square.Index, ks square.Index) square.Indexes {
	var tos square.Indexes

	check := func(square square.Index) {
		peace := board[square]
		if peace.IsEmptyOrBlack() {
			original := board[square]
			board[square] = board[s]
			board[s] = Empty

			if !board.SquareUnderAttackFromBlack(ks) {
				tos = append(tos, square)
			}

			board[s] = board[square]
			board[square] = original
		}
	}

	knightMoves := KnightMoves[s]

	for _, square := range knightMoves {
		check(square)
	}

	return tos
}

func (board Board) WhiteRookTos(s square.Index, ks square.Index) square.Indexes {
	var tos square.Indexes

	check := func(square square.Index) bool {
		peace := board[square]
		if peace.IsEmptyOrBlack() {
			original := board[square]
			board[square] = board[s]
			board[s] = Empty

			if !board.SquareUnderAttackFromBlack(ks) {
				tos = append(tos, square)
			}

			board[s] = board[square]
			board[square] = original
		}
		return peace.IsEmpty()
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
	if board[s] == WhiteKing {
		return board.WhiteKingTos(s)
	}
	if board[s] == WhiteRook {
		return board.WhiteRookTos(s, kingSquare)
	}
	if board[s] == WhiteKnight {
		return board.WhiteKnightTos(s, kingSquare)
	}
	return nil
}
