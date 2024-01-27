package peace

import (
	"fmt"
)

type Code uint8

func Combine(color Color, tp Kind) Code {
	return Code(uint8(color) | uint8(tp))
}

const (
	Null        Code = 0
	WhitePawn   Code = Code(uint8(White) | uint8(Pawn))
	WhiteBishop Code = Code(uint8(White) | uint8(Bishop))
	WhiteKnight Code = Code(uint8(White) | uint8(Knight))
	WhiteRook   Code = Code(uint8(White) | uint8(Rook))
	WhiteQueen  Code = Code(uint8(White) | uint8(Queen))
	WhiteKing   Code = Code(uint8(White) | uint8(King))
	BlackPawn   Code = Code(uint8(Black) | uint8(Pawn))
	BlackBishop Code = Code(uint8(Black) | uint8(Bishop))
	BlackKnight Code = Code(uint8(Black) | uint8(Knight))
	BlackRook   Code = Code(uint8(Black) | uint8(Rook))
	BlackQueen  Code = Code(uint8(Black) | uint8(Queen))
	BlackKing   Code = Code(uint8(Black) | uint8(King))
)

func (f Code) Color() Color {
	return Color(uint8(f) & ColorMask)
}

func (f Code) IsWhite() bool {
	return f&White == White
}

func (f Code) IsBlack() bool {
	return f&Black == Black
}

func (f Code) IsLinearMover() bool {
	return uint8(f)&LinearMoverMask != 0
}

func (f Code) IsLinearMoverFrom(color Color) bool {
	return uint8(f)&(LinearMoverMask|ColorMask) == LinearMoverMask|uint8(color)
}

func (f Code) IsWhiteLinearMover() bool {
	return uint8(f)&(LinearMoverMask|ColorMask) == LinearMoverMask|White
}

func (f Code) IsBlackLinearMover() bool {
	return uint8(f)&(LinearMoverMask|ColorMask) == LinearMoverMask|Black
}

func (f Code) IsDiagonalMover() bool {
	return uint8(f)&DiagonalMoverMask != 0
}

func (f Code) IsDiagonalMoverFrom(color Color) bool {
	return uint8(f)&(DiagonalMoverMask|ColorMask) == DiagonalMoverMask|uint8(color)
}

func (f Code) IsWhiteDiagonalMover() bool {
	return uint8(f)&(DiagonalMoverMask|ColorMask) == DiagonalMoverMask|White
}

func (f Code) IsBlackDiagonalMover() bool {
	return uint8(f)&(DiagonalMoverMask|ColorMask) == DiagonalMoverMask|Black
}

func (f Code) IsBishop() bool {
	return f == WhiteBishop || f == BlackBishop
}

func (f Code) IsKnight() bool {
	return f == WhiteKnight || f == BlackKnight
}

func (f Code) IsRook() bool {
	return f == WhiteRook || f == BlackRook
}

func (f Code) IsQueen() bool {
	return f == WhiteQueen || f == BlackQueen
}

func (f Code) IsKing() bool {
	return f == WhiteKing || f == BlackKing
}

func (f Code) IsNull() bool {
	return f == Null
}

func (f Code) IsNullOr(color Color) bool {
	return uint8(f)&ColorMask != uint8(color.Oponent())
}

func (f Code) IsNullOrNot(color Color) bool {
	return uint8(f)&ColorMask != uint8(color)
}

func (f Code) IsNullOrWhite() bool {
	return uint8(f)&Black == 0
}

func (f Code) IsNullOrBlack() bool {
	return uint8(f)&White == 0
}

func (f Code) Symbol() string {
	switch f {
	case Null:
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

func FromSymbol(symbol rune) (Code, error) {
	switch symbol {
	case 32:
		return Null, nil
	case 9820:
		return BlackRook, nil
	case 9822:
		return BlackKnight, nil
	case 9821:
		return BlackBishop, nil
	case 9819:
		return BlackQueen, nil
	case 9818:
		return BlackKing, nil
	case 9823:
		return BlackPawn, nil
	case 9817:
		return WhitePawn, nil
	case 9814:
		return WhiteRook, nil
	case 9816:
		return WhiteKnight, nil
	case 9815:
		return WhiteBishop, nil
	case 9813:
		return WhiteQueen, nil
	case 9812:
		return WhiteKing, nil
	}

	return Null, fmt.Errorf("unacceptable symbol: %d", symbol)
}
