package chess

import (
	"fmt"

	"github.com/ggeorgiev/instant-chess/src/bitboard"
	"github.com/ggeorgiev/instant-chess/src/board/matrix"
	"github.com/ggeorgiev/instant-chess/src/board/state"
	"github.com/ggeorgiev/instant-chess/src/math"
	"github.com/ggeorgiev/instant-chess/src/peace"
	"github.com/ggeorgiev/instant-chess/src/square"
)

func Generate(peacesString string) {
	peaces := peace.MustParseFigures(peacesString)
	n := uint64(len(peaces))

	rights := 0
	states := 0
	invalid := 0
	m1 := 0
	skipped := 0
	iterator := peace.NewPermutationIterator(peaces)

	//permutations := iterator.NumberPermutations()

	for perm := iterator.Next(); perm != nil; perm = iterator.Next() {
		for i := uint64(0); i < math.CountBitsets(n); i++ {
			bitset := math.IndexToBitset(n, i)
			indexes := square.ConvertBitboardMaskIntoIndexes(bitboard.Mask(bitset))

			matrix := matrix.Matrix{}
			for p, s := range indexes {
				matrix[s] = perm[p]
			}

			invalidBoard, offender := matrix.Invalid()
			if invalidBoard {
				if offender != square.InvalidIndex {
					skipTo := math.NextValidIndex(n, i, bitset, uint64(offender))
					skip := int(skipTo - i)
					states += skip
					invalid += skip
					skipped += skip
					i = skipTo
				} else {
					invalid++
				}
				continue
			}

			rightsList := matrix.MoveRights()
			rights += len(rightsList) - 1

			for _, rights := range rightsList {
				states++

				boardState := state.State{
					Matrix: &matrix,
					Rights: rights,
				}

				if boardState.M1() {
					m1++
					//store.Set(permutations*i, storage.Data{MateMoves: 1})
				}
			}
		}
	}

	//fmt.Printf("Ratio: %.02f%%\n", store.Ratio()*100.0)

	fmt.Printf("M1 positions: %d/%d\n", m1, states)
	fmt.Printf("Rights variations: %d/%d\n", rights, states)
	fmt.Printf("Skipped positions: %d/%d\n", skipped, states)
	fmt.Printf("Invalid positions: %d/%d\n", invalid, states)
}

// M1 positions: 33616/15249024
// Skipped positions: 2094283/15249024
// Invalid positions: 2094283/15249024
