package peace

import (
	"testing"

	"github.com/ggeorgiev/instant-chess/src/util"
	"github.com/stretchr/testify/assert"
)

func TestFigurePermutations(t *testing.T) {
	runes := util.Runes("♚♔♖♖")
	position := make([]int, len(runes))
	peaces := make(Figures, len(runes))
	for i, symbol := range runes {
		position[i] = 0
		peaces[i] = FromSymbol(symbol)
	}

	iterator := NewPermutationIterator(peaces)

	assert.Equal(t, uint64(12), iterator.NumberPermutations())

	assert.Equal(t, Figures{0x05, 0x05, 0x31, 0x32}, iterator.Next())
	assert.Equal(t, Figures{0x05, 0x05, 0x32, 0x31}, iterator.Next())
	assert.Equal(t, Figures{0x05, 0x31, 0x05, 0x32}, iterator.Next())
	assert.Equal(t, Figures{0x05, 0x31, 0x32, 0x05}, iterator.Next())
	assert.Equal(t, Figures{0x05, 0x32, 0x05, 0x31}, iterator.Next())
	assert.Equal(t, Figures{0x05, 0x32, 0x31, 0x05}, iterator.Next())
	assert.Equal(t, Figures{0x31, 0x05, 0x05, 0x32}, iterator.Next())
	assert.Equal(t, Figures{0x31, 0x05, 0x32, 0x05}, iterator.Next())
	assert.Equal(t, Figures{0x31, 0x32, 0x05, 0x05}, iterator.Next())
	assert.Equal(t, Figures{0x32, 0x05, 0x05, 0x31}, iterator.Next())
	assert.Equal(t, Figures{0x32, 0x05, 0x31, 0x05}, iterator.Next())
	assert.Equal(t, Figures{0x32, 0x31, 0x05, 0x05}, iterator.Next())
	assert.Equal(t, Figures(nil), iterator.Next())
}
