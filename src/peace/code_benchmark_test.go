package peace

import (
	"testing"
)

func BenchmarkPeaceColor_1m(b *testing.B) {
	for i := 1; i < 1000000; i++ {
		NoFigure.Color()

		WhitePawn.Color()
		WhiteBishop.Color()
		WhiteKnight.Color()
		WhiteRook.Color()
		WhiteQueen.Color()
		WhiteKing.Color()

		BlackPawn.Color()
		BlackBishop.Color()
		BlackKnight.Color()
		BlackRook.Color()
		BlackQueen.Color()
		BlackKing.Color()
	}
}

func BenchmarkPeaceIsWhite_1m(b *testing.B) {
	for i := 1; i < 1000000; i++ {
		NoFigure.IsWhite()

		WhitePawn.IsWhite()
		WhiteBishop.IsWhite()
		WhiteKnight.IsWhite()
		WhiteRook.IsWhite()
		WhiteQueen.IsWhite()
		WhiteKing.IsWhite()

		BlackPawn.IsWhite()
		BlackBishop.IsWhite()
		BlackKnight.IsWhite()
		BlackRook.IsWhite()
		BlackQueen.IsWhite()
		BlackKing.IsWhite()
	}
}

func BenchmarkPeaceIsBlack_1m(b *testing.B) {
	for i := 1; i < 1000000; i++ {
		NoFigure.IsBlack()

		WhitePawn.IsBlack()
		WhiteBishop.IsBlack()
		WhiteKnight.IsBlack()
		WhiteRook.IsBlack()
		WhiteQueen.IsBlack()
		WhiteKing.IsBlack()

		BlackPawn.IsBlack()
		BlackBishop.IsBlack()
		BlackKnight.IsBlack()
		BlackRook.IsBlack()
		BlackQueen.IsBlack()
		BlackKing.IsBlack()
	}
}

func BenchmarkPeaceNoFigureOr_1m(b *testing.B) {
	for i := 1; i < 1000000; i++ {
		color := []Color{WhiteColor, BlackColor}[i%2]
		NoFigure.IsNoFigureOr(color)

		WhitePawn.IsNoFigureOr(color)
		WhiteBishop.IsNoFigureOr(color)
		WhiteKnight.IsNoFigureOr(color)
		WhiteRook.IsNoFigureOr(color)
		WhiteQueen.IsNoFigureOr(color)
		WhiteKing.IsNoFigureOr(color)

		BlackPawn.IsNoFigureOr(color)
		BlackBishop.IsNoFigureOr(color)
		BlackKnight.IsNoFigureOr(color)
		BlackRook.IsNoFigureOr(color)
		BlackQueen.IsNoFigureOr(color)
		BlackKing.IsNoFigureOr(color)
	}
}

func BenchmarkPeaceNoFigureOrNot_1m(b *testing.B) {
	for i := 1; i < 1000000; i++ {
		color := []Color{WhiteColor, BlackColor}[i%2]

		NoFigure.IsNoFigureOrNot(color)

		WhitePawn.IsNoFigureOrNot(color)
		WhiteBishop.IsNoFigureOrNot(color)
		WhiteKnight.IsNoFigureOrNot(color)
		WhiteRook.IsNoFigureOrNot(color)
		WhiteQueen.IsNoFigureOrNot(color)
		WhiteKing.IsNoFigureOrNot(color)

		BlackPawn.IsNoFigureOrNot(color)
		BlackBishop.IsNoFigureOrNot(color)
		BlackKnight.IsNoFigureOrNot(color)
		BlackRook.IsNoFigureOrNot(color)
		BlackQueen.IsNoFigureOrNot(color)
		BlackKing.IsNoFigureOrNot(color)
	}
}

func BenchmarkPeaceIsRook_1m(b *testing.B) {
	for i := 1; i < 1000000; i++ {
		NoFigure.IsRook()

		WhitePawn.IsRook()
		WhiteBishop.IsRook()
		WhiteKnight.IsRook()
		WhiteRook.IsRook()
		WhiteQueen.IsRook()
		WhiteKing.IsRook()

		BlackPawn.IsRook()
		BlackBishop.IsRook()
		BlackKnight.IsRook()
		BlackRook.IsRook()
		BlackQueen.IsRook()
		BlackKing.IsRook()
	}
}

func BenchmarkPeaceIsLinearMover_1m(b *testing.B) {
	for i := 1; i < 1000000; i++ {
		NoFigure.IsLinearMover()

		WhitePawn.IsLinearMover()
		WhiteBishop.IsLinearMover()
		WhiteKnight.IsLinearMover()
		WhiteRook.IsLinearMover()
		WhiteQueen.IsLinearMover()
		WhiteKing.IsLinearMover()

		BlackPawn.IsLinearMover()
		BlackBishop.IsLinearMover()
		BlackKnight.IsLinearMover()
		BlackRook.IsLinearMover()
		BlackQueen.IsLinearMover()
		BlackKing.IsLinearMover()
	}
}

func BenchmarkPeaceIsLiniarFrom_1m(b *testing.B) {
	for i := 1; i < 1000000; i++ {
		color := []Color{WhiteColor, BlackColor}[i%2]

		NoFigure.IsLinearMoverFrom(color)

		WhitePawn.IsLinearMoverFrom(color)
		WhiteBishop.IsLinearMoverFrom(color)
		WhiteKnight.IsLinearMoverFrom(color)
		WhiteRook.IsLinearMoverFrom(color)
		WhiteQueen.IsLinearMoverFrom(color)
		WhiteKing.IsLinearMoverFrom(color)

		BlackPawn.IsLinearMoverFrom(color)
		BlackBishop.IsLinearMoverFrom(color)
		BlackKnight.IsLinearMoverFrom(color)
		BlackRook.IsLinearMoverFrom(color)
		BlackQueen.IsLinearMoverFrom(color)
		BlackKing.IsLinearMoverFrom(color)
	}
}
