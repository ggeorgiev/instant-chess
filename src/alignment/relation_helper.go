package alignment

import (
	"github.com/ggeorgiev/instant-chess/src/square"
)

func squareRelationsInternalHelper() Relations {
	var relations Relations

	for i := square.ZeroIndex; i <= square.LastIndex; i++ {
		for j := square.ZeroIndex; j <= square.LastIndex; j++ {
			if i == j {
				relations[i][j] = SameSquare
				continue
			}
			if IsSameFile(i, j) {
				if i > j {
					relations[i][j] = FileUnder
				} else {
					relations[i][j] = FileAbove
				}
			} else if IsSameRank(i, j) {
				if i > j {
					relations[i][j] = RankLeft
				} else {
					relations[i][j] = RankRight
				}
			} else if IsSameDiagonal(i, j) {
				if i > j {
					relations[i][j] = DiagonalUnder
				} else {
					relations[i][j] = DiagonalAbove
				}
			} else if IsSameCounterDiagonal(i, j) {
				if i > j {
					relations[i][j] = CounterDiagonalUnder
				} else {
					relations[i][j] = CounterDiagonalAbove
				}
			} else {
				relations[i][j] = NotAligned
			}
		}
	}

	return relations
}
