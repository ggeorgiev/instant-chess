package square

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSquareX(t *testing.T) {
	for f := ZeroFile; f <= LastFile; f++ {
		for r := ZeroRank; r <= LastRank; r++ {
			assert.Equal(t, NewIndex(f, r).File(), f)
		}
	}
}

func TestSquareY(t *testing.T) {
	for f := ZeroFile; f <= LastFile; f++ {
		for r := ZeroRank; r <= LastRank; r++ {
			assert.Equal(t, NewIndex(f, r).Rank(), r)
		}
	}
}
