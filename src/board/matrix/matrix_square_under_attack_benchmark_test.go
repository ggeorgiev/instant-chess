package matrix

import (
	"testing"

	"github.com/ggeorgiev/instant-chess/src/square"
)

var matrixes = []Matrix{
	MustParseMatrix(`
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
	MustParseMatrix(`
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
		for _, matrix := range matrixes {
			for s := square.ZeroIndex; s <= square.LastIndex; s++ {
				matrix.IsSquareUnderAttackFromWhiteKing(s)
			}
		}
	}
}
func BenchmarkUnderAttackCountKing_1k(b *testing.B) {
	for i := 1; i < 1000; i++ {
		for _, matrix := range matrixes {
			for s := square.ZeroIndex; s <= square.LastIndex; s++ {
				matrix.CountSquareUnderAttackFromWhiteKing(s)
			}
		}
	}
}

func BenchmarkUnderAttackIsKnight_1k(b *testing.B) {
	for i := 1; i < 1000; i++ {
		for _, matrix := range matrixes {
			for s := square.ZeroIndex; s <= square.LastIndex; s++ {
				matrix.IsSquareUnderAttackFromWhiteKnight(s)
			}
		}
	}
}

func BenchmarkUnderAttackCountKnight_1k(b *testing.B) {
	for i := 1; i < 1000; i++ {
		for _, matrix := range matrixes {
			for s := square.ZeroIndex; s <= square.LastIndex; s++ {
				matrix.CountSquareUnderAttackFromWhiteKnight(s)
			}
		}
	}
}

func BenchmarkUnderAttackIsPawn_1k(b *testing.B) {
	for i := 1; i < 1000; i++ {
		for _, matrix := range matrixes {
			for s := square.ZeroIndex; s <= square.LastIndex; s++ {
				matrix.IsSquareUnderAttackFromWhitePawn(s)
			}
		}
	}
}

func BenchmarkUnderAttackCountPawn_1k(b *testing.B) {
	for i := 1; i < 1000; i++ {
		for _, matrix := range matrixes {
			for s := square.ZeroIndex; s <= square.LastIndex; s++ {
				matrix.CountSquareUnderAttackFromWhitePawn(s)
			}
		}
	}
}

func BenchmarkUnderAttackIsDirect_1k(b *testing.B) {
	for i := 1; i < 1000; i++ {
		for _, matrix := range matrixes {
			for s := square.ZeroIndex; s <= square.LastIndex; s++ {
				matrix.IsSquareUnderDirectAttackFromWhite(s)
			}
		}
	}
}

func BenchmarkDirectAttackersOfSquare_1k(b *testing.B) {
	for i := 1; i < 1000; i++ {
		for _, matrix := range matrixes {
			for s := square.ZeroIndex; s <= square.LastIndex; s++ {
				matrix.DirectAttackersOfSquareFromWhite(s)
			}
		}
	}
}

func BenchmarkUnderAttack_1k(b *testing.B) {
	for i := 1; i < 1000; i++ {
		for _, matrix := range matrixes {
			for s := square.ZeroIndex; s <= square.LastIndex; s++ {
				matrix.IsSquareUnderAttackFromWhite(s)
			}
		}
	}
}

func BenchmarkAttackersOfSquare_1k(b *testing.B) {
	for i := 1; i < 1000; i++ {
		for _, matrix := range matrixes {
			for s := square.ZeroIndex; s <= square.LastIndex; s++ {
				matrix.BlockableAttackersOfSquareFromWhite(s)
			}
		}
	}
}
