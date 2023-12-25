package math

// CombinationIterator is an iterator for k-combinations of an n-element set
type CombinationIterator struct {
	comb     []int // Current combination
	n, k     int   // Set size and subset size
	finished bool  // Finished flag
}

// NewCombinationIterator creates a new iterator for k-combinations of n elements
func NewCombinationIterator(n, k int) *CombinationIterator {
	comb := make([]int, k)
	for i := range comb {
		comb[i] = i
	}
	return &CombinationIterator{n: n, k: k, comb: comb}
}

// Next returns the next k-combination of n elements, or nil if no more
func (it *CombinationIterator) Next() []int {
	if it.finished {
		return nil // No more combinations
	}
	// Copy the current combination to return it
	result := make([]int, it.k)
	copy(result, it.comb)

	// Generate the next combination
	k := len(it.comb)
	i := k - 1
	for i >= 0 && it.comb[i] == it.n-k+i {
		i--
	}
	if i < 0 {
		it.finished = true // No more combinations
		return result
	}
	it.comb[i]++
	for j := i + 1; j < k; j++ {
		it.comb[j] = it.comb[j-1] + 1
	}
	return result
}
