package peaceplaces

import "github.com/ggeorgiev/instant-chess/src/square"

const (
	WhiteKingStartingPlace          = square.Index(4)
	WhiteRookQueensideStartingPlace = square.Index(0)
	WhiteRookKingsideStartingPlace  = square.Index(7)

	BlackKingStartingPlace          = square.Index(60)
	BlackRookQueensideStartingPlace = square.Index(56)
	BlackRookKingsideStartingPlace  = square.Index(63)
)

const (
	WhiteKingKingsideCastlingPlace  = square.Index(6)
	WhiteKingQueensideCastlingPlace = square.Index(2)
	WhiteRookKingsideCastlingPlace  = square.Index(5)
	WhiteRookQueensideCastlingPlace = square.Index(3)

	BlackKingKingsideCastlingPlace  = square.Index(62)
	BlackKingQueensideCastlingPlace = square.Index(58)
	BlackRookKingsideCastlingPlace  = square.Index(61)
	BlackRookQueensideCastlingPlace = square.Index(59)
)
