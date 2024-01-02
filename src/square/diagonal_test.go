package square

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiagonal(t *testing.T) {
	assert.Equal(t, diagonalInternalHelper(), Diagonal)
}

func TestAntiDiagonal(t *testing.T) {
	assert.Equal(t, antiDiagonalInternalHelper(), AntiDiagonal)
}
