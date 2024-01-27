package attack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKnightBitboardMasksInternalHelper(t *testing.T) {
	masks := knightBitboardMasksInternalHelper()
	assert.Equal(t, masks, KnightBitboardMasks, masks.String())
}
