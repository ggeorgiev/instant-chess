package matrix

import (
	"github.com/ggeorgiev/instant-chess/src/peacealignment"
	"github.com/ggeorgiev/instant-chess/src/square"
)

func (m *Matrix) IsBlackChecked(kingSquare square.Index) bool {
	if m.IsSquareUnderDirectAttackFromWhite(kingSquare) {
		return true
	}

	if m.IsSquareUnderAttackFromWhiteFromLeft(kingSquare) != square.InvalidIndex {
		return true
	}

	if m.IsSquareUnderAttackFromWhiteFromRight(kingSquare) != square.InvalidIndex {
		return true
	}

	if m.IsSquareUnderAttackFromWhiteFromUnder(kingSquare) != square.InvalidIndex {
		return true
	}

	if m.IsSquareUnderAttackFromWhiteFromAbove(kingSquare) != square.InvalidIndex {
		return true
	}

	if m.IsSquareUnderAttackFromWhiteFromLeftUnder(kingSquare) != square.InvalidIndex {
		return true
	}

	if m.IsSquareUnderAttackFromWhiteFromLeftAbove(kingSquare) != square.InvalidIndex {
		return true
	}

	if m.IsSquareUnderAttackFromWhiteFromRightUnder(kingSquare) != square.InvalidIndex {
		return true
	}

	if m.IsSquareUnderAttackFromWhiteFromRightAbove(kingSquare) != square.InvalidIndex {
		return true
	}

	return false
}

func (m *Matrix) IsBlackCheckedToMoveCaptureOrBlock(kingSquare square.Index) (bool, square.Index, bool) {
	checked, attacker := m.SquareUnderDirectAttackExactlyFromWhite(kingSquare)
	if checked && attacker != square.InvalidIndex {
		return true, attacker, false
	}

	evaluateAttacker := m.IsSquareUnderAttackFromWhiteFromLeft(kingSquare)
	if evaluateAttacker != square.InvalidIndex {
		if attacker != square.InvalidIndex {
			return true, square.InvalidIndex, false
		}
		attacker = evaluateAttacker
	}

	evaluateAttacker = m.IsSquareUnderAttackFromWhiteFromRight(kingSquare)
	if evaluateAttacker != square.InvalidIndex {
		if attacker != square.InvalidIndex {
			return true, square.InvalidIndex, false
		}
		attacker = evaluateAttacker
	}

	evaluateAttacker = m.IsSquareUnderAttackFromWhiteFromUnder(kingSquare)
	if evaluateAttacker != square.InvalidIndex {
		if attacker != square.InvalidIndex {
			return true, square.InvalidIndex, false
		}
		attacker = evaluateAttacker
	}

	evaluateAttacker = m.IsSquareUnderAttackFromWhiteFromAbove(kingSquare)
	if evaluateAttacker != square.InvalidIndex {
		if attacker != square.InvalidIndex {
			return true, square.InvalidIndex, false
		}
		attacker = evaluateAttacker
	}

	evaluateAttacker = m.IsSquareUnderAttackFromWhiteFromLeftUnder(kingSquare)
	if evaluateAttacker != square.InvalidIndex {
		if attacker != square.InvalidIndex {
			return true, square.InvalidIndex, false
		}
		attacker = evaluateAttacker
	}

	evaluateAttacker = m.IsSquareUnderAttackFromWhiteFromLeftAbove(kingSquare)
	if evaluateAttacker != square.InvalidIndex {
		if attacker != square.InvalidIndex {
			return true, square.InvalidIndex, false
		}
		attacker = evaluateAttacker
	}

	evaluateAttacker = m.IsSquareUnderAttackFromWhiteFromRightUnder(kingSquare)
	if evaluateAttacker != square.InvalidIndex {
		if attacker != square.InvalidIndex {
			return true, square.InvalidIndex, false
		}
		attacker = evaluateAttacker
	}

	evaluateAttacker = m.IsSquareUnderAttackFromWhiteFromRightAbove(kingSquare)
	if evaluateAttacker != square.InvalidIndex {
		if attacker != square.InvalidIndex {
			return true, square.InvalidIndex, false
		}
		attacker = evaluateAttacker
	}

	return attacker != square.InvalidIndex, attacker, true
}

func (m *Matrix) IsBlackMaybeCheckedAfterMove(kingSquare square.Index, movedFrom square.Index) peacealignment.Vector {
	sq := peacealignment.SquareRelations[kingSquare][movedFrom]
	if sq == peacealignment.NotAligned {
		return peacealignment.NoVector
	}

	if sq == peacealignment.RankLeft {
		if m.IsSquareUnderAttackFromWhiteFromLeft(kingSquare) != square.InvalidIndex {
			return peacealignment.Rank
		}
	}

	if sq == peacealignment.RankRight {
		if m.IsSquareUnderAttackFromWhiteFromRight(kingSquare) != square.InvalidIndex {
			return peacealignment.Rank
		}
	}

	if sq == peacealignment.FileUnder {
		if m.IsSquareUnderAttackFromWhiteFromUnder(kingSquare) != square.InvalidIndex {
			return peacealignment.File
		}
	}

	if sq == peacealignment.FileAbove {
		if m.IsSquareUnderAttackFromWhiteFromAbove(kingSquare) != square.InvalidIndex {
			return peacealignment.File
		}
	}

	if sq == peacealignment.DiagonalUnder {
		if m.IsSquareUnderAttackFromWhiteFromLeftUnder(kingSquare) != square.InvalidIndex {
			return peacealignment.Diagonal
		}
	}

	if sq == peacealignment.DiagonalAbove {
		if m.IsSquareUnderAttackFromWhiteFromRightAbove(kingSquare) != square.InvalidIndex {
			return peacealignment.Diagonal
		}
	}

	if sq == peacealignment.CounterDiagonalUnder {
		if m.IsSquareUnderAttackFromWhiteFromRightUnder(kingSquare) != square.InvalidIndex {
			return peacealignment.CounterDiagonal
		}
	}

	if sq == peacealignment.CounterDiagonalAbove {
		if m.IsSquareUnderAttackFromWhiteFromLeftAbove(kingSquare) != square.InvalidIndex {
			return peacealignment.CounterDiagonal
		}
	}

	return peacealignment.NoVector
}
