package board

import (
	"github.com/ggeorgiev/instant-chess/src/peace"
	"github.com/ggeorgiev/instant-chess/src/peaceattacks"
	"github.com/ggeorgiev/instant-chess/src/square"
)

func (m Matrix) SquareUnderAttackFromWhiteKing(s square.Index) int {
	count := 0
	attackedFromKing := peaceattacks.FromKing[s]
	for _, kingSquare := range attackedFromKing {
		if m[kingSquare] == peace.WhiteKing {
			count++
		}
	}
	return count
}

func (m Matrix) SquareUnderAttackFromWhiteKnight(s square.Index) int {
	count := 0
	attackedFromKnight := peaceattacks.FromKnight[s]
	for _, knightSquare := range attackedFromKnight {
		if m[knightSquare] == peace.WhiteKnight {
			count++
		}
	}
	return count
}

func (m Matrix) SquareUnderAttackFromWhitePawn(s square.Index) int {
	count := 0
	attackedFromPawn := peaceattacks.FromWhitePawn[s]
	for _, pawnSquare := range attackedFromPawn {
		if m[pawnSquare] == peace.WhitePawn {
			count++
		}
	}
	return count
}

func (m Matrix) SquareUnderAttackFromWhiteFromLeft(s square.Index) bool {
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

func (m Matrix) SquareUnderAttackFromWhiteFromRight(s square.Index) bool {
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

func (m Matrix) SquareUnderAttackFromWhiteFromUnder(s square.Index) bool {
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

func (m Matrix) SquareUnderAttackFromWhiteFromAbove(s square.Index) bool {
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

func (m Matrix) SquareUnderAttackFromWhiteFromLeftUnder(s square.Index) bool {
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

func (m Matrix) SquareUnderAttackFromWhiteFromLeftAbove(s square.Index) bool {
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

func (m Matrix) SquareUnderAttackFromWhiteFromRightUnder(s square.Index) bool {
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

func (m Matrix) SquareUnderAttackFromWhiteFromRightAbove(s square.Index) bool {
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

func (m Matrix) SquareUnderAttackFromWhite(s square.Index, countTo int) int {
	count := 0

	count += m.SquareUnderAttackFromWhiteKing(s)
	if count >= countTo {
		return countTo
	}

	count += m.SquareUnderAttackFromWhiteKnight(s)
	if count >= countTo {
		return countTo
	}

	count += m.SquareUnderAttackFromWhitePawn(s)
	if count >= countTo {
		return countTo
	}

	if m.SquareUnderAttackFromWhiteFromLeft(s) {
		count++
		if count >= countTo {
			return count
		}
	}

	if m.SquareUnderAttackFromWhiteFromRight(s) {
		count++
		if count >= countTo {
			return count
		}
	}

	if m.SquareUnderAttackFromWhiteFromUnder(s) {
		count++
		if count >= countTo {
			return count
		}
	}

	if m.SquareUnderAttackFromWhiteFromAbove(s) {
		count++
		if count >= countTo {
			return count
		}
	}

	if m.SquareUnderAttackFromWhiteFromLeftUnder(s) {
		count++
		if count >= countTo {
			return count
		}
	}

	if m.SquareUnderAttackFromWhiteFromLeftAbove(s) {
		count++
		if count >= countTo {
			return count
		}
	}

	if m.SquareUnderAttackFromWhiteFromRightUnder(s) {
		count++
		if count >= countTo {
			return count
		}
	}

	if m.SquareUnderAttackFromWhiteFromRightAbove(s) {
		count++
		if count >= countTo {
			return count
		}
	}

	return count
}
