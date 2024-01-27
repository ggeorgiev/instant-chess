package matrix

import (
	"github.com/ggeorgiev/instant-chess/src/alignment"
	"github.com/ggeorgiev/instant-chess/src/move"
	"github.com/ggeorgiev/instant-chess/src/peace"
	"github.com/ggeorgiev/instant-chess/src/square"
)

func (m *Matrix) WhiteKingTos(s square.Index) square.Indexes {
	var tos square.Indexes

	m[s] = peace.NoFigure

	kingMoves := move.KingSquareIndexes[s]
	for _, square := range kingMoves {
		figure := m[square]
		if figure.IsNoFigureOrBlack() && !m.IsWhiteChecked(square) {
			tos = append(tos, square)
		}
	}

	m[s] = peace.WhiteKing
	return tos
}

func (m *Matrix) WhiteKnightNoCheckedTos(s square.Index) square.Indexes {
	var tos square.Indexes

	knightMoves := move.KnightSquareIndexes[s]
	for _, square := range knightMoves {
		if m[square].IsWhite() {
			continue
		}
		tos = append(tos, square)
	}

	return tos
}

func (m *Matrix) WhiteBishopNoCheckedTos(s square.Index, vector alignment.Vector) square.Indexes {
	var tos square.Indexes

	check := func(square square.Index) bool {
		figure := m[square]
		if figure.IsWhite() {
			return false
		}
		tos = append(tos, square)
		return figure.IsNoFigure()
	}

	file := s.File()
	rank := s.Rank()

	if vector != alignment.CounterDiagonal {
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

	if vector != alignment.Diagonal {
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

func (m *Matrix) WhiteBishopCapture(s square.Index, vector alignment.Vector, capture square.Index) square.Indexes {
	var tos square.Indexes

	check := func(sq square.Index) bool {
		figure := m[sq]
		if figure.IsWhite() {
			return false
		}
		if sq == capture {
			tos = square.Indexes{capture}
		}
		return figure.IsNoFigure()
	}

	file := s.File()
	rank := s.Rank()

	if vector != alignment.CounterDiagonal {
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

	if vector != alignment.Diagonal {
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

func (m *Matrix) WhiteRookNoCheckedTos(s square.Index, vector alignment.Vector) square.Indexes {
	// TODO: optimize
	var tos square.Indexes

	check := func(square square.Index) bool {
		figure := m[square]
		if figure.IsWhite() {
			return false
		}
		tos = append(tos, square)
		return figure.IsNoFigure()
	}

	file := s.File()
	rank := s.Rank()

	if vector != alignment.File {
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

	if vector != alignment.Rank {
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

func (m *Matrix) WhiteRookCapture(s square.Index, vector alignment.Vector, capture square.Index) square.Indexes {
	// TODO: optimize
	var tos square.Indexes

	check := func(sq square.Index) bool {
		figure := m[sq]
		if figure.IsWhite() {
			return false
		}
		if sq == capture {
			tos = square.Indexes{capture}
		}
		return figure.IsNoFigure()
	}

	file := s.File()
	rank := s.Rank()

	if vector != alignment.File {
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

	if vector != alignment.Rank {
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

func (m *Matrix) WhiteQueenNoCheckedTos(s square.Index, vector alignment.Vector) square.Indexes {
	var tos square.Indexes

	check := func(square square.Index) bool {
		figure := m[square]
		if figure.IsWhite() {
			return false
		}
		tos = append(tos, square)
		return figure.IsNoFigure()
	}

	file := s.File()
	rank := s.Rank()

	if vector != alignment.File {
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

	if vector != alignment.Rank {
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

	if vector != alignment.CounterDiagonal {
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

	if vector != alignment.Diagonal {
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
