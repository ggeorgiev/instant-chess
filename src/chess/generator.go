package chess

import (
	"fmt"

	"github.com/ggeorgiev/instant-chess/src/math"
)

func Generate(peacesString string) {
	runes := runes(peacesString)
	position := make([]int, len(runes))
	peaces := make([]Peace, len(runes))

	for i, symbol := range runes {
		position[i] = 0
		peaces[i] = PeaceFromSymbol(symbol)
	}

	// Create a new permutation iterator
	iterator := math.NewPermutationIterator(64, len(runes))
	i := 0
	invalid := 0
	for ; i < 1000000000; i++ {
		perm := iterator.Next()
		if perm == nil {
			break // No more permutations
		}

		var matrix [8][8]Peace
		for s, place := range perm {
			matrix[place%8][place/8] = peaces[s]
		}

		position := CreatePosition(matrix)
		if !position.Valid {
			invalid++
		}
	}
	fmt.Printf("Invalid positions: %d/%d\n", invalid, i)
}
