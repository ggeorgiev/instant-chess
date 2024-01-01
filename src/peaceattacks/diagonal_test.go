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

func TestRiseBitboardMasksInternalHelper(t *testing.T) {
	masks := riseBitboardMasksInternalHelper()
	assert.Equal(t, masks, RiseBitboardMasks, masks.String())
}

func TestFallBitboardMasksInternalHelper(t *testing.T) {
	masks := fallBitboardMasksInternalHelper()
	assert.Equal(t, masks, FallBitboardMasks, masks.String())
}
