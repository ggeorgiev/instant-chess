package chess

import (
	"fmt"

	"github.com/ggeorgiev/instant-chess/src/bitboard"
	"github.com/ggeorgiev/instant-chess/src/board"
	"github.com/ggeorgiev/instant-chess/src/math"
	"github.com/ggeorgiev/instant-chess/src/move"
	"github.com/ggeorgiev/instant-chess/src/peace"
	"github.com/ggeorgiev/instant-chess/src/square"
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
	n := uint64(len(peaces))

	whiteRooks := 0
	blackRooks := 0
	whitePawns := 0
	blackPawns := 0
	for _, figure := range peaces {
		if figure == peace.WhiteRook {
			whiteRooks++
		}
		if figure == peace.BlackRook {
			blackRooks++
		}
		if figure == peace.WhitePawn {
			whitePawns++
		}
		if figure == peace.BlackPawn {
			blackPawns++
		}
	}

	rightsList := move.GenerateRightsList(whiteRooks, blackRooks, whitePawns, blackPawns)

	invalid := 0
	m1 := 0
	skipped := 0
	for _, rights := range rightsList {
		iterator := peace.NewPermutationIterator(peaces)
		for perm := iterator.Next(); perm != nil; perm = iterator.Next() {
			for i := uint64(0); i < math.CountBitsets(n); i++ {
				bitset := math.IndexToBitset(n, i)
				indexes := square.ConvertBitboardMaskIntoIndexes(bitboard.Mask(bitset))

				boardState := board.State{
					Rights: &rights,
				}
				for p, s := range indexes {
					boardState.Matrix[s] = perm[p]
				}

				invalidBoard, offender := boardState.Invalid()
				if invalidBoard {
					if offender != square.InvalidIndex {
						skipTo := math.NextValidIndex(n, i, bitset, uint64(offender))
						skip := int(skipTo - i)
						invalid += skip
						skipped += skip
						i = skipTo
					} else {
						invalid++
					}
					continue
				}

				if boardState.M1() {
					//fmt.Println(boardState.Sprint())
					m1++
				}
			}
		}
	}

	total := uint64(len(rightsList)) * math.CountBitsets(n) * peace.NewPermutationIterator(peaces).NumberPermutations()
	fmt.Printf("M1 positions: %d/%d\n", m1, total)
	fmt.Printf("Skipped positions: %d/%d\n", skipped, total)
	fmt.Printf("Invalid positions: %d/%d\n", invalid, total)
}
