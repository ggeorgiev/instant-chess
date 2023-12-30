package chess

import "fmt"

const (
	ColorMask = 0b11
	White     = 0b01
	Black     = 0b10
)

type PeaceColor uint8

const (
	PeaceColorWhite PeaceColor = White
	PeaceColorBlack PeaceColor = Black
)

func (pc PeaceColor) Oponent() PeaceColor {
	return pc ^ ColorMask
}

type PeaceType uint8

const (
	Pawn   PeaceType = 0b001000
	Bishop PeaceType = 0b010000
	Knight PeaceType = 0b011000
	Rook   PeaceType = 0b100000
	Queen  PeaceType = 0b101000
	King   PeaceType = 0b110000
)

type Peace uint8

func Combine(color PeaceColor, tp PeaceType) Peace {
	return Peace(uint8(color) | uint8(tp))
}

const (
	Empty       Peace = 0
	WhitePawn   Peace = Peace(uint8(White) | uint8(Pawn))
	WhiteBishop Peace = Peace(uint8(White) | uint8(Bishop))
	WhiteKnight Peace = Peace(uint8(White) | uint8(Knight))
	WhiteRook   Peace = Peace(uint8(White) | uint8(Rook))
	WhiteQueen  Peace = Peace(uint8(White) | uint8(Queen))
	WhiteKing   Peace = Peace(uint8(White) | uint8(King))
	BlackPawn   Peace = Peace(uint8(Black) | uint8(Pawn))
	BlackBishop Peace = Peace(uint8(Black) | uint8(Bishop))
	BlackKnight Peace = Peace(uint8(Black) | uint8(Knight))
	BlackRook   Peace = Peace(uint8(Black) | uint8(Rook))
	BlackQueen  Peace = Peace(uint8(Black) | uint8(Queen))
	BlackKing   Peace = Peace(uint8(Black) | uint8(King))
)

func (p Peace) Color() PeaceColor {
	return PeaceColor(uint8(p) & 0b111)
}

func (p Peace) IsBlack() bool {
	return p&Black == Black
}

func (p Peace) IsWhite() bool {
	return p&White == White
}

func (p Peace) IsKing() bool {
	return p == WhiteKing || p == BlackKing
}

func (p Peace) IsRook() bool {
	return p == WhiteRook || p == BlackRook
}

func (p Peace) IsEmpty() bool {
	return p == Empty
}

func (p Peace) IsEmptyOr(color PeaceColor) bool {
	return uint8(p)&ColorMask != uint8(color.Oponent())
}

func (p Peace) IsEmptyOrNot(color PeaceColor) bool {
	return uint8(p)&ColorMask != uint8(color)
}

func (p Peace) IsEmptyOrWhite() bool {
	return uint8(p)&Black == 0
}

func (p Peace) IsEmptyOrBlack() bool {
	return uint8(p)&White == 0
}

func (p Peace) Symbol() string {
	switch p {
	case Empty:
		return ` `
	case WhitePawn:
		return `♙`
	case WhiteBishop:
		return `♗`
	case WhiteKnight:
		return `♘`
	case WhiteRook:
		return `♖`
	case WhiteQueen:
		return `♕`
	case WhiteKing:
		return `♔`
	case BlackPawn:
		return `♟︎`
	case BlackBishop:
		return `♝`
	case BlackKnight:
		return `♞`
	case BlackRook:
		return `♜`
	case BlackQueen:
		return `♛`
	case BlackKing:
		return `♚`
	}
	panic("impossible")
}

func PeaceFromSymbol(symbol rune) Peace {
	switch symbol {
	case 32:
		return Empty
	case 9820:
		return BlackRook
	case 9822:
		return BlackKnight
	case 9821:
		return BlackBishop
	case 9819:
		return BlackQueen
	case 9818:
		return BlackKing
	case 9823:
		return BlackPawn
	case 9817:
		return WhitePawn
	case 9814:
		return WhiteRook
	case 9816:
		return WhiteKnight
	case 9815:
		return WhiteBishop
	case 9813:
		return WhiteQueen
	case 9812:
		return WhiteKing
	}
	panic(fmt.Sprintf("imposible: %d", symbol))
}
