package attack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhitePawnSquareIndexesInternalHelper(t *testing.T) {
	indexes := whitePawnSquareIndexesInternalHelper()
	assert.Equal(t, indexes, FromWhitePawn)
}

func TestBlackPawnSquareIndexesInternalHelper(t *testing.T) {
	indexes := blackPawnSquareIndexesInternalHelper()
	assert.Equal(t, indexes, FromBlackPawn)
}

func TestWhitePawnBitboardMasksInternalHelper(t *testing.T) {
	masks := whitePawnBitboardMasksInternalHelper()
	assert.Equal(t, masks, WhitePawnBitboardMasks, masks.String())
}

func TestBlackPawnBitboardMasksInternalHelper(t *testing.T) {
	masks := blackPawnBitboardMasksInternalHelper()
	assert.Equal(t, masks, BlackPawnBitboardMasks, masks.String())
}
