package board

import (
	"github.com/ggeorgiev/instant-chess/src/peacealignment"
	"github.com/ggeorgiev/instant-chess/src/square"
)

func (m Matrix) IsWhiteChecked(kingSquare square.Index) bool {
	if m.IsSquareUnderDirectAttackFromBlack(kingSquare) {
		return true
	}

	if m.IsSquareUnderAttackFromBlackFromLeft(kingSquare) {
		return true
	}

	if m.IsSquareUnderAttackFromBlackFromRight(kingSquare) {
		return true
	}

	if m.IsSquareUnderAttackFromBlackFromUnder(kingSquare) {
		return true
	}

	if m.IsSquareUnderAttackFromBlackFromAbove(kingSquare) {
		return true
	}

	if m.IsSquareUnderAttackFromBlackFromLeftUnder(kingSquare) {
		return true
	}

	if m.IsSquareUnderAttackFromBlackFromLeftAbove(kingSquare) {
		return true
	}

	if m.IsSquareUnderAttackFromBlackFromRightUnder(kingSquare) {
		return true
	}

	if m.IsSquareUnderAttackFromBlackFromRightAbove(kingSquare) {
		return true
	}

	return false
}

func (m Matrix) IsWhiteCheckedToMove(kingSquare square.Index) (bool, bool) {
	if m.IsSquareUnderDirectAttackFromBlack(kingSquare) {
		return true, true
	}

	checked := false
	if m.IsSquareUnderAttackFromBlackFromLeft(kingSquare) {
		if checked {
			return true, true
		}
		checked = true
	}

	if m.IsSquareUnderAttackFromBlackFromRight(kingSquare) {
		if checked {
			return true, true
		}
		checked = true
	}

	if m.IsSquareUnderAttackFromBlackFromUnder(kingSquare) {
		if checked {
			return true, true
		}
		checked = true
	}

	if m.IsSquareUnderAttackFromBlackFromAbove(kingSquare) {
		if checked {
			return true, true
		}
		checked = true
	}

	if m.IsSquareUnderAttackFromBlackFromLeftUnder(kingSquare) {
		if checked {
			return true, true
		}
		checked = true
	}

	if m.IsSquareUnderAttackFromBlackFromLeftAbove(kingSquare) {
		if checked {
			return true, true
		}
		checked = true
	}

	if m.IsSquareUnderAttackFromBlackFromRightUnder(kingSquare) {
		if checked {
			return true, true
		}
		checked = true
	}

	if m.IsSquareUnderAttackFromBlackFromRightAbove(kingSquare) {
		if checked {
			return true, true
		}
		checked = true
	}

	return checked, false
}

func (m Matrix) IsWhiteMaybeCheckedAfterMove(kingSquare square.Index, movedFrom square.Index) peacealignment.Vector {
	sq := peacealignment.SquareRelations[kingSquare][movedFrom]
	if sq == peacealignment.NotAligned {
		return peacealignment.NoVector
	}

	if sq == peacealignment.RankLeft {
		if m.IsSquareUnderAttackFromBlackFromLeft(kingSquare) {
			return peacealignment.Rank
		}
	}

	if sq == peacealignment.RankRight {
		if m.IsSquareUnderAttackFromBlackFromRight(kingSquare) {
			return peacealignment.Rank
		}
	}

	if sq == peacealignment.FileUnder {
		if m.IsSquareUnderAttackFromBlackFromUnder(kingSquare) {
			return peacealignment.File
		}
	}

	if sq == peacealignment.FileAbove {
		if m.IsSquareUnderAttackFromBlackFromAbove(kingSquare) {
			return peacealignment.File
		}
	}

	if sq == peacealignment.DiagonalUnder {
		if m.IsSquareUnderAttackFromBlackFromLeftUnder(kingSquare) {
			return peacealignment.Diagonal
		}
	}

	if sq == peacealignment.DiagonalAbove {
		if m.IsSquareUnderAttackFromBlackFromLeftAbove(kingSquare) {
			return peacealignment.Diagonal
		}
	}

	if sq == peacealignment.CounterDiagonalAbove {
		if m.IsSquareUnderAttackFromBlackFromRightUnder(kingSquare) {
			return peacealignment.CounterDiagonal
		}
	}

	if sq == peacealignment.CounterDiagonalUnder {
		if m.IsSquareUnderAttackFromBlackFromRightAbove(kingSquare) {
			return peacealignment.CounterDiagonal
		}
	}

	return peacealignment.NoVector
}
