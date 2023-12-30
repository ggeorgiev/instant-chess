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

func BenchmarkSquareUnderAttackXY(b *testing.B) {
	for x := 0; x < 8; x++ {
		for y := 0; y < 8; y++ {
			positionA.Board.SquareUnderAttack(x, y, PeaceColorWhite)
		}
	}
}

func BenchmarkSquareUnderAttackYX(b *testing.B) {
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			positionA.Board.SquareUnderAttack(x, y, PeaceColorWhite)
		}
	}
}
