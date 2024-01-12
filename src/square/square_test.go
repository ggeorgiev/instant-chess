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

func TestMin(t *testing.T) {
	assert.Equal(t, Indexes{}.Min(), InvalidIndex)
	assert.Equal(t, Indexes{InvalidIndex}.Min(), InvalidIndex)
	assert.Equal(t, Indexes{ZeroIndex}.Min(), ZeroIndex)
	assert.Equal(t, Indexes{ZeroIndex, LastIndex}.Min(), ZeroIndex)
	assert.Equal(t, Indexes{LastIndex, ZeroIndex}.Min(), ZeroIndex)
}

func TestMax(t *testing.T) {
	assert.Equal(t, Indexes{}.Max(), InvalidIndex)
	assert.Equal(t, Indexes{InvalidIndex}.Max(), InvalidIndex)
	assert.Equal(t, Indexes{ZeroIndex}.Max(), ZeroIndex)
	assert.Equal(t, Indexes{ZeroIndex, LastIndex}.Max(), LastIndex)
	assert.Equal(t, Indexes{LastIndex, ZeroIndex}.Max(), LastIndex)
}

func TestMaxBase(t *testing.T) {
	assert.Equal(t, Indexes{}.MaxBase(), InvalidIndex)
	assert.Equal(t, Indexes{InvalidIndex}.MaxBase(), InvalidIndex)
	assert.Equal(t, Indexes{ZeroIndex}.MaxBase(), ZeroIndex)
	assert.Equal(t, Indexes{ZeroIndex, LastIndex}.MaxBase(), LastIndex)
	assert.Equal(t, Indexes{LastIndex, ZeroIndex}.MaxBase(), LastIndex)
	assert.Equal(t, Indexes{LastIndex, LastIndex - 1, ZeroIndex}.MaxBase(), LastIndex-1)
	assert.Equal(t, Indexes{LastIndex - 2, LastIndex, LastIndex - 1, ZeroIndex}.MaxBase(), LastIndex-2)
}
