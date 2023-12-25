package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestCombinationIterator tests the basic functionality of the CombinationIterator
func TestCombinationIterator(t *testing.T) {
	// Create an iterator for combinations of 4 elements taken 2 at a time
	iterator := NewCombinationIterator(4, 2)

	// Define the expected series of combinations
	expectedCombinations := [][]int{
		{0, 1},
		{0, 2},
		{0, 3},
		{1, 2},
		{1, 3},
		{2, 3},
	}

	for _, expected := range expectedCombinations {
		combo := iterator.Next()
		assert.Equal(t, expected, combo)
	}

	assert.Nil(t, iterator.Next())
}
