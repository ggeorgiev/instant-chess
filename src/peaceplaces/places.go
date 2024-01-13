package peaceplaces

import "github.com/ggeorgiev/instant-chess/src/square"

const (
	WhiteKingStartingPlace          = square.E1
	WhiteRookQueensideStartingPlace = square.A1
	WhiteRookKingsideStartingPlace  = square.H1

	BlackKingStartingPlace          = square.E8
	BlackRookQueensideStartingPlace = square.A8
	BlackRookKingsideStartingPlace  = square.H8
)

const (
	WhiteKingKingsideCastlingPlace  = square.G1
	WhiteKingQueensideCastlingPlace = square.C1
	WhiteRookKingsideCastlingPlace  = square.F1
	WhiteRookQueensideCastlingPlace = square.D1

	BlackKingKingsideCastlingPlace  = square.G8
	BlackKingQueensideCastlingPlace = square.C8
	BlackRookKingsideCastlingPlace  = square.F8
	BlackRookQueensideCastlingPlace = square.D8
)
