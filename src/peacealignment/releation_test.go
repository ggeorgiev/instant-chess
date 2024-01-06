package peacealignment

import (
	"testing"

	"github.com/ggeorgiev/instant-chess/src/square"
	"github.com/stretchr/testify/assert"
)

func (r Relations) String() string {
	result := ""
	for i := square.ZeroIndex; i <= square.LastIndex; i++ {
		result += "{\n"
		for j := square.ZeroIndex; j <= square.LastIndex; j++ {
			switch r[i][j] {
			case NotAligned:
				result += "na, "
			case SameSquare:
				result += "ss, "
			case FileAbove:
				result += "fa, "
			case FileUnder:
				result += "fu, "
			case RankLeft:
				result += "rl, "
			case RankRight:
				result += "rr, "
			case DiagonalAbove:
				result += "da, "
			case DiagonalUnder:
				result += "du, "
			case CounterDiagonalAbove:
				result += "ca, "
			case CounterDiagonalUnder:
				result += "cu, "
			}
			if j%8 == 7 {
				result += "\n"
			}
		}
		result += "},\n"
	}

	return result
}

func TestSquareRelations(t *testing.T) {
	squareRelations := squareRelationsInternalHelper()
	assert.Equal(t, squareRelations, SquareRelations, squareRelations.String())
}
