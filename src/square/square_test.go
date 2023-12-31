package square

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSquareX(t *testing.T) {
	for x := int8(0); x < 8; x++ {
		for y := int8(0); y < 8; y++ {
			assert.Equal(t, NewIndex(x, y).X(), x)
		}
	}
}

func TestSquareY(t *testing.T) {
	for x := int8(0); x < 8; x++ {
		for y := int8(0); y < 8; y++ {
			assert.Equal(t, NewIndex(x, y).Y(), y)
		}
	}
}
