package peacealignment

import (
	"testing"

	"github.com/ggeorgiev/instant-chess/src/square"
	"github.com/stretchr/testify/assert"
)

func (v Vectors) String() string {
	result := ""
	for i := square.ZeroIndex; i <= square.LastIndex; i++ {
		result += "{\n"
		for j := square.ZeroIndex; j <= square.LastIndex; j++ {
			switch v[i][j] {
			case NoVector:
				result += "n, "
			case File:
				result += "f, "
			case Rank:
				result += "r, "
			case Diagonal:
				result += "d, "
			case CounterDiagonal:
				result += "c, "
			}
			if j%8 == 7 {
				result += "\n"
			}
		}
		result += "},\n"
	}

	return result
}

func TestSquareVectors(t *testing.T) {
	squareVectors := squareVectorsInternalHelper()
	assert.Equal(t, squareVectors, SquareVectors, squareVectors.String())
}
