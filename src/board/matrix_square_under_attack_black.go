package board

import (
	"github.com/ggeorgiev/instant-chess/src/peace"
	"github.com/ggeorgiev/instant-chess/src/peaceattacks"
	"github.com/ggeorgiev/instant-chess/src/square"
)

func (m Matrix) SquareUnderAttackFromBlackKing(s square.Index) int {
	count := 0
	attackedFromKing := peaceattacks.FromKing[s]
	for _, kingSquare := range attackedFromKing {
		if m[kingSquare] == peace.BlackKing {
			count++
		}
	}
	return count
}

func (m Matrix) SquareUnderAttackFromBlackKnight(s square.Index) int {
	count := 0
	attackedFromKnight := peaceattacks.FromKnight[s]
	for _, knightSquare := range attackedFromKnight {
		if m[knightSquare] == peace.BlackKnight {
			count++
		}
	}
	return count
}

func (m Matrix) SquareUnderAttackFromBlackPawn(s square.Index) int {
	count := 0
	attackedFromPawn := peaceattacks.FromBlackPawn[s]
	for _, pawnSquare := range attackedFromPawn {
		if m[pawnSquare] == peace.BlackPawn {
			count++
		}
	}
	return count
}

func (m Matrix) SquareUnderAttackFromBlackFromLeft(s square.Index) bool {
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

func (m Matrix) SquareUnderAttackFromBlackFromRight(s square.Index) bool {
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

func (m Matrix) SquareUnderAttackFromBlackFromUnder(s square.Index) bool {
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

func (m Matrix) SquareUnderAttackFromBlackFromAbove(s square.Index) bool {
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

func (m Matrix) SquareUnderAttackFromBlackFromLeftUnder(s square.Index) bool {
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

func (m Matrix) SquareUnderAttackFromBlackFromLeftAbove(s square.Index) bool {
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

func (m Matrix) SquareUnderAttackFromBlackFromRightUnder(s square.Index) bool {
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

func (m Matrix) SquareUnderAttackFromBlackFromRightAbove(s square.Index) bool {
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

func (m Matrix) SquareUnderAttackFromBlack(s square.Index, countTo int) int {
	count := 0

	count += m.SquareUnderAttackFromBlackKing(s)
	if count >= countTo {
		return countTo
	}

	count += m.SquareUnderAttackFromBlackKnight(s)
	if count >= countTo {
		return countTo
	}

	count += m.SquareUnderAttackFromBlackPawn(s)
	if count >= countTo {
		return countTo
	}

	if m.SquareUnderAttackFromBlackFromLeft(s) {
		count++
		if count >= countTo {
			return count
		}
	}

	if m.SquareUnderAttackFromBlackFromRight(s) {
		count++
		if count >= countTo {
			return count
		}
	}

	if m.SquareUnderAttackFromBlackFromUnder(s) {
		count++
		if count >= countTo {
			return count
		}
	}

	if m.SquareUnderAttackFromBlackFromAbove(s) {
		count++
		if count >= countTo {
			return count
		}
	}

	if m.SquareUnderAttackFromBlackFromLeftUnder(s) {
		count++
		if count >= countTo {
			return count
		}
	}

	if m.SquareUnderAttackFromBlackFromLeftAbove(s) {
		count++
		if count >= countTo {
			return count
		}
	}

	if m.SquareUnderAttackFromBlackFromRightUnder(s) {
		count++
		if count >= countTo {
			return count
		}
	}

	if m.SquareUnderAttackFromBlackFromRightAbove(s) {
		count++
		if count >= countTo {
			return count
		}
	}

	return count
}
