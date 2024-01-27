package peace

import (
	"sort"
)

// PermutationIterator is an iterator for permutations
type PermutationIterator struct {
	codes       Codes
	permutation Codes
}

// NewPermutationIterator creates a new PermutationIterator
func NewPermutationIterator(codes Codes) (*PermutationIterator, Codes) {
	perm := make(Codes, len(codes))
	copy(perm, codes)

	sort.Sort(perm)

	return &PermutationIterator{
		codes:       codes,
		permutation: perm,
	}, perm
}

func (it *PermutationIterator) NumberPermutations() uint64 {
	totalPermutations := uint64(1)
	repCount := uint64(1)

	for i := 1; i < len(it.codes); i++ {
		totalPermutations *= uint64(i + 1)
		if it.permutation[i] == it.permutation[i-1] {
			repCount++
			totalPermutations /= repCount
		} else {
			repCount = 1
		}
	}

	return totalPermutations
}

// Next advances the iterator to the next unique permutation and returns it.
// If there are no more permutations, it returns nil.
func (it *PermutationIterator) Next() Codes {
	array := it.permutation
	// Find the largest index k such that a[k] < a[k + 1]
	var k int = -1
	for i := len(array) - 2; i >= 0; i-- {
		if array[i] < array[i+1] {
			k = i
			break
		}
	}
	if k == -1 {
		// This is the last permutation
		return nil
	}

	// Find the smallest element after index k which is greater than a[k]
	var l int
	for i := len(array) - 1; i > k; i-- {
		if array[k] < array[i] {
			l = i
			break
		}
	}

	// Swap a[k] and a[l]
	array[k], array[l] = array[l], array[k]

	// Reverse the sequence from a[k + 1] to the end
	for i, j := k+1, len(array)-1; i < j; i, j = i+1, j-1 {
		array[i], array[j] = array[j], array[i]
	}

	return array
}
