package matrix

import (
	"github.com/ggeorgiev/instant-chess/src/peacealignment"
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

func (m *Matrix) IsWhiteMaybeCheckedAfterMove(kingSquare square.Index, movedFrom square.Index) peacealignment.Vector {
	sq := peacealignment.SquareRelations[kingSquare][movedFrom]
	if sq == peacealignment.NotAligned {
		return peacealignment.NoVector
	}

	if sq == peacealignment.RankLeft {
		if m.IsSquareUnderAttackFromBlackFromLeft(kingSquare) != square.InvalidIndex {
			return peacealignment.Rank
		}
	}

	if sq == peacealignment.RankRight {
		if m.IsSquareUnderAttackFromBlackFromRight(kingSquare) != square.InvalidIndex {
			return peacealignment.Rank
		}
	}

	if sq == peacealignment.FileUnder {
		if m.IsSquareUnderAttackFromBlackFromUnder(kingSquare) != square.InvalidIndex {
			return peacealignment.File
		}
	}

	if sq == peacealignment.FileAbove {
		if m.IsSquareUnderAttackFromBlackFromAbove(kingSquare) != square.InvalidIndex {
			return peacealignment.File
		}
	}

	if sq == peacealignment.DiagonalUnder {
		if m.IsSquareUnderAttackFromBlackFromLeftUnder(kingSquare) != square.InvalidIndex {
			return peacealignment.Diagonal
		}
	}

	if sq == peacealignment.DiagonalAbove {
		if m.IsSquareUnderAttackFromBlackFromRightAbove(kingSquare) != square.InvalidIndex {
			return peacealignment.Diagonal
		}
	}

	if sq == peacealignment.CounterDiagonalUnder {
		if m.IsSquareUnderAttackFromBlackFromRightUnder(kingSquare) != square.InvalidIndex {
			return peacealignment.CounterDiagonal
		}
	}

	if sq == peacealignment.CounterDiagonalAbove {
		if m.IsSquareUnderAttackFromBlackFromLeftAbove(kingSquare) != square.InvalidIndex {
			return peacealignment.CounterDiagonal
		}
	}

	return peacealignment.NoVector
}
