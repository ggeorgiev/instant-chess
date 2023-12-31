package peace

import "fmt"

type Figure uint8
type Figures []Figure

func Combine(color Color, tp Kind) Figure {
	return Figure(uint8(color) | uint8(tp))
}

const (
	NoFigure    Figure = 0
	WhitePawn   Figure = Figure(uint8(White) | uint8(Pawn))
	WhiteBishop Figure = Figure(uint8(White) | uint8(Bishop))
	WhiteKnight Figure = Figure(uint8(White) | uint8(Knight))
	WhiteRook   Figure = Figure(uint8(White) | uint8(Rook))
	WhiteQueen  Figure = Figure(uint8(White) | uint8(Queen))
	WhiteKing   Figure = Figure(uint8(White) | uint8(King))
	BlackPawn   Figure = Figure(uint8(Black) | uint8(Pawn))
	BlackBishop Figure = Figure(uint8(Black) | uint8(Bishop))
	BlackKnight Figure = Figure(uint8(Black) | uint8(Knight))
	BlackRook   Figure = Figure(uint8(Black) | uint8(Rook))
	BlackQueen  Figure = Figure(uint8(Black) | uint8(Queen))
	BlackKing   Figure = Figure(uint8(Black) | uint8(King))
)

func (f Figure) Color() Color {
	return Color(uint8(f) & ColorMask)
}

func (f Figure) IsWhite() bool {
	return f&White == White
}

func (f Figure) IsBlack() bool {
	return f&Black == Black
}

func (f Figure) IsLinearMover() bool {
	return uint8(f)&LinearMoverMask != 0
}

func (f Figure) IsLinearMoverFrom(color Color) bool {
	return uint8(f)&(LinearMoverMask|ColorMask) == LinearMoverMask|uint8(color)
}

func (f Figure) IsWhiteLinearMover() bool {
	return uint8(f)&(LinearMoverMask|ColorMask) == LinearMoverMask|White
}

func (f Figure) IsBlackLinearMover() bool {
	return uint8(f)&(LinearMoverMask|ColorMask) == LinearMoverMask|Black
}

func (f Figure) IsDiagonalMover() bool {
	return uint8(f)&DiagonalMoverMask != 0
}

func (f Figure) IsDiagonalMoverFrom(color Color) bool {
	return uint8(f)&(DiagonalMoverMask|ColorMask) == DiagonalMoverMask|uint8(color)
}

func (f Figure) IsWhiteDiagonalMover() bool {
	return uint8(f)&(DiagonalMoverMask|ColorMask) == DiagonalMoverMask|White
}

func (f Figure) IsBlackDiagonalMover() bool {
	return uint8(f)&(DiagonalMoverMask|ColorMask) == DiagonalMoverMask|Black
}

func (f Figure) IsKnight() bool {
	return f == WhiteKnight || f == BlackKnight
}

func (f Figure) IsRook() bool {
	return f == WhiteRook || f == BlackRook
}

func (f Figure) IsQueen() bool {
	return f == WhiteQueen || f == BlackQueen
}

func (f Figure) IsKing() bool {
	return f == WhiteKing || f == BlackKing
}

func (f Figure) IsNoFigure() bool {
	return f == NoFigure
}

func (f Figure) IsNoFigureOr(color Color) bool {
	return uint8(f)&ColorMask != uint8(color.Oponent())
}

func (f Figure) IsNoFigureOrNot(color Color) bool {
	return uint8(f)&ColorMask != uint8(color)
}

func (f Figure) IsNoFigureOrWhite() bool {
	return uint8(f)&Black == 0
}

func (f Figure) IsNoFigureOrBlack() bool {
	return uint8(f)&White == 0
}

func (f Figure) Symbol() string {
	switch f {
	case NoFigure:
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

func FromSymbol(symbol rune) Figure {
	switch symbol {
	case 32:
		return NoFigure
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
