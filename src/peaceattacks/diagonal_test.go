package peaceattacks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRiseLeftSquareIndexesInternalHelper(t *testing.T) {
	indexes := riseLeftSquareIndexesInternalHelper()
	assert.Equal(t, indexes, RiseLeftSquareIndexes)
}

func TestRiseRightSquareIndexesInternalHelper(t *testing.T) {
	indexes := riseRightSquareIndexesInternalHelper()
	assert.Equal(t, indexes, RiseRightSquareIndexes)
}

func TestFallLeftSquareIndexesInternalHelper(t *testing.T) {
	indexes := fallLeftSquareIndexesInternalHelper()
	assert.Equal(t, indexes, FallLeftSquareIndexes)
}

func TestFallRightSquareIndexesInternalHelper(t *testing.T) {
	indexes := fallRightSquareIndexesInternalHelper()
	assert.Equal(t, indexes, FallRightSquareIndexes)
}

func TestRiseLeftBitboardMasksInternalHelper(t *testing.T) {
	masks := riseLeftBitboardMasksInternalHelper()
	assert.Equal(t, masks, RiseLeftBitboardMasks, masks.String())
}

func TestRiseRightBitboardMasksInternalHelper(t *testing.T) {
	masks := riseRightBitboardMasksInternalHelper()
	assert.Equal(t, masks, RiseRightBitboardMasks, masks.String())
}

func TestFallLeftBitboardMasksInternalHelper(t *testing.T) {
	masks := fallLeftBitboardMasksInternalHelper()
	assert.Equal(t, masks, FallLeftBitboardMasks, masks.String())
}

func TestFallRightBitboardMasksInternalHelper(t *testing.T) {
	masks := fallRightBitboardMasksInternalHelper()
	assert.Equal(t, masks, FallRightBitboardMasks, masks.String())
}

func TestDiagonalsRiseBitboardMasksInternalHelper(t *testing.T) {
	masks := diagonalsRiseBitboardMasksInternalHelper()
	assert.Equal(t, masks, DiagonalsRiseBitboardMasks, masks.String())
}

func TestDiagonalsFallBitboardMasksInternalHelper(t *testing.T) {
	masks := diagonalsFallBitboardMasksInternalHelper()
	assert.Equal(t, masks, DiagonalsFallBitboardMasks, masks.String())
}

func TestDiagonalsBitboardMasksInternalHelper(t *testing.T) {
	masks := diagonalsBitboardMasksInternalHelper()
	assert.Equal(t, masks, DiagonalsBitboardMasks, masks.String())
}
