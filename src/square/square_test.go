package square

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSquareX(t *testing.T) {
	for f := File(0); f < 8; f++ {
		for r := Rank(0); r < 8; r++ {
			assert.Equal(t, NewIndex(f, r).File(), f)
		}
	}
}

func TestSquareY(t *testing.T) {
	for f := File(0); f < 8; f++ {
		for r := Rank(0); r < 8; r++ {
			assert.Equal(t, NewIndex(f, r).Rank(), r)
		}
	}
}
