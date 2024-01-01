package peaceattacks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKingBitboardMasksInternalHelper(t *testing.T) {
	masks := kingBitboardMasksInternalHelper()
	assert.Equal(t, masks, KingBitboardMasks, masks.String())
}
