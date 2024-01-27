package matrix

import (
	"github.com/ggeorgiev/instant-chess/src/alignment"
	"github.com/ggeorgiev/instant-chess/src/square"
)

func (m *Matrix) IsWhiteChecked(kingSquare square.Index) bool {
	if m.IsSquareUnderDirectAttackFromBlack(kingSquare) {
		return true
	}

	if m.IsSquareUnderAttackFromBlackFromLeft(kingSquare) != square.InvalidIndex {
		return true
	}

	if m.IsSquareUnderAttackFromBlackFromRight(kingSquare) != square.InvalidIndex {
		return true
	}

	if m.IsSquareUnderAttackFromBlackFromUnder(kingSquare) != square.InvalidIndex {
		return true
	}

	if m.IsSquareUnderAttackFromBlackFromAbove(kingSquare) != square.InvalidIndex {
		return true
	}

	if m.IsSquareUnderAttackFromBlackFromLeftUnder(kingSquare) != square.InvalidIndex {
		return true
	}

	if m.IsSquareUnderAttackFromBlackFromLeftAbove(kingSquare) != square.InvalidIndex {
		return true
	}

	if m.IsSquareUnderAttackFromBlackFromRightUnder(kingSquare) != square.InvalidIndex {
		return true
	}

	if m.IsSquareUnderAttackFromBlackFromRightAbove(kingSquare) != square.InvalidIndex {
		return true
	}

	return false
}

func (m *Matrix) IsWhiteCheckedToMoveCaptureOrBlock(kingSquare square.Index) (bool, square.Index, bool) {
	checked, attacker := m.SquareUnderDirectAttackExactlyFromBlack(kingSquare)
	if checked && attacker != square.InvalidIndex {
		return true, attacker, false
	}

	evaluateAttacker := m.IsSquareUnderAttackFromBlackFromLeft(kingSquare)
	if evaluateAttacker != square.InvalidIndex {
		if attacker != square.InvalidIndex {
			return true, square.InvalidIndex, false
		}
		attacker = evaluateAttacker
	}

	evaluateAttacker = m.IsSquareUnderAttackFromBlackFromRight(kingSquare)
	if evaluateAttacker != square.InvalidIndex {
		if attacker != square.InvalidIndex {
			return true, square.InvalidIndex, false
		}
		attacker = evaluateAttacker
	}

	evaluateAttacker = m.IsSquareUnderAttackFromBlackFromUnder(kingSquare)
	if evaluateAttacker != square.InvalidIndex {
		if attacker != square.InvalidIndex {
			return true, square.InvalidIndex, false
		}
		attacker = evaluateAttacker
	}

	evaluateAttacker = m.IsSquareUnderAttackFromBlackFromAbove(kingSquare)
	if evaluateAttacker != square.InvalidIndex {
		if attacker != square.InvalidIndex {
			return true, square.InvalidIndex, false
		}
		attacker = evaluateAttacker
	}

	evaluateAttacker = m.IsSquareUnderAttackFromBlackFromLeftUnder(kingSquare)
	if evaluateAttacker != square.InvalidIndex {
		if attacker != square.InvalidIndex {
			return true, square.InvalidIndex, false
		}
		attacker = evaluateAttacker
	}

	evaluateAttacker = m.IsSquareUnderAttackFromBlackFromLeftAbove(kingSquare)
	if evaluateAttacker != square.InvalidIndex {
		if attacker != square.InvalidIndex {
			return true, square.InvalidIndex, false
		}
		attacker = evaluateAttacker
	}

	evaluateAttacker = m.IsSquareUnderAttackFromBlackFromRightUnder(kingSquare)
	if evaluateAttacker != square.InvalidIndex {
		if attacker != square.InvalidIndex {
			return true, square.InvalidIndex, false
		}
		attacker = evaluateAttacker
	}

	evaluateAttacker = m.IsSquareUnderAttackFromBlackFromRightAbove(kingSquare)
	if evaluateAttacker != square.InvalidIndex {
		if attacker != square.InvalidIndex {
			return true, square.InvalidIndex, false
		}
		attacker = evaluateAttacker
	}

	return attacker != square.InvalidIndex, attacker, true
}

func (m *Matrix) IsWhiteMaybeCheckedAfterMove(kingSquare square.Index, movedFrom square.Index) alignment.Vector {
	sq := alignment.SquareRelations[kingSquare][movedFrom]
	if sq == alignment.NotAligned {
		return alignment.NoVector
	}

	if sq == alignment.RankLeft {
		if m.IsSquareUnderAttackFromBlackFromLeft(kingSquare) != square.InvalidIndex {
			return alignment.Rank
		}
	}

	if sq == alignment.RankRight {
		if m.IsSquareUnderAttackFromBlackFromRight(kingSquare) != square.InvalidIndex {
			return alignment.Rank
		}
	}

	if sq == alignment.FileUnder {
		if m.IsSquareUnderAttackFromBlackFromUnder(kingSquare) != square.InvalidIndex {
			return alignment.File
		}
	}

	if sq == alignment.FileAbove {
		if m.IsSquareUnderAttackFromBlackFromAbove(kingSquare) != square.InvalidIndex {
			return alignment.File
		}
	}

	if sq == alignment.DiagonalUnder {
		if m.IsSquareUnderAttackFromBlackFromLeftUnder(kingSquare) != square.InvalidIndex {
			return alignment.Diagonal
		}
	}

	if sq == alignment.DiagonalAbove {
		if m.IsSquareUnderAttackFromBlackFromRightAbove(kingSquare) != square.InvalidIndex {
			return alignment.Diagonal
		}
	}

	if sq == alignment.CounterDiagonalUnder {
		if m.IsSquareUnderAttackFromBlackFromRightUnder(kingSquare) != square.InvalidIndex {
			return alignment.CounterDiagonal
		}
	}

	if sq == alignment.CounterDiagonalAbove {
		if m.IsSquareUnderAttackFromBlackFromLeftAbove(kingSquare) != square.InvalidIndex {
			return alignment.CounterDiagonal
		}
	}

	return alignment.NoVector
}
