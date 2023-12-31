package chess

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKingMovesInternalHelper(t *testing.T) {
	assert.Equal(t, kingMovesInternalHelper(), KingMoves)
}

func TestKnightMovesInternalHelper(t *testing.T) {
	assert.Equal(t, knightMovesInternalHelper(), KnightMoves)
}
