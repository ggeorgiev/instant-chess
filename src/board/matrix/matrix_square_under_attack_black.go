package matrix

import (
	"github.com/ggeorgiev/instant-chess/src/attack"
	"github.com/ggeorgiev/instant-chess/src/peace"
	"github.com/ggeorgiev/instant-chess/src/square"
)

func (m *Matrix) CountSquareUnderAttackFromBlackKing(s square.Index) int {
	count := 0
	attackedFromKing := attack.FromKing[s]
	for _, kingSquare := range attackedFromKing {
		if m[kingSquare] == peace.BlackKing {
			count++
		}
	}
	return count
}

func (m *Matrix) IsSquareUnderAttackFromBlackKing(s square.Index) bool {
	attackedFromKing := attack.FromKing[s]
	for _, kingSquare := range attackedFromKing {
		if m[kingSquare] == peace.BlackKing {
			return true
		}
	}
	return false
}

func (m *Matrix) CountSquareUnderAttackFromBlackKnight(s square.Index) int {
	count := 0
	attackedFromKnight := attack.FromKnight[s]
	for _, knightSquare := range attackedFromKnight {
		if m[knightSquare] == peace.BlackKnight {
			count++
		}
	}
	return count
}

func (m *Matrix) IsSquareUnderAttackFromBlackKnight(s square.Index) bool {
	attackedFromKnight := attack.FromKnight[s]
	for _, knightSquare := range attackedFromKnight {
		if m[knightSquare] == peace.BlackKnight {
			return true
		}
	}
	return false
}

func (m *Matrix) CountSquareUnderAttackFromBlackPawn(s square.Index) int {
	count := 0
	attackedFromPawn := attack.FromBlackPawn[s]
	for _, pawnSquare := range attackedFromPawn {
		if m[pawnSquare] == peace.BlackPawn {
			count++
		}
	}
	return count
}

func (m *Matrix) IsSquareUnderAttackFromBlackPawn(s square.Index) bool {
	attackedFromPawn := attack.FromBlackPawn[s]
	for _, pawnSquare := range attackedFromPawn {
		if m[pawnSquare] == peace.BlackPawn {
			return true
		}
	}
	return false
}

func (m *Matrix) DirectAttackersOfSquareFromBlack(s square.Index) square.Indexes {
	var attackers square.Indexes
	attackedDirectly := attack.BlackDirectsList[s]
	for _, direct := range attackedDirectly {
		if m[direct.Index] == direct.Peace {
			attackers = append(attackers, direct.Index)
		}
	}
	return attackers
}

func (m *Matrix) SquareUnderDirectAttackExactlyFromBlack(s square.Index) (bool, square.Index) {
	attacker := square.InvalidIndex
	count := 0
	attackedDirectly := attack.BlackDirectsList[s]
	for _, direct := range attackedDirectly {
		if m[direct.Index] == direct.Peace {
			attacker = direct.Index
			count++
			if count > 1 {
				break
			}
		}
	}
	return count == 1, attacker
}

func (m *Matrix) IsSquareUnderDirectAttackFromBlack(s square.Index) bool {
	attackedDirectly := attack.BlackDirectsList[s]
	for _, direct := range attackedDirectly {
		if m[direct.Index] == direct.Peace {
			return true
		}
	}
	return false
}

func (m *Matrix) IsSquareUnderAttackFromBlackFromLeft(s square.Index) square.Index {
	rank := s.Rank()
	for f := s.File() - 1; f >= square.ZeroFile; f-- {
		attacker := square.NewIndex(f, rank)
		pc := m[attacker]
		if pc.IsBlackLinearMover() {
			return attacker
		}
		if pc != peace.Null {
			return square.InvalidIndex
		}
	}
	return square.InvalidIndex
}

func (m *Matrix) IsSquareUnderAttackFromBlackFromRight(s square.Index) square.Index {
	rank := s.Rank()
	for f := s.File() + 1; f <= square.LastFile; f++ {
		attacker := square.NewIndex(f, rank)
		pc := m[attacker]
		if pc.IsBlackLinearMover() {
			return attacker
		}
		if pc != peace.Null {
			return square.InvalidIndex
		}
	}
	return square.InvalidIndex
}

func (m *Matrix) IsSquareUnderAttackFromBlackFromUnder(s square.Index) square.Index {
	file := s.File()
	for r := s.Rank() - 1; r >= square.ZeroRank; r-- {
		attacker := square.NewIndex(file, r)
		pc := m[attacker]
		if pc.IsBlackLinearMover() {
			return attacker
		}
		if pc != peace.Null {
			return square.InvalidIndex
		}
	}
	return square.InvalidIndex
}

func (m *Matrix) IsSquareUnderAttackFromBlackFromAbove(s square.Index) square.Index {
	file := s.File()
	for r := s.Rank() + 1; r <= square.LastRank; r++ {
		attacker := square.NewIndex(file, r)
		pc := m[attacker]
		if pc.IsBlackLinearMover() {
			return attacker
		}
		if pc != peace.Null {
			return square.InvalidIndex
		}
	}
	return square.InvalidIndex
}

func (m *Matrix) IsSquareUnderAttackFromBlackFromLeftUnder(s square.Index) square.Index {
	f := s.File() - 1
	r := s.Rank() - 1
	for f >= square.ZeroFile && r >= square.ZeroRank {
		attacker := square.NewIndex(f, r)
		pc := m[attacker]
		if pc.IsBlackDiagonalMover() {
			return attacker
		}
		if pc != peace.Null {
			return square.InvalidIndex
		}
		f--
		r--
	}
	return square.InvalidIndex
}

func (m *Matrix) IsSquareUnderAttackFromBlackFromLeftAbove(s square.Index) square.Index {
	f := s.File() - 1
	r := s.Rank() + 1
	for f >= square.ZeroFile && r <= square.LastRank {
		attacker := square.NewIndex(f, r)
		pc := m[attacker]
		if pc.IsBlackDiagonalMover() {
			return attacker
		}
		if pc != peace.Null {
			return square.InvalidIndex
		}
		f--
		r++
	}
	return square.InvalidIndex
}

func (m *Matrix) IsSquareUnderAttackFromBlackFromRightUnder(s square.Index) square.Index {
	f := s.File() + 1
	r := s.Rank() - 1
	for f <= square.LastFile && r >= square.ZeroRank {
		attacker := square.NewIndex(f, r)
		pc := m[attacker]
		if pc.IsBlackDiagonalMover() {
			return attacker
		}
		if pc != peace.Null {
			return square.InvalidIndex
		}
		f++
		r--
	}
	return square.InvalidIndex
}

func (m *Matrix) IsSquareUnderAttackFromBlackFromRightAbove(s square.Index) square.Index {
	f := s.File() + 1
	r := s.Rank() + 1
	for f <= square.LastFile && r <= square.LastRank {
		attacker := square.NewIndex(f, r)
		pc := m[attacker]
		if pc.IsBlackDiagonalMover() {
			return attacker
		}
		if pc != peace.Null {
			return square.InvalidIndex
		}
		f++
		r++
	}
	return square.InvalidIndex
}

func (m *Matrix) IsSquareUnderAttackFromBlack(s square.Index) bool {
	if m.IsSquareUnderDirectAttackFromBlack(s) {
		return true
	}

	if m.IsSquareUnderAttackFromBlackFromLeft(s) != square.InvalidIndex {
		return true
	}

	if m.IsSquareUnderAttackFromBlackFromRight(s) != square.InvalidIndex {
		return true
	}

	if m.IsSquareUnderAttackFromBlackFromUnder(s) != square.InvalidIndex {
		return true
	}

	if m.IsSquareUnderAttackFromBlackFromAbove(s) != square.InvalidIndex {
		return true
	}

	if m.IsSquareUnderAttackFromBlackFromLeftUnder(s) != square.InvalidIndex {
		return true
	}

	if m.IsSquareUnderAttackFromBlackFromLeftAbove(s) != square.InvalidIndex {
		return true
	}

	if m.IsSquareUnderAttackFromBlackFromRightUnder(s) != square.InvalidIndex {
		return true
	}

	if m.IsSquareUnderAttackFromBlackFromRightAbove(s) != square.InvalidIndex {
		return true
	}

	return false
}

func (m *Matrix) BlockableAttackersOfSquareFromBlack(s square.Index) square.Indexes {
	var attackers square.Indexes

	attacker := m.IsSquareUnderAttackFromBlackFromLeft(s)
	if attacker != square.InvalidIndex {
		attackers = append(attackers, attacker)
	}

	attacker = m.IsSquareUnderAttackFromBlackFromRight(s)
	if attacker != square.InvalidIndex {
		attackers = append(attackers, attacker)
	}

	attacker = m.IsSquareUnderAttackFromBlackFromUnder(s)
	if attacker != square.InvalidIndex {
		attackers = append(attackers, attacker)
	}

	attacker = m.IsSquareUnderAttackFromBlackFromAbove(s)
	if attacker != square.InvalidIndex {
		attackers = append(attackers, attacker)
	}

	attacker = m.IsSquareUnderAttackFromBlackFromLeftUnder(s)
	if attacker != square.InvalidIndex {
		attackers = append(attackers, attacker)
	}

	attacker = m.IsSquareUnderAttackFromBlackFromLeftAbove(s)
	if attacker != square.InvalidIndex {
		attackers = append(attackers, attacker)
	}

	attacker = m.IsSquareUnderAttackFromBlackFromRightUnder(s)
	if attacker != square.InvalidIndex {
		attackers = append(attackers, attacker)
	}

	attacker = m.IsSquareUnderAttackFromBlackFromRightAbove(s)
	if attacker != square.InvalidIndex {
		attackers = append(attackers, attacker)
	}

	return attackers
}

func (m *Matrix) AttackersOfSquareFromBlack(s square.Index) square.Indexes {
	var attackers square.Indexes
	attackers = append(attackers, m.DirectAttackersOfSquareFromBlack(s)...)
	attackers = append(attackers, m.BlockableAttackersOfSquareFromBlack(s)...)
	return attackers
}
