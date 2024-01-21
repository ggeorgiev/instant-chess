package matrix

import (
	"testing"

	"github.com/ggeorgiev/instant-chess/src/square"
)

func BenchmarkIsChecked_1k(b *testing.B) {
	for i := 1; i < 1000; i++ {
		for _, matrix := range matrixes {
			for s := square.ZeroIndex; s <= square.LastIndex; s++ {
				matrix.IsWhiteChecked(s)
			}
		}
	}
}

func BenchmarkIsCheckedToMove_1k(b *testing.B) {
	for i := 1; i < 1000; i++ {
		for _, matrix := range matrixes {
			for s := square.ZeroIndex; s <= square.LastIndex; s++ {
				matrix.IsWhiteCheckedToMoveCaptureOrBlock(s)
			}
		}
	}
}
