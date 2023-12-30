package chess

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSquareX(t *testing.T) {
	for x := int8(0); x < 8; x++ {
		for y := int8(0); y < 8; y++ {
			assert.Equal(t, NewSquare(x, y).X(), x)
		}
	}
}

func TestSquareY(t *testing.T) {
	for x := int8(0); x < 8; x++ {
		for y := int8(0); y < 8; y++ {
			assert.Equal(t, NewSquare(x, y).Y(), y)
		}
	}
}
