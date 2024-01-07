package board

import (
	"github.com/ggeorgiev/instant-chess/src/peace"
	"github.com/ggeorgiev/instant-chess/src/peacemoves"
	"github.com/ggeorgiev/instant-chess/src/square"
)

func (m Matrix) BlackKingTos(s square.Index) square.Indexes {
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

func (m Matrix) BlackKnightNoCheckedTos(s square.Index, kingSquare square.Index) square.Indexes {
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

func (m Matrix) BlackRookNoCheckedTos(s square.Index, kingSquare square.Index) square.Indexes {
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

	return tos
}

func (m Matrix) BlackTos(s square.Index, kingSquare square.Index) square.Indexes {
	figure := m[s]
	if figure == peace.BlackKing {
		return m.BlackKingTos(s)
	}

	if figure.IsBlack() {
		original := m[s]
		m[s] = peace.NoFigure
		if m.IsBlackCheckedAfterMove(kingSquare, s) {
			return nil
		}
		m[s] = original
	}

	switch figure {
	case peace.BlackRook:
		return m.BlackRookNoCheckedTos(s, kingSquare)
	case peace.BlackKnight:
		return m.BlackKnightNoCheckedTos(s, kingSquare)
	}
	return nil
}
