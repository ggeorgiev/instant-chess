package chess

import (
	"testing"
)

func BenchmarkPeaceColor_1m(b *testing.B) {
	for i := 1; i < 1000000; i++ {
		Empty.Color()

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
		Empty.IsWhite()

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
		Empty.IsBlack()

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

func BenchmarkPeaceIsEmptyOr_1m(b *testing.B) {
	for i := 1; i < 1000000; i++ {
		color := []PeaceColor{PeaceColorWhite, PeaceColorBlack}[i%2]
		Empty.IsEmptyOr(color)

		WhitePawn.IsEmptyOr(color)
		WhiteBishop.IsEmptyOr(color)
		WhiteKnight.IsEmptyOr(color)
		WhiteRook.IsEmptyOr(color)
		WhiteQueen.IsEmptyOr(color)
		WhiteKing.IsEmptyOr(color)

		BlackPawn.IsEmptyOr(color)
		BlackBishop.IsEmptyOr(color)
		BlackKnight.IsEmptyOr(color)
		BlackRook.IsEmptyOr(color)
		BlackQueen.IsEmptyOr(color)
		BlackKing.IsEmptyOr(color)
	}
}

func BenchmarkPeaceIsEmptyOrNot_1m(b *testing.B) {
	for i := 1; i < 1000000; i++ {
		color := []PeaceColor{PeaceColorWhite, PeaceColorBlack}[i%2]

		Empty.IsEmptyOrNot(color)

		WhitePawn.IsEmptyOrNot(color)
		WhiteBishop.IsEmptyOrNot(color)
		WhiteKnight.IsEmptyOrNot(color)
		WhiteRook.IsEmptyOrNot(color)
		WhiteQueen.IsEmptyOrNot(color)
		WhiteKing.IsEmptyOrNot(color)

		BlackPawn.IsEmptyOrNot(color)
		BlackBishop.IsEmptyOrNot(color)
		BlackKnight.IsEmptyOrNot(color)
		BlackRook.IsEmptyOrNot(color)
		BlackQueen.IsEmptyOrNot(color)
		BlackKing.IsEmptyOrNot(color)
	}
}

func BenchmarkPeaceIsRook_1m(b *testing.B) {
	for i := 1; i < 1000000; i++ {
		Empty.IsRook()

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

func BenchmarkPeaceIsLiniar_1m(b *testing.B) {
	for i := 1; i < 1000000; i++ {
		Empty.IsLiniar()

		WhitePawn.IsLiniar()
		WhiteBishop.IsLiniar()
		WhiteKnight.IsLiniar()
		WhiteRook.IsLiniar()
		WhiteQueen.IsLiniar()
		WhiteKing.IsLiniar()

		BlackPawn.IsLiniar()
		BlackBishop.IsLiniar()
		BlackKnight.IsLiniar()
		BlackRook.IsLiniar()
		BlackQueen.IsLiniar()
		BlackKing.IsLiniar()
	}
}

func BenchmarkPeaceIsLiniarFrom_1m(b *testing.B) {
	for i := 1; i < 1000000; i++ {
		color := []PeaceColor{PeaceColorWhite, PeaceColorBlack}[i%2]

		Empty.IsLiniarFrom(color)

		WhitePawn.IsLiniarFrom(color)
		WhiteBishop.IsLiniarFrom(color)
		WhiteKnight.IsLiniarFrom(color)
		WhiteRook.IsLiniarFrom(color)
		WhiteQueen.IsLiniarFrom(color)
		WhiteKing.IsLiniarFrom(color)

		BlackPawn.IsLiniarFrom(color)
		BlackBishop.IsLiniarFrom(color)
		BlackKnight.IsLiniarFrom(color)
		BlackRook.IsLiniarFrom(color)
		BlackQueen.IsLiniarFrom(color)
		BlackKing.IsLiniarFrom(color)
	}
}
