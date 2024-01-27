package peace

import (
	"github.com/ggeorgiev/instant-chess/src/move"
	"github.com/ggeorgiev/instant-chess/src/place"
	"github.com/ggeorgiev/instant-chess/src/square"
)

func rightsFigureInternalHelper() []FigureMap {
	var peacesList []FigureMap
	for r := move.NoRights; r < move.Rights(0xFF); r++ {
		figures := make(FigureMap)
		if r.IsWhiteAnyCastling() {
			figures[place.WhiteKingStarting] = WhiteKing
		}
		if r.IsWhiteKingsideCastling() {
			figures[place.WhiteRookKingsideStarting] = WhiteRook
		}
		if r.IsWhiteQueensideCastling() {
			figures[place.WhiteRookQueensideStarting] = WhiteRook
		}
		if r.IsBlackAnyCastling() {
			figures[place.BlackKingStarting] = BlackKing
		}
		if r.IsBlackKingsideCastling() {
			figures[place.BlackRookKingsideStarting] = BlackRook
		}
		if r.IsBlackQueensideCastling() {
			figures[place.BlackRookQueensideStarting] = BlackRook
		}
		if r.EnPassantFile() != square.InvalidFile {
			figures[square.NewIndex(r.EnPassantFile(), place.BlackPawnsJumpRank)] = WhitePawn
		}

		peacesList = append(peacesList, figures)
	}

	return peacesList
}
