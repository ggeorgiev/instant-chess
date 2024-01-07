package board

import (
	"github.com/ggeorgiev/instant-chess/src/peace"
	"github.com/ggeorgiev/instant-chess/src/peaceattacks"
	"github.com/ggeorgiev/instant-chess/src/square"
)

func (m Matrix) CountSquareUnderAttackFromBlackKing(s square.Index) int {
	count := 0
	attackedFromKing := peaceattacks.FromKing[s]
	for _, kingSquare := range attackedFromKing {
		if m[kingSquare] == peace.BlackKing {
			count++
		}
	}
	return count
}

func (m Matrix) IsSquareUnderAttackFromBlackKing(s square.Index) bool {
	attackedFromKing := peaceattacks.FromKing[s]
	for _, kingSquare := range attackedFromKing {
		if m[kingSquare] == peace.BlackKing {
			return true
		}
	}
	return false
}

func (m Matrix) CountSquareUnderAttackFromBlackKnight(s square.Index) int {
	count := 0
	attackedFromKnight := peaceattacks.FromKnight[s]
	for _, knightSquare := range attackedFromKnight {
		if m[knightSquare] == peace.BlackKnight {
			count++
		}
	}
	return count
}

func (m Matrix) IsSquareUnderAttackFromBlackKnight(s square.Index) bool {
	attackedFromKnight := peaceattacks.FromKnight[s]
	for _, knightSquare := range attackedFromKnight {
		if m[knightSquare] == peace.BlackKnight {
			return true
		}
	}
	return false
}

func (m Matrix) CountSquareUnderAttackFromBlackPawn(s square.Index) int {
	count := 0
	attackedFromPawn := peaceattacks.FromBlackPawn[s]
	for _, pawnSquare := range attackedFromPawn {
		if m[pawnSquare] == peace.BlackPawn {
			count++
		}
	}
	return count
}

func (m Matrix) IsSquareUnderAttackFromBlackPawn(s square.Index) bool {
	attackedFromPawn := peaceattacks.FromBlackPawn[s]
	for _, pawnSquare := range attackedFromPawn {
		if m[pawnSquare] == peace.BlackPawn {
			return true
		}
	}
	return false
}

func (m Matrix) CountSquareUnderDirectAttackFromBlack(s square.Index) int {
	count := 0
	attackedDirectly := peaceattacks.BlackDirectsList[s]
	for _, direct := range attackedDirectly {
		if m[direct.Index] == direct.Peace {
			count++
		}
	}
	return count
}

func (m Matrix) IsSquareUnderDirectAttackFromBlack(s square.Index) bool {
	attackedDirectly := peaceattacks.BlackDirectsList[s]
	for _, direct := range attackedDirectly {
		if m[direct.Index] == direct.Peace {
			return true
		}
	}
	return false
}

func (m Matrix) IsSquareUnderAttackFromBlackFromLeft(s square.Index) bool {
	rank := s.Rank()
	for f := s.File() - 1; f >= square.ZeroFile; f-- {
		figure := m[square.NewIndex(f, rank)]
		if figure.IsBlackLinearMover() {
			return true
		}
		if figure != peace.NoFigure {
			return false
		}
	}
	return false
}

func (m Matrix) IsSquareUnderAttackFromBlackFromRight(s square.Index) bool {
	rank := s.Rank()
	for f := s.File() + 1; f <= square.LastFile; f++ {
		figure := m[square.NewIndex(f, rank)]
		if figure.IsBlackLinearMover() {
			return true
		}
		if figure != peace.NoFigure {
			return false
		}
	}
	return false
}

func (m Matrix) IsSquareUnderAttackFromBlackFromUnder(s square.Index) bool {
	file := s.File()
	for r := s.Rank() - 1; r >= square.ZeroRank; r-- {
		figure := m[square.NewIndex(file, r)]
		if figure.IsBlackLinearMover() {
			return true
		}
		if figure != peace.NoFigure {
			return false
		}
	}
	return false
}

func (m Matrix) IsSquareUnderAttackFromBlackFromAbove(s square.Index) bool {
	file := s.File()
	for r := s.Rank() + 1; r <= square.LastRank; r++ {
		figure := m[square.NewIndex(file, r)]
		if figure.IsBlackLinearMover() {
			return true
		}
		if figure != peace.NoFigure {
			return false
		}
	}
	return false
}

func (m Matrix) IsSquareUnderAttackFromBlackFromLeftUnder(s square.Index) bool {
	f := s.File() - 1
	r := s.Rank() - 1
	for f >= square.ZeroFile && r >= square.ZeroRank {
		figure := m[square.NewIndex(f, r)]
		if figure.IsBlackDiagonalMover() {
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

func (m Matrix) IsSquareUnderAttackFromBlackFromLeftAbove(s square.Index) bool {
	f := s.File() - 1
	r := s.Rank() + 1
	for f >= square.ZeroFile && r <= square.LastRank {
		figure := m[square.NewIndex(f, r)]
		if figure.IsBlackDiagonalMover() {
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

func (m Matrix) IsSquareUnderAttackFromBlackFromRightUnder(s square.Index) bool {
	f := s.File() + 1
	r := s.Rank() - 1
	for f <= square.LastFile && r >= square.ZeroRank {
		figure := m[square.NewIndex(f, r)]
		if figure.IsBlackDiagonalMover() {
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

func (m Matrix) IsSquareUnderAttackFromBlackFromRightAbove(s square.Index) bool {
	f := s.File() + 1
	r := s.Rank() + 1
	for f <= square.LastFile && r <= square.LastRank {
		figure := m[square.NewIndex(f, r)]
		if figure.IsBlackDiagonalMover() {
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

func (m Matrix) IsSquareUnderAttackFromBlack(s square.Index) bool {
	if m.IsSquareUnderAttackFromBlackKing(s) {
		return true
	}
	if m.IsSquareUnderAttackFromBlackKnight(s) {
		return true
	}
	if m.IsSquareUnderAttackFromBlackPawn(s) {
		return true
	}

	if m.IsSquareUnderAttackFromBlackFromLeft(s) {
		return true
	}

	if m.IsSquareUnderAttackFromBlackFromRight(s) {
		return true
	}

	if m.IsSquareUnderAttackFromBlackFromUnder(s) {
		return true
	}

	if m.IsSquareUnderAttackFromBlackFromAbove(s) {
		return true
	}

	if m.IsSquareUnderAttackFromBlackFromLeftUnder(s) {
		return true
	}

	if m.IsSquareUnderAttackFromBlackFromLeftAbove(s) {
		return true
	}

	if m.IsSquareUnderAttackFromBlackFromRightUnder(s) {
		return true
	}

	if m.IsSquareUnderAttackFromBlackFromRightAbove(s) {
		return true
	}

	return false
}

func (m Matrix) CountSquareUnderAttackFromBlack(s square.Index) int {
	count := 0

	count += m.CountSquareUnderAttackFromBlackKing(s)
	count += m.CountSquareUnderAttackFromBlackKnight(s)
	count += m.CountSquareUnderAttackFromBlackPawn(s)

	if m.IsSquareUnderAttackFromBlackFromLeft(s) {
		count++
	}

	if m.IsSquareUnderAttackFromBlackFromRight(s) {
		count++
	}

	if m.IsSquareUnderAttackFromBlackFromUnder(s) {
		count++
	}

	if m.IsSquareUnderAttackFromBlackFromAbove(s) {
		count++
	}

	if m.IsSquareUnderAttackFromBlackFromLeftUnder(s) {
		count++
	}

	if m.IsSquareUnderAttackFromBlackFromLeftAbove(s) {
		count++
	}

	if m.IsSquareUnderAttackFromBlackFromRightUnder(s) {
		count++
	}

	if m.IsSquareUnderAttackFromBlackFromRightAbove(s) {
		count++
	}

	return count
}
