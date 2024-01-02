package chess

import (
	"testing"

	"github.com/ggeorgiev/instant-chess/src/peace"
	"github.com/ggeorgiev/instant-chess/src/square"
)

var positionB = ParsePosition(`
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

func BenchmarkTos_1k(b *testing.B) {
	kingSquare := positionB.Board.FindPeace(peace.WhiteKing)
	for i := 1; i < 1000; i++ {
		for s := square.ZeroIndex; s <= square.LastIndex; s++ {
			positionB.Board.WhiteTos(s, kingSquare)
		}
	}
}

var positionC = ParsePosition(`
    a   b   c   d   e   f   g   h
  +---+---+---+---+---+---+---+---+
8 |   |   |   |   |   |   |   |   | 8
  +---+---+---+---+---+---+---+---+
7 |   |   |   |   |   |   |   |   | 7
  +---+---+---+---+---+---+---+---+
6 |   |   |   |   | ♚ |   |   |   | 6
  +---+---+---+---+---+---+---+---+
5 |   |   |   |   |   |   |   |   | 5
  +---+---+---+---+---+---+---+---+
4 |   |   |   |   |   |   |   |   | 4
  +---+---+---+---+---+---+---+---+
3 |   |   |   |   | ♔ |   |   |   | 3
  +---+---+---+---+---+---+---+---+
2 |   |   |   |   |   |   |   |   | 2
  +---+---+---+---+---+---+---+---+
1 | ♖ | ♘ | ♗ | ♕ |   | ♗ | ♘ | ♖ | 1
  +---+---+---+---+---+---+---+---+
    a   b   c   d   e   f   g   h
`)

func BenchmarkMoves_1k(b *testing.B) {
	for i := 1; i < 1000; i++ {
		positionC.Board.Moves()
	}
}
