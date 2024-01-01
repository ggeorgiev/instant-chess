package peacemoves

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKingMovesInternalHelper(t *testing.T) {
	assert.Equal(t, kingMovesInternalHelper(), KingSquareIndexes)
}
