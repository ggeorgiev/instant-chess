package peace

import (
	"fmt"
	"strings"

	"github.com/ggeorgiev/instant-chess/src/move"
	"github.com/ggeorgiev/instant-chess/src/square"
	"github.com/ggeorgiev/instant-chess/src/util"
)

type Figure uint8
type Figures []Figure

// Provide the sort interface for Figures
func (figures Figures) Len() int           { return len(figures) }
func (figures Figures) Less(i, j int) bool { return figures[i] < figures[j] }
func (figures Figures) Swap(i, j int)      { figures[i], figures[j] = figures[j], figures[i] }

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

func (f Figure) IsBishop() bool {
	return f == WhiteBishop || f == BlackBishop
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

func FromSymbol(symbol rune) (Figure, error) {
	switch symbol {
	case 32:
		return NoFigure, nil
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

	return NoFigure, fmt.Errorf("unacceptable symbol: %d", symbol)
}

func ParseFigures(text string) (Figures, error) {
	runes := util.Runes(text)
	position := make([]int, len(runes))
	peaces := make(Figures, len(runes))
	for i, symbol := range runes {
		position[i] = 0
		var err error
		peaces[i], err = FromSymbol(symbol)
		if err != nil {
			return nil, err
		}
	}
	return peaces, nil
}

func MustParseFigures(text string) Figures {
	figures, err := ParseFigures(text)
	if err != nil {
		panic(err.Error())
	}
	return figures
}

func (fs Figures) MoveRights() move.RightsList {
	whiteKing := false
	blackKing := false
	whiteRooks := 0
	blackRooks := 0
	whitePawn := false
	blackPawn := false

	for _, f := range fs {
		switch f {
		case WhiteKing:
			whiteKing = true
		case BlackKing:
			blackKing = true
		case WhiteRook:
			whiteRooks++
		case BlackRook:
			blackRooks++
		case WhitePawn:
			whitePawn = true
		case BlackPawn:
			blackPawn = true
		}
	}

	whiteCasting := move.RightsList{move.NoRights}
	if whiteKing {
		if whiteRooks >= 1 {
			whiteCasting = append(whiteCasting, move.WhiteKingsideCastlingRights, move.WhiteQueensideCastlingRights)
		}
		if whiteRooks >= 2 {
			whiteCasting = append(whiteCasting, move.WhiteBothCastlingRights)
		}
	}

	blackCasting := move.RightsList{move.NoRights}
	if blackKing {
		if blackRooks >= 1 {
			blackCasting = append(blackCasting, move.BlackKingsideCastlingRights, move.BlackQueensideCastlingRights)
		}
		if blackRooks >= 2 {
			blackCasting = append(blackCasting, move.BlackBothCastlingRights)
		}
	}

	enPassantFiles := move.RightsList{move.NoEnPassantFile}
	if whitePawn && blackPawn {
		for f := square.ZeroFile; f <= square.LastFile; f++ {
			enPassantFiles = append(enPassantFiles, move.NoRights.SetEnPassantFile(f))
		}
	}

	return move.CombineRights(whiteCasting, blackCasting, enPassantFiles)
}

func (fs Figures) Copy() Figures {
	peaces := make(Figures, len(fs))
	copy(peaces, fs)
	return peaces
}

func (fs Figures) String() string {
	var sb strings.Builder
	for _, f := range fs {
		sb.WriteString(f.Symbol())
	}
	return sb.String()
}

func (fs Figures) RemoveOne(f Figure) Figures {
	for i, figure := range fs {
		if figure == f {
			newSize := len(fs) - 1
			fs[i] = fs[newSize]
			return fs[:newSize]
		}
	}
	panic("Removing figure that does not exists")
}
