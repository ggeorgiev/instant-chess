package peace

import (
	"github.com/ggeorgiev/instant-chess/src/move"
	"github.com/ggeorgiev/instant-chess/src/peaceplaces"
	"github.com/ggeorgiev/instant-chess/src/square"
)

func rightsFigureInternalHelper() []FigureMap {
	var peacesList []FigureMap
	for r := move.NoRights; r < move.Rights(0xFF); r++ {
		figures := make(FigureMap)
		if r.IsWhiteAnyCastling() {
			figures[peaceplaces.WhiteKingStartingPlace] = WhiteKing
		}
		if r.IsWhiteKingsideCastling() {
			figures[peaceplaces.WhiteRookKingsideStartingPlace] = WhiteRook
		}
		if r.IsWhiteQueensideCastling() {
			figures[peaceplaces.WhiteRookQueensideStartingPlace] = WhiteRook
		}
		if r.IsBlackAnyCastling() {
			figures[peaceplaces.BlackKingStartingPlace] = BlackKing
		}
		if r.IsBlackKingsideCastling() {
			figures[peaceplaces.BlackRookKingsideStartingPlace] = BlackRook
		}
		if r.IsBlackQueensideCastling() {
			figures[peaceplaces.BlackRookQueensideStartingPlace] = BlackRook
		}
		if r.EnPassantFile() != square.InvalidFile {
			figures[square.NewIndex(r.EnPassantFile(), peaceplaces.BlackPawnsJumpRank)] = WhitePawn
		}

		peacesList = append(peacesList, figures)
	}

	return peacesList
}
