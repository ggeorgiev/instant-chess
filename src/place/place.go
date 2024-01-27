package place

import "github.com/ggeorgiev/instant-chess/src/square"

const (
	WhiteKingStarting          = square.E1
	WhiteRookQueensideStarting = square.A1
	WhiteRookKingsideStarting  = square.H1

	BlackKingStarting          = square.E8
	BlackRookQueensideStarting = square.A8
	BlackRookKingsideStarting  = square.H8
)

const (
	WhiteKingKingsideCastling  = square.G1
	WhiteKingQueensideCastling = square.C1
	WhiteRookKingsideCastling  = square.F1
	WhiteRookQueensideCastling = square.D1

	BlackKingKingsideCastling  = square.G8
	BlackKingQueensideCastling = square.C8
	BlackRookKingsideCastling  = square.F8
	BlackRookQueensideCastling = square.D8
)
