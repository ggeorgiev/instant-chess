package board

import (
	"github.com/ggeorgiev/instant-chess/src/peace"
	"github.com/ggeorgiev/instant-chess/src/peaceattacks"
	"github.com/ggeorgiev/instant-chess/src/square"
)

func (m *Matrix) CountSquareUnderAttackFromWhiteKing(s square.Index) int {
	count := 0
	attackedFromKing := peaceattacks.FromKing[s]
	for _, kingSquare := range attackedFromKing {
		if m[kingSquare] == peace.WhiteKing {
			count++
		}
	}
	return count
}

func (m *Matrix) IsSquareUnderAttackFromWhiteKing(s square.Index) bool {
	attackedFromKing := peaceattacks.FromKing[s]
	for _, kingSquare := range attackedFromKing {
		if m[kingSquare] == peace.WhiteKing {
			return true
		}
	}
	return false
}

func (m *Matrix) CountSquareUnderAttackFromWhiteKnight(s square.Index) int {
	count := 0
	attackedFromKnight := peaceattacks.FromKnight[s]
	for _, knightSquare := range attackedFromKnight {
		if m[knightSquare] == peace.WhiteKnight {
			count++
		}
	}
	return count
}

func (m *Matrix) IsSquareUnderAttackFromWhiteKnight(s square.Index) bool {
	attackedFromKnight := peaceattacks.FromKnight[s]
	for _, knightSquare := range attackedFromKnight {
		if m[knightSquare] == peace.WhiteKnight {
			return true
		}
	}
	return false
}

func (m *Matrix) CountSquareUnderAttackFromWhitePawn(s square.Index) int {
	count := 0
	attackedFromPawn := peaceattacks.FromWhitePawn[s]
	for _, pawnSquare := range attackedFromPawn {
		if m[pawnSquare] == peace.WhitePawn {
			count++
		}
	}
	return count
}

func (m *Matrix) IsSquareUnderAttackFromWhitePawn(s square.Index) bool {
	attackedFromPawn := peaceattacks.FromWhitePawn[s]
	for _, pawnSquare := range attackedFromPawn {
		if m[pawnSquare] == peace.WhitePawn {
			return true
		}
	}
	return false
}

func (m *Matrix) DirectAttackersOfSquareFromWhite(s square.Index) square.Indexes {
	var attackers square.Indexes
	attackedDirectly := peaceattacks.WhiteDirectsList[s]
	for _, direct := range attackedDirectly {
		if m[direct.Index] == direct.Peace {
			attackers = append(attackers, direct.Index)
		}
	}
	return attackers
}

func (m *Matrix) SquareUnderDirectAttackExactlyFromWhite(s square.Index) (bool, square.Index) {
	attacker := square.InvalidIndex
	count := 0
	attackedDirectly := peaceattacks.WhiteDirectsList[s]
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

func (m *Matrix) IsSquareUnderDirectAttackFromWhite(s square.Index) bool {
	attackedDirectly := peaceattacks.WhiteDirectsList[s]
	for _, direct := range attackedDirectly {
		if m[direct.Index] == direct.Peace {
			return true
		}
	}
	return false
}

func (m *Matrix) IsSquareUnderAttackFromWhiteFromLeft(s square.Index) square.Index {
	rank := s.Rank()
	for f := s.File() - 1; f >= square.ZeroFile; f-- {
		attacker := square.NewIndex(f, rank)
		figure := m[attacker]
		if figure.IsWhiteLinearMover() {
			return attacker
		}
		if figure != peace.NoFigure {
			return square.InvalidIndex
		}
	}
	return square.InvalidIndex
}

func (m *Matrix) IsSquareUnderAttackFromWhiteFromRight(s square.Index) square.Index {
	rank := s.Rank()
	for f := s.File() + 1; f <= square.LastFile; f++ {
		attacker := square.NewIndex(f, rank)
		figure := m[attacker]
		if figure.IsWhiteLinearMover() {
			return attacker
		}
		if figure != peace.NoFigure {
			return square.InvalidIndex
		}
	}
	return square.InvalidIndex
}

func (m *Matrix) IsSquareUnderAttackFromWhiteFromUnder(s square.Index) square.Index {
	file := s.File()
	for r := s.Rank() - 1; r >= square.ZeroRank; r-- {
		attacker := square.NewIndex(file, r)
		figure := m[attacker]
		if figure.IsWhiteLinearMover() {
			return attacker
		}
		if figure != peace.NoFigure {
			return square.InvalidIndex
		}
	}
	return square.InvalidIndex
}

func (m *Matrix) IsSquareUnderAttackFromWhiteFromAbove(s square.Index) square.Index {
	file := s.File()
	for r := s.Rank() + 1; r <= square.LastRank; r++ {
		attacker := square.NewIndex(file, r)
		figure := m[attacker]
		if figure.IsWhiteLinearMover() {
			return attacker
		}
		if figure != peace.NoFigure {
			return square.InvalidIndex
		}
	}
	return square.InvalidIndex
}

func (m *Matrix) IsSquareUnderAttackFromWhiteFromLeftUnder(s square.Index) square.Index {
	f := s.File() - 1
	r := s.Rank() - 1
	for f >= square.ZeroFile && r >= square.ZeroRank {
		attacker := square.NewIndex(f, r)
		figure := m[attacker]
		if figure.IsWhiteDiagonalMover() {
			return attacker
		}
		if figure != peace.NoFigure {
			return square.InvalidIndex
		}
		f--
		r--
	}
	return square.InvalidIndex
}

func (m *Matrix) IsSquareUnderAttackFromWhiteFromLeftAbove(s square.Index) square.Index {
	f := s.File() - 1
	r := s.Rank() + 1
	for f >= square.ZeroFile && r <= square.LastRank {
		attacker := square.NewIndex(f, r)
		figure := m[attacker]
		if figure.IsWhiteDiagonalMover() {
			return attacker
		}
		if figure != peace.NoFigure {
			return square.InvalidIndex
		}
		f--
		r++
	}
	return square.InvalidIndex
}

func (m *Matrix) IsSquareUnderAttackFromWhiteFromRightUnder(s square.Index) square.Index {
	f := s.File() + 1
	r := s.Rank() - 1
	for f <= square.LastFile && r >= square.ZeroRank {
		attacker := square.NewIndex(f, r)
		figure := m[attacker]
		if figure.IsWhiteDiagonalMover() {
			return attacker
		}
		if figure != peace.NoFigure {
			return square.InvalidIndex
		}
		f++
		r--
	}
	return square.InvalidIndex
}

func (m *Matrix) IsSquareUnderAttackFromWhiteFromRightAbove(s square.Index) square.Index {
	f := s.File() + 1
	r := s.Rank() + 1
	for f <= square.LastFile && r <= square.LastRank {
		attacker := square.NewIndex(f, r)
		figure := m[attacker]
		if figure.IsWhiteDiagonalMover() {
			return attacker
		}
		if figure != peace.NoFigure {
			return square.InvalidIndex
		}
		f++
		r++
	}
	return square.InvalidIndex
}

func (m *Matrix) IsSquareUnderAttackFromWhite(s square.Index) bool {
	if m.IsSquareUnderDirectAttackFromWhite(s) {
		return true
	}

	if m.IsSquareUnderAttackFromWhiteFromLeft(s) != square.InvalidIndex {
		return true
	}

	if m.IsSquareUnderAttackFromWhiteFromRight(s) != square.InvalidIndex {
		return true
	}

	if m.IsSquareUnderAttackFromWhiteFromUnder(s) != square.InvalidIndex {
		return true
	}

	if m.IsSquareUnderAttackFromWhiteFromAbove(s) != square.InvalidIndex {
		return true
	}

	if m.IsSquareUnderAttackFromWhiteFromLeftUnder(s) != square.InvalidIndex {
		return true
	}

	if m.IsSquareUnderAttackFromWhiteFromLeftAbove(s) != square.InvalidIndex {
		return true
	}

	if m.IsSquareUnderAttackFromWhiteFromRightUnder(s) != square.InvalidIndex {
		return true
	}

	if m.IsSquareUnderAttackFromWhiteFromRightAbove(s) != square.InvalidIndex {
		return true
	}

	return false
}

func (m *Matrix) BlockableAttackersOfSquareFromWhite(s square.Index) square.Indexes {
	var attackers square.Indexes

	attacker := m.IsSquareUnderAttackFromWhiteFromLeft(s)
	if attacker != square.InvalidIndex {
		attackers = append(attackers, attacker)
	}

	attacker = m.IsSquareUnderAttackFromWhiteFromRight(s)
	if attacker != square.InvalidIndex {
		attackers = append(attackers, attacker)
	}

	attacker = m.IsSquareUnderAttackFromWhiteFromUnder(s)
	if attacker != square.InvalidIndex {
		attackers = append(attackers, attacker)
	}

	attacker = m.IsSquareUnderAttackFromWhiteFromAbove(s)
	if attacker != square.InvalidIndex {
		attackers = append(attackers, attacker)
	}

	attacker = m.IsSquareUnderAttackFromWhiteFromLeftUnder(s)
	if attacker != square.InvalidIndex {
		attackers = append(attackers, attacker)
	}

	attacker = m.IsSquareUnderAttackFromWhiteFromLeftAbove(s)
	if attacker != square.InvalidIndex {
		attackers = append(attackers, attacker)
	}

	attacker = m.IsSquareUnderAttackFromWhiteFromRightUnder(s)
	if attacker != square.InvalidIndex {
		attackers = append(attackers, attacker)
	}

	attacker = m.IsSquareUnderAttackFromWhiteFromRightAbove(s)
	if attacker != square.InvalidIndex {
		attackers = append(attackers, attacker)
	}

	return attackers
}

func (m *Matrix) AttackersOfSquareFromWhite(s square.Index) square.Indexes {
	var attackers square.Indexes
	attackers = append(attackers, m.DirectAttackersOfSquareFromWhite(s)...)
	attackers = append(attackers, m.BlockableAttackersOfSquareFromWhite(s)...)
	return attackers
}
