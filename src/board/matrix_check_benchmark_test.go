package board

import (
	"testing"

	"github.com/ggeorgiev/instant-chess/src/square"
)

func BenchmarkIsChecked_1k(b *testing.B) {
	for i := 1; i < 1000; i++ {
		for _, state := range states {
			for s := square.ZeroIndex; s <= square.LastIndex; s++ {
				state.Matrix.IsWhiteChecked(s)
			}
		}
	}
}

func BenchmarkIsCheckedToMove_1k(b *testing.B) {
	for i := 1; i < 1000; i++ {
		for _, state := range states {
			for s := square.ZeroIndex; s <= square.LastIndex; s++ {
				state.Matrix.IsWhiteCheckedToMoveCaptureOrBlock(s)
			}
		}
	}
}
