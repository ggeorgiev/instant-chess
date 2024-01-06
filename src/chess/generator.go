package chess

import (
	"fmt"

	"github.com/ggeorgiev/instant-chess/src/math"
	"github.com/ggeorgiev/instant-chess/src/peace"
	"github.com/ggeorgiev/instant-chess/src/util"
)

func Generate(peacesString string) {
	runes := util.Runes(peacesString)
	position := make([]int, len(runes))
	peaces := make(peace.Figures, len(runes))

	for i, symbol := range runes {
		position[i] = 0
		peaces[i] = peace.FromSymbol(symbol)
	}

	// Create a new permutation iterator
	iterator := math.NewPermutationIterator(64, len(runes))
	i := 0
	invalid := 0
	m1 := 0
	for ; i < 100000000; i++ {
		perm := iterator.Next()
		if perm == nil {
			break // No more permutations
		}

		var matrix [64]peace.Figure
		for s, place := range perm {
			matrix[place] = peaces[s]
		}

		position := CreatePosition(matrix)
		if !position.Valid {
			invalid++
		} else if position.M1() {
			//position.Print()
			m1++
		}
	}
	fmt.Printf("M1 positions: %d/%d\n", m1, i)
	fmt.Printf("Invalid positions: %d/%d\n", invalid, i)
}
