package peaceattacks

import (
	"fmt"
	"testing"

	"github.com/ggeorgiev/instant-chess/src/peace"
	"github.com/ggeorgiev/instant-chess/src/square"
	"github.com/stretchr/testify/assert"
)

func TestWhiteDirectsInternalHelper(t *testing.T) {
	directsList := whiteDirectInternalHelper()
	assert.Equal(t, directsList, WhiteDirectsList, directsList.String())
}

func TestBlackDirectsInternalHelper(t *testing.T) {
	directsList := blackDirectInternalHelper()
	assert.Equal(t, directsList, BlackDirectsList, directsList.String())
}

func (dl DirectsList) String() string {
	result := ""
	for s := square.ZeroIndex; s <= square.LastIndex; s++ {
		result += "{"
		for _, direct := range dl[s] {
			shortPeace := ""
			switch direct.Peace {
			case peace.WhitePawn:
				shortPeace = "wp"
			case peace.WhiteKnight:
				shortPeace = "wn"
			case peace.WhiteKing:
				shortPeace = "wk"
			case peace.BlackPawn:
				shortPeace = "bp"
			case peace.BlackKnight:
				shortPeace = "bn"
			case peace.BlackKing:
				shortPeace = "bk"

			}
			result += fmt.Sprintf("{%d, %s},", direct.Index, shortPeace)
		}
		result += "},\n"
	}

	return result
}
