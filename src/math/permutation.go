package math

// PermutationIterator is an iterator for permutations of k elements from an n-element set
type PermutationIterator struct {
	combIterator *CombinationIterator
	currentPerm  []int
	currentComb  []int
}

// NewPermutationIterator creates a new iterator for k-permutations of n elements
func NewPermutationIterator(n, k int) *PermutationIterator {
	return &PermutationIterator{
		combIterator: NewCombinationIterator(n, k),
	}
}

// Next returns the next permutation, or nil if no more
func (it *PermutationIterator) Next() []int {
	// When there's no current combination or no more permutations of the current combination
	if it.currentComb == nil || !nextPermutation(it.currentPerm) {
		it.currentComb = it.combIterator.Next() // Get the next combination
		if it.currentComb == nil {
			return nil // No more combinations, hence no more permutations
		}
		it.currentPerm = make([]int, len(it.currentComb))
		copy(it.currentPerm, it.currentComb) // Reset the current permutation
	}
	result := make([]int, len(it.currentPerm))
	copy(result, it.currentPerm) // Return a copy of the current permutation
	return result
}

// nextPermutation generates the next permutation of the array in lexicographic order
func nextPermutation(perm []int) bool {
	// Find non-increasing suffix
	i := len(perm) - 2
	for i >= 0 && perm[i] >= perm[i+1] {
		i--
	}
	if i < 0 {
		return false
	}

	// Find successor to pivot
	j := len(perm) - 1
	for perm[j] <= perm[i] {
		j--
	}
	perm[i], perm[j] = perm[j], perm[i]

	// Reverse suffix
	for k, l := i+1, len(perm)-1; k < l; k, l = k+1, l-1 {
		perm[k], perm[l] = perm[l], perm[k]
	}
	return true
}
