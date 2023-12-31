package chess

func (board Board) SquareUnderAttackFromBlack(s Square) bool {
	attackedFromKing := AttackedFromKing[s]
	for _, kingSquare := range attackedFromKing {
		if board[kingSquare] == BlackKing {
			return true
		}
	}

	attackedFromKnight := AttackedFromKnight[s]
	for _, knightSquare := range attackedFromKnight {
		if board[knightSquare] == BlackKnight {
			return true
		}
	}

	x := s.X()
	y := s.Y()

	for i := x - 1; i >= 0; i-- {
		peace := board[NewSquare(i, y)]
		if peace.IsBlackLinearMover() {
			return true
		}
		if peace != Empty {
			break
		}
	}
	for i := x + 1; i < 8; i++ {
		peace := board[NewSquare(i, y)]
		if peace.IsBlackLinearMover() {
			return true
		}
		if peace != Empty {
			break
		}
	}
	for i := y - 1; i >= 0; i-- {
		peace := board[NewSquare(x, i)]
		if peace.IsBlackLinearMover() {
			return true
		}
		if peace != Empty {
			break
		}
	}
	for i := y + 1; i < 8; i++ {
		peace := board[NewSquare(x, i)]
		if peace.IsBlackLinearMover() {
			return true
		}
		if peace != Empty {
			break
		}
	}

	i := x - 1
	j := y - 1
	for i >= 0 && j >= 0 {
		peace := board[NewSquare(i, j)]
		if peace.IsBlackDiagonalMover() {
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
		peace := board[NewSquare(i, j)]
		if peace.IsBlackDiagonalMover() {
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
		peace := board[NewSquare(i, j)]
		if peace.IsBlackDiagonalMover() {
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
		peace := board[NewSquare(i, j)]
		if peace.IsBlackDiagonalMover() {
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

func (board Board) BlackKingTos(s Square) []Square {
	var tos []Square

	original := board[s]
	board[s] = Empty

	kingMoves := KingMoves[s]

	for _, square := range kingMoves {
		peace := board[square]
		if peace.IsEmptyOrWhite() && !board.SquareUnderAttackFromWhite(square) {
			tos = append(tos, square)
		}
	}

	board[s] = original
	return tos
}

func (board Board) BlackRookTos(s Square, ks Square) []Square {
	var tos []Square

	check := func(square Square) bool {
		peace := board[square]
		if peace.IsEmptyOrWhite() {
			original := board[square]
			board[square] = board[s]
			board[s] = Empty

			if !board.SquareUnderAttackFromWhite(ks) {
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
		if !check(NewSquare(i, y)) {
			break
		}
	}
	for i := x + 1; i < 8; i++ {
		if !check(NewSquare(i, y)) {
			break
		}
	}

	for i := y - 1; i >= 0; i-- {
		if !check(NewSquare(x, i)) {
			break
		}
	}
	for i := y + 1; i < 8; i++ {
		if !check(NewSquare(x, i)) {
			break
		}
	}
	return tos
}

func (board Board) BlackTos(s Square, kingSquare Square) []Square {
	if board[s] == BlackKing {
		return board.BlackKingTos(s)
	}
	if board[s]== BlackRook {
		return board.BlackRookTos(s, kingSquare)
	}
	return nil
}
