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
	return false
}