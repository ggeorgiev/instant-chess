package square

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiagonals(t *testing.T) {
	assert.Equal(t, diagonalsInternalHelper(), IndexDiagonals)
}

func TestCounterDiagonals(t *testing.T) {
	assert.Equal(t, counterDiagonalsInternalHelper(), IndexCounterDiagonals)
}
