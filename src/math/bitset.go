package math

import (
	"math/big"
	"math/bits"
)

func binomial(n uint64, k uint64) uint64 {
	if k > n-k {
		k = n - k
	}
	result := big.NewInt(1)
	for i := uint64(0); i < k; i++ {
		result.Mul(result, big.NewInt(int64(n-i)))
		result.Div(result, big.NewInt(int64(i+1)))
	}
	return result.Uint64()
}

func IndexToBitset(n uint64, index uint64) uint64 {
	bitset := uint64(0)

	for i := uint64(0); n > 0; i++ {
		b := binomial(63-i, n-1)
		if index >= b {
			index -= b
		} else {
			bitset |= 1 << i
			n--
		}
	}

	return bitset
}

func BitsetToIndex(n uint64, bitset uint64) uint64 {
	index := uint64(0)
	for i := uint64(0); n > 0; i++ {
		if bitset&(1<<i) != 0 {
			n--
		} else {
			index += binomial(63-i, n-1)
		}
	}
	return index
}

func CountBitsets(n uint64) uint64 {
	if n == 64 {
		return 1
	}
	return binomial(64, n)
}

func NextValidIndex(n uint64, bitset uint64, offendingBitPosition uint64) uint64 {
	sequenceOfOnes := bitset >> offendingBitPosition
	if sequenceOfOnes&(sequenceOfOnes+1) != 0 {
		panic("Incorrectly identified offending bit, this is not the first time it is in that position")
	}

	sequenceOfOnes >>= 1
	countSetBits := bits.OnesCount64(sequenceOfOnes)

	shiftedBits := sequenceOfOnes << (64 - countSetBits)

	offendingBitMask := uint64(1 << offendingBitPosition)
	rightMask := (offendingBitMask - 1)

	result := (bitset & rightMask) | shiftedBits | offendingBitMask

	return BitsetToIndex(n, result)
}
