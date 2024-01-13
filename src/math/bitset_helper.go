package math

import (
	"math/big"
)

func binomialValue(n uint64, k uint64) uint64 {
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

func binomialInternalHelper() [64][]uint64 {
	b := [64][]uint64{}

	for i := uint64(0); i < 64; i++ {
		b[i] = make([]uint64, i+1)
		for j := uint64(0); j <= i; j++ {
			b[i][j] = binomialValue(i, j)
		}
	}
	return b
}

func binomialCountInternalHelper() [65]uint64 {
	b := [65]uint64{}

	for i := uint64(0); i <= 64; i++ {
		b[i] = binomialValue(64, i)
	}
	return b
}
