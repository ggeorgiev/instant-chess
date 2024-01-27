package peace

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFigurePermutations(t *testing.T) {
	peaces := MustParseFigures("♚♔♖♖")

	iterator, perm := NewPermutationIterator(peaces)

	assert.Equal(t, uint64(12), iterator.NumberPermutations())

	assert.Equal(t, Figures{0x05, 0x05, 0x31, 0x32}, perm)
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
