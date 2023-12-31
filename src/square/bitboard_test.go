package square

import (
	"testing"

	"github.com/ggeorgiev/instant-chess/src/bitboard"
	"github.com/stretchr/testify/assert"
)

func TestPrintMask(t *testing.T) {
	mask := bitboard.Mask(0b0000000001100110111111111111111101111110001111000001100000000000)
	actual := "\n" + SprintMask(mask)
	expected := `
····a···b···c···d···e···f···g···h····
··+---+---+---+---+---+---+---+---+··
8·|   |   |   |   |   |   |   |   |·8
··+---+---+---+---+---+---+---+---+··
7·|   | ● | ● |   |   | ● | ● |   |·7
··+---+---+---+---+---+---+---+---+··
6·| ● | ● | ● | ● | ● | ● | ● | ● |·6
··+---+---+---+---+---+---+---+---+··
5·| ● | ● | ● | ● | ● | ● | ● | ● |·5
··+---+---+---+---+---+---+---+---+··
4·|   | ● | ● | ● | ● | ● | ● |   |·4
··+---+---+---+---+---+---+---+---+··
3·|   |   | ● | ● | ● | ● |   |   |·3
··+---+---+---+---+---+---+---+---+··
2·|   |   |   | ● | ● |   |   |   |·2
··+---+---+---+---+---+---+---+---+··
1·|   |   |   |   |   |   |   |   |·1
··+---+---+---+---+---+---+---+---+··
····a···b···c···d···e···f···g···h····
`
	assert.Equal(t, expected, actual)
}

func TestConvertBitboardMaskIntoIndexes(t *testing.T) {
	for s := ZeroIndex; s <= LastIndex; s++ {
		assert.Equal(t, Indexes{s}, ConvertBitboardMaskIntoIndexes(IndexMask[s]))
	}

	assert.Equal(t, Indexes(nil), ConvertBitboardMaskIntoIndexes(bitboard.Empty))
	assert.Equal(t, Indexes{5, 10}, ConvertBitboardMaskIntoIndexes(IndexMask[5]|IndexMask[10]))
	assert.Equal(t, Indexes{0, 17, 23, 63},
		ConvertBitboardMaskIntoIndexes(IndexMask[0]|IndexMask[17]|IndexMask[23]|IndexMask[63]))
}
