package move

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKnightMovesInternalHelper(t *testing.T) {
	assert.Equal(t, knightMovesInternalHelper(), KnightSquareIndexes)
}
