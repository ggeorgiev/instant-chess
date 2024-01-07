package board

import (
	"github.com/ggeorgiev/instant-chess/src/peace"
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
