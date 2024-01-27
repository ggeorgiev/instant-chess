package peace

import (
	"testing"
)

func BenchmarkPeaceColor_1m(b *testing.B) {
	for i := 1; i < 1000000; i++ {
		Null.Color()

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
		Null.IsWhite()

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
		Null.IsBlack()

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

func BenchmarkPeaceNullOr_1m(b *testing.B) {
	for i := 1; i < 1000000; i++ {
		color := []Color{WhiteColor, BlackColor}[i%2]
		Null.IsNullOr(color)

		WhitePawn.IsNullOr(color)
		WhiteBishop.IsNullOr(color)
		WhiteKnight.IsNullOr(color)
		WhiteRook.IsNullOr(color)
		WhiteQueen.IsNullOr(color)
		WhiteKing.IsNullOr(color)

		BlackPawn.IsNullOr(color)
		BlackBishop.IsNullOr(color)
		BlackKnight.IsNullOr(color)
		BlackRook.IsNullOr(color)
		BlackQueen.IsNullOr(color)
		BlackKing.IsNullOr(color)
	}
}

func BenchmarkPeaceNullOrNot_1m(b *testing.B) {
	for i := 1; i < 1000000; i++ {
		color := []Color{WhiteColor, BlackColor}[i%2]

		Null.IsNullOrNot(color)

		WhitePawn.IsNullOrNot(color)
		WhiteBishop.IsNullOrNot(color)
		WhiteKnight.IsNullOrNot(color)
		WhiteRook.IsNullOrNot(color)
		WhiteQueen.IsNullOrNot(color)
		WhiteKing.IsNullOrNot(color)

		BlackPawn.IsNullOrNot(color)
		BlackBishop.IsNullOrNot(color)
		BlackKnight.IsNullOrNot(color)
		BlackRook.IsNullOrNot(color)
		BlackQueen.IsNullOrNot(color)
		BlackKing.IsNullOrNot(color)
	}
}

func BenchmarkPeaceIsRook_1m(b *testing.B) {
	for i := 1; i < 1000000; i++ {
		Null.IsRook()

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
		Null.IsLinearMover()

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

		Null.IsLinearMoverFrom(color)

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
