package chess

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestkingMovesInternalHelper(t *testing.T) {
	assert.Equal(t, kingMovesInternalHelper(), KingMoves)
}
