package alignment

import (
	"github.com/ggeorgiev/instant-chess/src/square"
)

func squareVectorsInternalHelper() Vectors {
	var vectors Vectors

	for i := square.ZeroIndex; i <= square.LastIndex; i++ {
		for j := square.ZeroIndex; j <= square.LastIndex; j++ {
			if i == j {
				vectors[i][j] = NoVector
				continue
			}
			if IsSameFile(i, j) {
				vectors[i][j] = File
			} else if IsSameRank(i, j) {
				vectors[i][j] = Rank
			} else if IsSameDiagonal(i, j) {
				vectors[i][j] = Diagonal
			} else if IsSameCounterDiagonal(i, j) {
				vectors[i][j] = CounterDiagonal
			} else {
				vectors[i][j] = NoVector
			}
		}
	}

	return vectors
}
