package chess

import (
	"testing"

	"github.com/ggeorgiev/instant-chess/src/square"
)

func BenchmarkAttackMask_1k(b *testing.B) {
	for i := 1; i < 1000; i++ {
		for s := square.ZeroIndex; s <= square.LastIndex; s++ {
			positionA.Board.AttackBitboardMaskFromWhite()
		}
	}
}
