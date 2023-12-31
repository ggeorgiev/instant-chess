package chess

import (
	"testing"
)

var positionA = ParsePosition(`
    a   b   c   d   e   f   g   h
  +---+---+---+---+---+---+---+---+
8 |   |   |   |   |   |   |   |   | 8
  +---+---+---+---+---+---+---+---+
7 |   |   |   |   |   |   |   |   | 7
  +---+---+---+---+---+---+---+---+
6 |   |   |   |   |   |   |   |   | 6
  +---+---+---+---+---+---+---+---+
5 |   |   |   |   |   |   |   |   | 5
  +---+---+---+---+---+---+---+---+
4 |   | ♜ |   | ♔ |   |   |   |   | 4
  +---+---+---+---+---+---+---+---+
3 |   |   |   |   |   |   |   |   | 3
  +---+---+---+---+---+---+---+---+
2 |   | ♚ |   |   |   |   |   |   | 2
  +---+---+---+---+---+---+---+---+
1 |   |   |   |   |   |   |   |   | 1
  +---+---+---+---+---+---+---+---+
    a   b   c   d   e   f   g   h
`)

func BenchmarkSquareUnderAttackXY_1k(b *testing.B) {
	for i := 1; i < 1000; i++ {
		for s := Square(0); s < 64; s++ {
			positionA.Board.SquareUnderAttackFromWhite(s)
		}
	}
}

func BenchmarkSquareUnderAttackYX_1k(b *testing.B) {
	for i := 1; i < 1000; i++ {
		for s := Square(0); s < 64; s++ {
			positionA.Board.SquareUnderAttackFromWhite(s)
		}
	}
}
