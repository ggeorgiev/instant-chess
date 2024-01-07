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

func (m Matrix) IsSquareUnderDirectAttackFromWhite(s square.Index) bool {
	attackedDirectly := peaceattacks.WhiteDirectsList[s]
	for _, direct := range attackedDirectly {
		if m[direct.Index] == direct.Peace {
			return true
		}
	}
	return false
}

func (m Matrix) IsSquareUnderAttackFromWhiteFromLeft(s square.Index) bool {
	rank := s.Rank()
	for f := s.File() - 1; f >= square.ZeroFile; f-- {
		figure := m[square.NewIndex(f, rank)]
		if figure.IsWhiteLinearMover() {
			return true
		}
		if figure != peace.NoFigure {
			return false
		}
	}
	return false
}

func (m Matrix) IsSquareUnderAttackFromWhiteFromRight(s square.Index) bool {
	rank := s.Rank()
	for f := s.File() + 1; f <= square.LastFile; f++ {
		figure := m[square.NewIndex(f, rank)]
		if figure.IsWhiteLinearMover() {
			return true
		}
		if figure != peace.NoFigure {
			return false
		}
	}
	return false
}

func (m Matrix) IsSquareUnderAttackFromWhiteFromUnder(s square.Index) bool {
	file := s.File()
	for r := s.Rank() - 1; r >= square.ZeroRank; r-- {
		figure := m[square.NewIndex(file, r)]
		if figure.IsWhiteLinearMover() {
			return true
		}
		if figure != peace.NoFigure {
			return false
		}
	}
	return false
}

func (m Matrix) IsSquareUnderAttackFromWhiteFromAbove(s square.Index) bool {
	file := s.File()
	for r := s.Rank() + 1; r <= square.LastRank; r++ {
		figure := m[square.NewIndex(file, r)]
		if figure.IsWhiteLinearMover() {
			return true
		}
		if figure != peace.NoFigure {
			return false
		}
	}
	return false
}

func (m Matrix) IsSquareUnderAttackFromWhiteFromLeftUnder(s square.Index) bool {
	f := s.File() - 1
	r := s.Rank() - 1
	for f >= square.ZeroFile && r >= square.ZeroRank {
		figure := m[square.NewIndex(f, r)]
		if figure.IsWhiteDiagonalMover() {
			return true
		}
		if figure != peace.NoFigure {
			return false
		}
		f--
		r--
	}
	return false
}

func (m Matrix) IsSquareUnderAttackFromWhiteFromLeftAbove(s square.Index) bool {
	f := s.File() - 1
	r := s.Rank() + 1
	for f >= square.ZeroFile && r <= square.LastRank {
		figure := m[square.NewIndex(f, r)]
		if figure.IsWhiteDiagonalMover() {
			return true
		}
		if figure != peace.NoFigure {
			return false
		}
		f--
		r++
	}
	return false
}

func (m Matrix) IsSquareUnderAttackFromWhiteFromRightUnder(s square.Index) bool {
	f := s.File() + 1
	r := s.Rank() - 1
	for f <= square.LastFile && r >= square.ZeroRank {
		figure := m[square.NewIndex(f, r)]
		if figure.IsWhiteDiagonalMover() {
			return true
		}
		if figure != peace.NoFigure {
			return false
		}
		f++
		r--
	}
	return false
}

func (m Matrix) IsSquareUnderAttackFromWhiteFromRightAbove(s square.Index) bool {
	f := s.File() + 1
	r := s.Rank() + 1
	for f <= square.LastFile && r <= square.LastRank {
		figure := m[square.NewIndex(f, r)]
		if figure.IsWhiteDiagonalMover() {
			return true
		}
		if figure != peace.NoFigure {
			return false
		}
		f++
		r++
	}
	return false
}

func (m Matrix) IsSquareUnderAttackFromWhite(s square.Index) bool {
	if m.IsSquareUnderDirectAttackFromWhite(s) {
		return true
	}

	if m.IsSquareUnderAttackFromWhiteFromLeft(s) {
		return true
	}

	if m.IsSquareUnderAttackFromWhiteFromRight(s) {
		return true
	}

	if m.IsSquareUnderAttackFromWhiteFromUnder(s) {
		return true
	}

	if m.IsSquareUnderAttackFromWhiteFromAbove(s) {
		return true
	}

	if m.IsSquareUnderAttackFromWhiteFromLeftUnder(s) {
		return true
	}

	if m.IsSquareUnderAttackFromWhiteFromLeftAbove(s) {
		return true
	}

	if m.IsSquareUnderAttackFromWhiteFromRightUnder(s) {
		return true
	}

	if m.IsSquareUnderAttackFromWhiteFromRightAbove(s) {
		return true
	}

	return false
}

func (m Matrix) CountSquareUnderAttackFromWhite(s square.Index) int {
	count := 0

	count += m.CountSquareUnderDirectAttackFromWhite(s)

	if m.IsSquareUnderAttackFromWhiteFromLeft(s) {
		count++
	}

	if m.IsSquareUnderAttackFromWhiteFromRight(s) {
		count++
	}

	if m.IsSquareUnderAttackFromWhiteFromUnder(s) {
		count++
	}

	if m.IsSquareUnderAttackFromWhiteFromAbove(s) {
		count++
	}

	if m.IsSquareUnderAttackFromWhiteFromLeftUnder(s) {
		count++
	}

	if m.IsSquareUnderAttackFromWhiteFromLeftAbove(s) {
		count++
	}

	if m.IsSquareUnderAttackFromWhiteFromRightUnder(s) {
		count++
	}

	if m.IsSquareUnderAttackFromWhiteFromRightAbove(s) {
		count++
	}

	return count
}
