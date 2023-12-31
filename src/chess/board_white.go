package chess

func (board Board) SquareUnderAttackWhite(s Square) bool {
	attackedFromKing := AttackedFromKing[s]

	for _, kingSquare := range attackedFromKing {
		if board[kingSquare] == WhiteKing {
			return true
		}
	}

	x := s.X()
	y := s.Y()

	for i := x - 1; i >= 0; i-- {
		peace := board[NewSquare(i, y)]
		if peace.IsWhiteLinearMover() {
			return true
		}
		if peace != Empty {
			break
		}
	}
	for i := x + 1; i < 8; i++ {
		peace := board[NewSquare(i, y)]
		if peace.IsWhiteLinearMover() {
			return true
		}
		if peace != Empty {
			break
		}
	}
	for i := y - 1; i >= 0; i-- {
		peace := board[NewSquare(x, i)]
		if peace.IsWhiteLinearMover() {
			return true
		}
		if peace != Empty {
			break
		}
	}
	for i := y + 1; i < 8; i++ {
		peace := board[NewSquare(x, i)]
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
		peace := board[NewSquare(i, j)]
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
		peace := board[NewSquare(i, j)]
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
		peace := board[NewSquare(i, j)]
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
		peace := board[NewSquare(i, j)]
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