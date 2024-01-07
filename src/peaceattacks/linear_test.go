package peaceattacks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRiseSquareIndexesInternalHelper(t *testing.T) {
	indexes := riseSquareIndexesInternalHelper()
	assert.Equal(t, indexes, RiseSquareIndexes)
}

func TestRightSquareIndexesInternalHelper(t *testing.T) {
	indexes := rightSquareIndexesInternalHelper()
	assert.Equal(t, indexes, RightSquareIndexes)
}

func TestFallSquareIndexesInternalHelper(t *testing.T) {
	indexes := fallSquareIndexesInternalHelper()
	assert.Equal(t, indexes, FallSquareIndexes)
}

func TestLeftSquareIndexesInternalHelper(t *testing.T) {
	indexes := leftSquareIndexesInternalHelper()
	assert.Equal(t, indexes, LeftSquareIndexes)
}

func TestRiseBitboardMasksInternalHelper(t *testing.T) {
	masks := riseBitboardMasksInternalHelper()
	assert.Equal(t, masks, RiseBitboardMasks, masks.String())
}

func TestRightBitboardMasksInternalHelper(t *testing.T) {
	masks := rightBitboardMasksInternalHelper()
	assert.Equal(t, masks, RightBitboardMasks, masks.String())
}

func TestFallBitboardMasksInternalHelper(t *testing.T) {
	masks := fallBitboardMasksInternalHelper()
	assert.Equal(t, masks, FallBitboardMasks, masks.String())
}

func TestLeftBitboardMasksInternalHelper(t *testing.T) {
	masks := leftBitboardMasksInternalHelper()
	assert.Equal(t, masks, LeftBitboardMasks, masks.String())
}

func TestLinearsRiseRightBitboardMasksInternalHelper(t *testing.T) {
	masks := linearsRiseRightBitboardMasksInternalHelper()
	assert.Equal(t, masks, LinearsRiseRightBitboardMasks, masks.String())
}

func TestLinearsFallLeftBitboardMasksInternalHelper(t *testing.T) {
	masks := linearsFallLeftBitboardMasksInternalHelper()
	assert.Equal(t, masks, LinearsFallLeftBitboardMasks, masks.String())
}

func TestLinearsBitboardMasksInternalHelper(t *testing.T) {
	masks := linearsBitboardMasksInternalHelper()
	assert.Equal(t, masks, LinearsBitboardMasks, masks.String())
}
