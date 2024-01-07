package board

import (
	"github.com/ggeorgiev/instant-chess/src/peacealignment"
	"github.com/ggeorgiev/instant-chess/src/square"
)

func (m Matrix) IsBlackChecked(kingSquare square.Index) bool {
	if m.IsSquareUnderDirectAttackFromWhite(kingSquare) {
		return true
	}

	if m.IsSquareUnderAttackFromWhiteFromLeft(kingSquare) {
		return true
	}

	if m.IsSquareUnderAttackFromWhiteFromRight(kingSquare) {
		return true
	}

	if m.IsSquareUnderAttackFromWhiteFromUnder(kingSquare) {
		return true
	}

	if m.IsSquareUnderAttackFromWhiteFromAbove(kingSquare) {
		return true
	}

	if m.IsSquareUnderAttackFromWhiteFromLeftUnder(kingSquare) {
		return true
	}

	if m.IsSquareUnderAttackFromWhiteFromLeftAbove(kingSquare) {
		return true
	}

	if m.IsSquareUnderAttackFromWhiteFromRightUnder(kingSquare) {
		return true
	}

	if m.IsSquareUnderAttackFromWhiteFromRightAbove(kingSquare) {
		return true
	}

	return false
}

func (m Matrix) IsBlackCheckedToMove(kingSquare square.Index) (bool, bool) {
	if m.IsSquareUnderDirectAttackFromWhite(kingSquare) {
		return true, true
	}

	checked := false
	if m.IsSquareUnderAttackFromWhiteFromLeft(kingSquare) {
		if checked {
			return true, true
		}
		checked = true
	}

	if m.IsSquareUnderAttackFromWhiteFromRight(kingSquare) {
		if checked {
			return true, true
		}
		checked = true
	}

	if m.IsSquareUnderAttackFromWhiteFromUnder(kingSquare) {
		if checked {
			return true, true
		}
		checked = true
	}

	if m.IsSquareUnderAttackFromWhiteFromAbove(kingSquare) {
		if checked {
			return true, true
		}
		checked = true
	}

	if m.IsSquareUnderAttackFromWhiteFromLeftUnder(kingSquare) {
		if checked {
			return true, true
		}
		checked = true
	}

	if m.IsSquareUnderAttackFromWhiteFromLeftAbove(kingSquare) {
		if checked {
			return true, true
		}
		checked = true
	}

	if m.IsSquareUnderAttackFromWhiteFromRightUnder(kingSquare) {
		if checked {
			return true, true
		}
		checked = true
	}

	if m.IsSquareUnderAttackFromWhiteFromRightAbove(kingSquare) {
		if checked {
			return true, true
		}
		checked = true
	}

	return checked, false
}

func (m Matrix) IsBlackMaybeCheckedAfterMove(kingSquare square.Index, movedFrom square.Index) (bool, peacealignment.Relation) {
	sq := peacealignment.SquareRelations[kingSquare][movedFrom]
	if sq == peacealignment.NotAligned {
		return false, sq
	}

	if sq == peacealignment.RankLeft {
		if m.IsSquareUnderAttackFromWhiteFromLeft(kingSquare) {
			return true, sq
		}
	}

	if sq == peacealignment.RankRight {
		if m.IsSquareUnderAttackFromWhiteFromRight(kingSquare) {
			return true, sq
		}
	}

	if sq == peacealignment.FileUnder {
		if m.IsSquareUnderAttackFromWhiteFromUnder(kingSquare) {
			return true, sq
		}
	}

	if sq == peacealignment.FileAbove {
		if m.IsSquareUnderAttackFromWhiteFromAbove(kingSquare) {
			return true, sq
		}
	}

	if sq == peacealignment.DiagonalUnder {
		if m.IsSquareUnderAttackFromWhiteFromLeftUnder(kingSquare) {
			return true, sq
		}
	}

	if sq == peacealignment.DiagonalAbove {
		if m.IsSquareUnderAttackFromWhiteFromLeftAbove(kingSquare) {
			return true, sq
		}
	}

	if sq == peacealignment.CounterDiagonalAbove {
		if m.IsSquareUnderAttackFromWhiteFromRightUnder(kingSquare) {
			return true, sq
		}
	}

	if sq == peacealignment.CounterDiagonalUnder {
		if m.IsSquareUnderAttackFromWhiteFromRightAbove(kingSquare) {
			return true, sq
		}
	}

	return false, sq
}
