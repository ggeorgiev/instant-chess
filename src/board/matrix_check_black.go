package board

import (
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

func (m Matrix) IsBlackCheckedAfterMove(kingSquare square.Index, movedFrom square.Index) bool {
	// TODO: implement optimized
	return m.IsBlackChecked(kingSquare)
}
