package peace

import (
	"github.com/ggeorgiev/instant-chess/src/move"
	"github.com/ggeorgiev/instant-chess/src/place"
	"github.com/ggeorgiev/instant-chess/src/square"
)

func rightsFigureInternalHelper() []RightsMap {
	var peacesList []RightsMap
	for r := move.NoRights; r < move.Rights(0xFF); r++ {
		codes := make(RightsMap)
		if r.IsWhiteAnyCastling() {
			codes[place.WhiteKingStarting] = WhiteKing
		}
		if r.IsWhiteKingsideCastling() {
			codes[place.WhiteRookKingsideStarting] = WhiteRook
		}
		if r.IsWhiteQueensideCastling() {
			codes[place.WhiteRookQueensideStarting] = WhiteRook
		}
		if r.IsBlackAnyCastling() {
			codes[place.BlackKingStarting] = BlackKing
		}
		if r.IsBlackKingsideCastling() {
			codes[place.BlackRookKingsideStarting] = BlackRook
		}
		if r.IsBlackQueensideCastling() {
			codes[place.BlackRookQueensideStarting] = BlackRook
		}
		if r.EnPassantFile() != square.InvalidFile {
			codes[square.NewIndex(r.EnPassantFile(), place.BlackPawnsJumpRank)] = WhitePawn
		}

		peacesList = append(peacesList, codes)
	}

	return peacesList
}
