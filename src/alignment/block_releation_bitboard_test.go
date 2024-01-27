package alignment

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBlockRelationBitboardMasksInternalHelper(t *testing.T) {
	maskLists := blockRelationBitboardMasksInternalHelper()
	assert.Equal(t, maskLists, BlockRelationMasksList, maskLists.String())
}
