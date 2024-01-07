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

	m[s] = peace.NoFigure

	// TODO: move this to the caller outside, it can be combined
	if !m.IsBlackCheckedAfterMove(kingSquare, s) {
		knightMoves := peacemoves.KnightSquareIndexes[s]
		for _, square := range knightMoves {
			if m[square].IsBlack() {
				continue
			}
			tos = append(tos, square)
		}
	}

	m[s] = peace.BlackKnight
	return tos
}
