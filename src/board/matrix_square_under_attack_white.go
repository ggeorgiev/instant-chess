package board

import (
	"github.com/ggeorgiev/instant-chess/src/peace"
	"github.com/ggeorgiev/instant-chess/src/peaceattacks"
	"github.com/ggeorgiev/instant-chess/src/square"
)

func (m Matrix) CountSquareUnderAttackFromWhiteKing(s square.Index) int {
	count := 0
	attackedFromKing := peaceattacks.FromKing[s]
	for _, kingSquare := range attackedFromKing {
		if m[kingSquare] == peace.WhiteKing {
			count++
		}
	}
	return count
}

func (m Matrix) IsSquareUnderAttackFromWhiteKing(s square.Index) bool {
	attackedFromKing := peaceattacks.FromKing[s]
	for _, kingSquare := range attackedFromKing {
		if m[kingSquare] == peace.WhiteKing {
			return true
		}
	}
	return false
}

func (m Matrix) CountSquareUnderAttackFromWhiteKnight(s square.Index) int {
	count := 0
	attackedFromKnight := peaceattacks.FromKnight[s]
	for _, knightSquare := range attackedFromKnight {
		if m[knightSquare] == peace.WhiteKnight {
			count++
		}
	}
	return count
}

func (m Matrix) IsSquareUnderAttackFromWhiteKnight(s square.Index) bool {
	attackedFromKnight := peaceattacks.FromKnight[s]
	for _, knightSquare := range attackedFromKnight {
		if m[knightSquare] == peace.WhiteKnight {
			return true
		}
	}
	return false
}

func (m Matrix) CountSquareUnderAttackFromWhitePawn(s square.Index) int {
	count := 0
	attackedFromPawn := peaceattacks.FromWhitePawn[s]
	for _, pawnSquare := range attackedFromPawn {
		if m[pawnSquare] == peace.WhitePawn {
			count++
		}
	}
	return count
}

func (m Matrix) IsSquareUnderAttackFromWhitePawn(s square.Index) bool {
	attackedFromPawn := peaceattacks.FromWhitePawn[s]
	for _, pawnSquare := range attackedFromPawn {
		if m[pawnSquare] == peace.WhitePawn {
			return true
		}
	}
	return false
}

func (m Matrix) CountSquareUnderDirectAttackFromWhite(s square.Index) int {
	count := 0
	attackedDirectly := peaceattacks.WhiteDirectsList[s]
	for _, direct := range attackedDirectly {
		if m[direct.Index] == direct.Peace {
			count++
		}
	}
	return count
}

func (m Matrix) SquareUnderDirectAttackExactlyFromWhite(s square.Index) (bool, square.Index) {
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

func (m Matrix) IsSquareUnderDirectAttackFromWhite(s square.Index) bool {
	attackedDirectly := peaceattacks.WhiteDirectsList[s]
	for _, direct := range attackedDirectly {
		if m[direct.Index] == direct.Peace {
			return true
		}
	}
	return false
}

func (m Matrix) IsSquareUnderAttackFromWhiteFromLeft(s square.Index) square.Index {
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

func (m Matrix) IsSquareUnderAttackFromWhiteFromRight(s square.Index) square.Index {
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

func (m Matrix) IsSquareUnderAttackFromWhiteFromUnder(s square.Index) square.Index {
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

func (m Matrix) IsSquareUnderAttackFromWhiteFromAbove(s square.Index) square.Index {
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

func (m Matrix) IsSquareUnderAttackFromWhiteFromLeftUnder(s square.Index) square.Index {
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

func (m Matrix) IsSquareUnderAttackFromWhiteFromLeftAbove(s square.Index) square.Index {
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

func (m Matrix) IsSquareUnderAttackFromWhiteFromRightUnder(s square.Index) square.Index {
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

func (m Matrix) IsSquareUnderAttackFromWhiteFromRightAbove(s square.Index) square.Index {
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

func (m Matrix) IsSquareUnderAttackFromWhite(s square.Index) bool {
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

func (m Matrix) CountSquareUnderAttackFromWhite(s square.Index) int {
	count := 0

	count += m.CountSquareUnderDirectAttackFromWhite(s)

	if m.IsSquareUnderAttackFromWhiteFromLeft(s) != square.InvalidIndex {
		count++
	}

	if m.IsSquareUnderAttackFromWhiteFromRight(s) != square.InvalidIndex {
		count++
	}

	if m.IsSquareUnderAttackFromWhiteFromUnder(s) != square.InvalidIndex {
		count++
	}

	if m.IsSquareUnderAttackFromWhiteFromAbove(s) != square.InvalidIndex {
		count++
	}

	if m.IsSquareUnderAttackFromWhiteFromLeftUnder(s) != square.InvalidIndex {
		count++
	}

	if m.IsSquareUnderAttackFromWhiteFromLeftAbove(s) != square.InvalidIndex {
		count++
	}

	if m.IsSquareUnderAttackFromWhiteFromRightUnder(s) != square.InvalidIndex {
		count++
	}

	if m.IsSquareUnderAttackFromWhiteFromRightAbove(s) != square.InvalidIndex {
		count++
	}

	return count
}
