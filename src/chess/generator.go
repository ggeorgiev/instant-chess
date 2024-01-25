package chess

import (
	"fmt"

	"github.com/ggeorgiev/instant-chess/src/board/batch"
	"github.com/ggeorgiev/instant-chess/src/peace"
	"github.com/ggeorgiev/instant-chess/src/square"
)

// Generate generates all possible positions for the given peaces
func Generate(peacesString string) error {
	peaces := peace.MustParseFigures(peacesString)
	iterator := peace.NewPermutationIterator(peaces)
	rightsList := peaces.MoveRights()

	var batches batch.Batches
	for perm := iterator.Next(); perm != nil; perm = iterator.Next() {
		for _, rights := range rightsList {
			b, err := batch.Create(perm, rights)
			if err != nil {
				return err
			}
			batches = append(batches, b)
		}
	}

	stats := &batch.Stats{}
	for _, batch := range batches {
		for i := uint64(0); i < batch.CountBitsets(); i++ {
			boardState, bitset, offender := batch.GenerateState(i)
			if boardState == nil {
				i = batch.SkipOffender(i, bitset, offender, stats)
				continue
			}

			invalidBoard, offender := boardState.Invalid()
			if invalidBoard {
				if offender != square.InvalidIndex {
					i = batch.SkipOffender(i, bitset, offender, stats)
				} else {
					stats.Invalid++
				}
				continue
			}

			stats.States++

			if boardState.M1() {
				stats.M1++
			}
		}
	}

	fmt.Printf("Rights positions: %d\n", len(rightsList))
	fmt.Printf("Batches positions: %d\n", len(batches))
	fmt.Printf("M1 positions: %d(%.02f%%)\n", stats.M1, float64(stats.M1)/float64(stats.States))
	fmt.Printf("Skipped positions: %d(%.02f%%)\n", stats.Skipped, float64(stats.Skipped)/float64(stats.States))
	fmt.Printf("Invalid positions: %d(%.02f%%)\n", stats.Invalid, float64(stats.Invalid)/float64(stats.States))
	fmt.Printf("States positions: %d\n", stats.States)

	return nil
}

// M1 positions: 33616/15249024
// Skipped positions: 2094283/15249024
// Invalid positions: 2094283/15249024
