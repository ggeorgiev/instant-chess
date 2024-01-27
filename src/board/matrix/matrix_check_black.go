package matrix

import (
	"github.com/ggeorgiev/instant-chess/src/alignment"
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

func (m *Matrix) IsBlackMaybeCheckedAfterMove(kingSquare square.Index, movedFrom square.Index) alignment.Vector {
	sq := alignment.SquareRelations[kingSquare][movedFrom]
	if sq == alignment.NotAligned {
		return alignment.NoVector
	}

	if sq == alignment.RankLeft {
		if m.IsSquareUnderAttackFromWhiteFromLeft(kingSquare) != square.InvalidIndex {
			return alignment.Rank
		}
	}

	if sq == alignment.RankRight {
		if m.IsSquareUnderAttackFromWhiteFromRight(kingSquare) != square.InvalidIndex {
			return alignment.Rank
		}
	}

	if sq == alignment.FileUnder {
		if m.IsSquareUnderAttackFromWhiteFromUnder(kingSquare) != square.InvalidIndex {
			return alignment.File
		}
	}

	if sq == alignment.FileAbove {
		if m.IsSquareUnderAttackFromWhiteFromAbove(kingSquare) != square.InvalidIndex {
			return alignment.File
		}
	}

	if sq == alignment.DiagonalUnder {
		if m.IsSquareUnderAttackFromWhiteFromLeftUnder(kingSquare) != square.InvalidIndex {
			return alignment.Diagonal
		}
	}

	if sq == alignment.DiagonalAbove {
		if m.IsSquareUnderAttackFromWhiteFromRightAbove(kingSquare) != square.InvalidIndex {
			return alignment.Diagonal
		}
	}

	if sq == alignment.CounterDiagonalUnder {
		if m.IsSquareUnderAttackFromWhiteFromRightUnder(kingSquare) != square.InvalidIndex {
			return alignment.CounterDiagonal
		}
	}

	if sq == alignment.CounterDiagonalAbove {
		if m.IsSquareUnderAttackFromWhiteFromLeftAbove(kingSquare) != square.InvalidIndex {
			return alignment.CounterDiagonal
		}
	}

	return alignment.NoVector
}
