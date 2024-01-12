package board

import (
	"github.com/ggeorgiev/instant-chess/src/peace"
	"github.com/ggeorgiev/instant-chess/src/peacealignment"
	"github.com/ggeorgiev/instant-chess/src/peacemoves"
	"github.com/ggeorgiev/instant-chess/src/square"
)

func (m *Matrix) BlackKingTos(s square.Index) square.Indexes {
	var tos square.Indexes

	m[s] = peace.NoFigure

	kingMoves := peacemoves.KingSquareIndexes[s]
	for _, square := range kingMoves {
		figure := m[square]
		if figure.IsNoFigureOrWhite() && !m.IsBlackChecked(square) {
			tos = append(tos, square)
		}
	}

	m[s] = peace.BlackKing
	return tos
}

func (m *Matrix) BlackKnightNoCheckedTos(s square.Index) square.Indexes {
	var tos square.Indexes

	knightMoves := peacemoves.KnightSquareIndexes[s]
	for _, square := range knightMoves {
		if m[square].IsBlack() {
			continue
		}
		tos = append(tos, square)
	}

	return tos
}

func (m *Matrix) BlackBishopNoCheckedTos(s square.Index, vector peacealignment.Vector) square.Indexes {
	var tos square.Indexes

	check := func(square square.Index) bool {
		figure := m[square]
		if figure.IsBlack() {
			return false
		}
		tos = append(tos, square)
		return figure.IsNoFigure()
	}

	file := s.File()
	rank := s.Rank()

	if vector != peacealignment.CounterDiagonal {
		f := file - 1
		r := rank - 1
		for f >= square.ZeroFile && r >= square.ZeroRank {
			if !check(square.NewIndex(f, r)) {
				break
			}
			f--
			r--
		}

		f = file + 1
		r = rank + 1
		for f <= square.LastFile && r <= square.LastRank {
			if !check(square.NewIndex(f, r)) {
				break
			}
			f++
			r++
		}
	}

	if vector != peacealignment.Diagonal {
		f := file + 1
		r := rank - 1
		for f <= square.LastFile && r >= square.ZeroRank {
			if !check(square.NewIndex(f, r)) {
				break
			}
			f++
			r--
		}

		f = file - 1
		r = rank + 1
		for f >= square.ZeroFile && r <= square.LastRank {
			if !check(square.NewIndex(f, r)) {
				break
			}
			f--
			r++
		}
	}

	return tos
}

func (m *Matrix) BlackBishopCapture(s square.Index, vector peacealignment.Vector, capture square.Index) square.Indexes {
	var tos square.Indexes

	check := func(sq square.Index) bool {
		figure := m[sq]
		if figure.IsBlack() {
			return false
		}
		if sq == capture {
			tos = square.Indexes{capture}
		}
		return figure.IsNoFigure()
	}

	file := s.File()
	rank := s.Rank()

	if vector != peacealignment.CounterDiagonal {
		f := file - 1
		r := rank - 1
		for f >= square.ZeroFile && r >= square.ZeroRank {
			if !check(square.NewIndex(f, r)) {
				break
			}
			f--
			r--
		}

		f = file + 1
		r = rank + 1
		for f <= square.LastFile && r <= square.LastRank {
			if !check(square.NewIndex(f, r)) {
				break
			}
			f++
			r++
		}
	}

	if vector != peacealignment.Diagonal {
		f := file + 1
		r := rank - 1
		for f <= square.LastFile && r >= square.ZeroRank {
			if !check(square.NewIndex(f, r)) {
				break
			}
			f++
			r--
		}

		f = file - 1
		r = rank + 1
		for f >= square.ZeroFile && r <= square.LastRank {
			if !check(square.NewIndex(f, r)) {
				break
			}
			f--
			r++
		}
	}

	return tos
}

func (m *Matrix) BlackRookNoCheckedTos(s square.Index, vector peacealignment.Vector) square.Indexes {
	// TODO: optimize
	var tos square.Indexes

	check := func(square square.Index) bool {
		figure := m[square]
		if figure.IsBlack() {
			return false
		}
		tos = append(tos, square)
		return figure.IsNoFigure()
	}

	file := s.File()
	rank := s.Rank()

	if vector != peacealignment.File {
		for f := file - 1; f >= square.ZeroFile; f-- {
			if !check(square.NewIndex(f, rank)) {
				break
			}
		}
		for f := file + 1; f <= square.LastFile; f++ {
			if !check(square.NewIndex(f, rank)) {
				break
			}
		}
	}

	if vector != peacealignment.Rank {
		for r := rank - 1; r >= square.ZeroRank; r-- {
			if !check(square.NewIndex(file, r)) {
				break
			}
		}
		for r := rank + 1; r <= square.LastRank; r++ {
			if !check(square.NewIndex(file, r)) {
				break
			}
		}
	}

	return tos
}

func (m *Matrix) BlackRookCapture(s square.Index, vector peacealignment.Vector, capture square.Index) square.Indexes {
	// TODO: optimize
	var tos square.Indexes

	check := func(sq square.Index) bool {
		figure := m[sq]
		if figure.IsBlack() {
			return false
		}
		if sq == capture {
			tos = square.Indexes{capture}
		}
		return figure.IsNoFigure()
	}

	file := s.File()
	rank := s.Rank()

	if vector != peacealignment.File {
		for f := file - 1; f >= square.ZeroFile; f-- {
			if !check(square.NewIndex(f, rank)) {
				break
			}
		}
		for f := file + 1; f <= square.LastFile; f++ {
			if !check(square.NewIndex(f, rank)) {
				break
			}
		}
	}

	if vector != peacealignment.Rank {
		for r := rank - 1; r >= square.ZeroRank; r-- {
			if !check(square.NewIndex(file, r)) {
				break
			}
		}
		for r := rank + 1; r <= square.LastRank; r++ {
			if !check(square.NewIndex(file, r)) {
				break
			}
		}
	}

	return tos
}

func (m *Matrix) BlackQueenNoCheckedTos(s square.Index, vector peacealignment.Vector) square.Indexes {
	var tos square.Indexes

	check := func(square square.Index) bool {
		figure := m[square]
		if figure.IsBlack() {
			return false
		}
		tos = append(tos, square)
		return figure.IsNoFigure()
	}

	file := s.File()
	rank := s.Rank()

	if vector != peacealignment.File {
		for f := file - 1; f >= square.ZeroFile; f-- {
			if !check(square.NewIndex(f, rank)) {
				break
			}
		}
		for f := file + 1; f <= square.LastFile; f++ {
			if !check(square.NewIndex(f, rank)) {
				break
			}
		}
	}

	if vector != peacealignment.Rank {
		for r := rank - 1; r >= square.ZeroRank; r-- {
			if !check(square.NewIndex(file, r)) {
				break
			}
		}
		for r := rank + 1; r <= square.LastRank; r++ {
			if !check(square.NewIndex(file, r)) {
				break
			}
		}
	}

	if vector != peacealignment.CounterDiagonal {
		f := file - 1
		r := rank - 1
		for f >= square.ZeroFile && r >= square.ZeroRank {
			if !check(square.NewIndex(f, r)) {
				break
			}
			f--
			r--
		}

		f = file + 1
		r = rank + 1
		for f <= square.LastFile && r <= square.LastRank {
			if !check(square.NewIndex(f, r)) {
				break
			}
			f++
			r++
		}
	}

	if vector != peacealignment.Diagonal {
		f := file + 1
		r := rank - 1
		for f <= square.LastFile && r >= square.ZeroRank {
			if !check(square.NewIndex(f, r)) {
				break
			}
			f++
			r--
		}

		f = file - 1
		r = rank + 1
		for f >= square.ZeroFile && r <= square.LastRank {
			if !check(square.NewIndex(f, r)) {
				break
			}
			f--
			r++
		}
	}

	return tos
}
