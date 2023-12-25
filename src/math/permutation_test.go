package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestPermutationIterator tests the basic functionality of the PermutationIterator
func TestPermutation3_3Iterator(t *testing.T) {
	// Create an iterator for permutations of 3 elements
	iterator := NewPermutationIterator(3, 3)

	// Define the expected series of permutations
	expectedPermutations := [][]int{
		{0, 1, 2},
		{0, 2, 1},
		{1, 0, 2},
		{1, 2, 0},
		{2, 0, 1},
		{2, 1, 0},
	}

	// Iterate through the iterator and check against the expected permutations
	for _, expected := range expectedPermutations {
		perm := iterator.Next()
		assert.Equal(t, expected, perm)
	}

	assert.Nil(t, iterator.Next())
}

func TestPermutation4_2Iterator(t *testing.T) {
	iterator := NewPermutationIterator(4, 2)

	expectedPermutations := [][]int{
		{0, 1},
		{1, 0},
		{0, 2},
		{2, 0},
		{0, 3},
		{3, 0},
		{1, 2},
		{2, 1},
		{1, 3},
		{3, 1},
		{2, 3},
		{3, 2},
	}

	for _, expected := range expectedPermutations {
		perm := iterator.Next()
		assert.Equal(t, expected, perm)
	}

	assert.Nil(t, iterator.Next())
}
