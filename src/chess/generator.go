package chess

import (
	"fmt"

	"github.com/ggeorgiev/instant-chess/src/bitboard"
	"github.com/ggeorgiev/instant-chess/src/board"
	"github.com/ggeorgiev/instant-chess/src/math"
	"github.com/ggeorgiev/instant-chess/src/peace"
	"github.com/ggeorgiev/instant-chess/src/square"
	"github.com/ggeorgiev/instant-chess/src/util"
)

func Generate(peacesString string) {
	runes := util.Runes("♚♔♖♖")
	position := make([]int, len(runes))
	peaces := make(peace.Figures, len(runes))
	for i, symbol := range runes {
		position[i] = 0
		peaces[i] = peace.FromSymbol(symbol)
	}
	n := uint64(len(peaces))

	invalid := 0
	m1 := 0
	skipped := 0
	iterator := peace.NewPermutationIterator(peaces)
	for perm := iterator.Next(); perm != nil; perm = iterator.Next() {
		for i := uint64(0); i < math.CountBitsets(n); i++ {
			bitset := math.IndexToBitset(n, i)
			indexes := square.ConvertBitboardMaskIntoIndexes(bitboard.Mask(bitset))
			var boardState board.State
			for p, s := range indexes {
				boardState.Matrix[s] = peaces[p]
			}

			invalidBoard, offender := boardState.Invalid()

			if invalidBoard {
				if offender != square.InvalidIndex && offender != indexes[len(indexes)-1] {
					skipTo := math.NextValidIndex(n, bitset, uint64(offender))
					skip := int(skipTo-i) + 1
					invalid += skip
					skipped += skip
					i = skipTo
				} else {
					invalid++
				}

			} else if boardState.M1() {
				m1++
			}
		}
	}

	total := math.CountBitsets(n) * iterator.NumberPermutations()
	fmt.Printf("M1 positions: %d/%d\n", m1, total)
	fmt.Printf("Skipped positions: %d/%d\n", skipped, total)
	fmt.Printf("Invalid positions: %d/%d\n", invalid, total)
}

//M1 positions: 462720/7624512
//Invalid positions: 2744448/7624512
