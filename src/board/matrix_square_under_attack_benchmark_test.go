package board

import (
	"testing"

	"github.com/ggeorgiev/instant-chess/src/square"
)

var states = []State{
	MustParseState(`
····a···b···c···d···e···f···g···h····
··+---+---+---+---+---+---+---+---+··
8·| ♜ | ♞ | ♝ | ♛ | ♚ | ♝ | ♞ | ♜ |·8
··+---+---+---+---+---+---+---+---+··
7·| ♟︎ | ♟︎ | ♟︎ | ♟︎ | ♟︎ | ♟︎ | ♟︎ | ♟︎ |·7
··+---+---+---+---+---+---+---+---+··
6·|   |   |   |   |   |   |   |   |·6
··+---+---+---+---+---+---+---+---+··
5·|   |   |   |   |   |   |   |   |·5
··+---+---+---+---+---+---+---+---+··
4·|   |   |   |   |   |   |   |   |·4
··+---+---+---+---+---+---+---+---+··
3·|   |   |   |   |   |   |   |   |·3
··+---+---+---+---+---+---+---+---+··
2·| ♙ | ♙ | ♙ | ♙ | ♙ | ♙ | ♙ | ♙ |·2
··+---+---+---+---+---+---+---+---+··
1·| ♖ | ♘ | ♗ | ♕ | ♔ | ♗ | ♘ | ♖ |·1
··+---+---+---+---+---+---+---+---+··
····a···b···c···d···e···f···g···h····
· O-O: --, O-O-O: --, En Passant: - ·
`),
	MustParseState(`
····a···b···c···d···e···f···g···h····
··+---+---+---+---+---+---+---+---+··
8·|   |   |   |   |   |   |   |   |·8
··+---+---+---+---+---+---+---+---+··
7·|   |   |   |   |   |   |   |   |·7
··+---+---+---+---+---+---+---+---+··
6·|   |   |   |   |   |   |   |   |·6
··+---+---+---+---+---+---+---+---+··
5·|   |   |   |   |   |   |   |   |·5
··+---+---+---+---+---+---+---+---+··
4·|   | ♜ |   | ♔ |   |   |   |   |·4
··+---+---+---+---+---+---+---+---+··
3·|   |   |   |   |   |   |   |   |·3
··+---+---+---+---+---+---+---+---+··
2·|   | ♚ |   |   |   |   |   |   |·2
··+---+---+---+---+---+---+---+---+··
1·|   |   |   |   |   |   |   |   |·1
··+---+---+---+---+---+---+---+---+··
····a···b···c···d···e···f···g···h····
· O-O: --, O-O-O: --, En Passant: - ·
`),
}

func BenchmarkUnderAttackIsKing_1k(b *testing.B) {
	for i := 1; i < 1000; i++ {
		for _, state := range states {
			for s := square.ZeroIndex; s <= square.LastIndex; s++ {
				state.Matrix.IsSquareUnderAttackFromWhiteKing(s)
			}
		}
	}
}
func BenchmarkUnderAttackCountKing_1k(b *testing.B) {
	for i := 1; i < 1000; i++ {
		for _, state := range states {
			for s := square.ZeroIndex; s <= square.LastIndex; s++ {
				state.Matrix.CountSquareUnderAttackFromWhiteKing(s)
			}
		}
	}
}

func BenchmarkUnderAttackIsKnight_1k(b *testing.B) {
	for i := 1; i < 1000; i++ {
		for _, state := range states {
			for s := square.ZeroIndex; s <= square.LastIndex; s++ {
				state.Matrix.IsSquareUnderAttackFromWhiteKnight(s)
			}
		}
	}
}

func BenchmarkUnderAttackCountKnight_1k(b *testing.B) {
	for i := 1; i < 1000; i++ {
		for _, state := range states {
			for s := square.ZeroIndex; s <= square.LastIndex; s++ {
				state.Matrix.CountSquareUnderAttackFromWhiteKnight(s)
			}
		}
	}
}

func BenchmarkUnderAttackIsPawn_1k(b *testing.B) {
	for i := 1; i < 1000; i++ {
		for _, state := range states {
			for s := square.ZeroIndex; s <= square.LastIndex; s++ {
				state.Matrix.IsSquareUnderAttackFromWhitePawn(s)
			}
		}
	}
}

func BenchmarkUnderAttackCountPawn_1k(b *testing.B) {
	for i := 1; i < 1000; i++ {
		for _, state := range states {
			for s := square.ZeroIndex; s <= square.LastIndex; s++ {
				state.Matrix.CountSquareUnderAttackFromWhitePawn(s)
			}
		}
	}
}

func BenchmarkUnderAttackIsDirect_1k(b *testing.B) {
	for i := 1; i < 1000; i++ {
		for _, state := range states {
			for s := square.ZeroIndex; s <= square.LastIndex; s++ {
				state.Matrix.IsSquareUnderDirectAttackFromWhite(s)
			}
		}
	}
}

func BenchmarkDirectAttackersOfSquare_1k(b *testing.B) {
	for i := 1; i < 1000; i++ {
		for _, state := range states {
			for s := square.ZeroIndex; s <= square.LastIndex; s++ {
				state.Matrix.DirectAttackersOfSquareFromWhite(s)
			}
		}
	}
}

func BenchmarkUnderAttack_1k(b *testing.B) {
	for i := 1; i < 1000; i++ {
		for _, state := range states {
			for s := square.ZeroIndex; s <= square.LastIndex; s++ {
				state.Matrix.IsSquareUnderAttackFromWhite(s)
			}
		}
	}
}

func BenchmarkAttackersOfSquare_1k(b *testing.B) {
	for i := 1; i < 1000; i++ {
		for _, state := range states {
			for s := square.ZeroIndex; s <= square.LastIndex; s++ {
				state.Matrix.BlockableAttackersOfSquareFromWhite(s)
			}
		}
	}
}
