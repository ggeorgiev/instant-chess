package chess

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	Generate("♚♔♖♜")
	assert.False(t, true)
}
