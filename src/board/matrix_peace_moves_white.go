package board

import (
	"github.com/ggeorgiev/instant-chess/src/peace"
	"github.com/ggeorgiev/instant-chess/src/peacealignment"
	"github.com/ggeorgiev/instant-chess/src/peacemoves"
	"github.com/ggeorgiev/instant-chess/src/square"
)

func (m Matrix) WhiteKingTos(s square.Index) square.Indexes {
	var tos square.Indexes

	m[s] = peace.NoFigure

	kingMoves := peacemoves.KingSquareIndexes[s]
	for _, square := range kingMoves {
		figure := m[square]
		if figure.IsNoFigureOrBlack() && !m.IsWhiteChecked(square) {
			tos = append(tos, square)
		}
	}

	m[s] = peace.WhiteKing
	return tos
}

func (m Matrix) WhiteKnightNoCheckedTos(s square.Index) square.Indexes {
	var tos square.Indexes

	knightMoves := peacemoves.KnightSquareIndexes[s]
	for _, square := range knightMoves {
		if m[square].IsWhite() {
			continue
		}
		tos = append(tos, square)
	}

	return tos
}

func (m Matrix) WhiteRookNoCheckedTos(s square.Index) square.Indexes {
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

func (m Matrix) WhiteTos(s square.Index, kingSquare square.Index) square.Indexes {
	figure := m[s]
	if !figure.IsWhite() {
		return nil
	}

	if figure == peace.WhiteKing {
		return m.WhiteKingTos(s)
	}

	var tos square.Indexes
	switch figure {
	case peace.WhiteRook:
		tos = m.WhiteRookNoCheckedTos(s)
	case peace.WhiteKnight:
		tos = m.WhiteKnightNoCheckedTos(s)
	}

	original := m[s]
	m[s] = peace.NoFigure
	maybeChecked, relation := m.IsWhiteMaybeCheckedAfterMove(kingSquare, s)
	m[s] = original

	if !maybeChecked {
		return tos
	}

	kingRelation := peacealignment.SquareRelations[kingSquare]
	for _, to := range tos {
		if relation == kingRelation[to] {
			return square.Indexes{to}
		}
	}
	return nil
}
