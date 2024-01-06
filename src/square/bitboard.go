package square

import (
	"fmt"
	"strings"

	"github.com/ggeorgiev/instant-chess/src/bitboard"
)

var IndexMask = [64]bitboard.Mask{
	0b0000000000000000000000000000000000000000000000000000000000000001,
	0b0000000000000000000000000000000000000000000000000000000000000010,
	0b0000000000000000000000000000000000000000000000000000000000000100,
	0b0000000000000000000000000000000000000000000000000000000000001000,
	0b0000000000000000000000000000000000000000000000000000000000010000,
	0b0000000000000000000000000000000000000000000000000000000000100000,
	0b0000000000000000000000000000000000000000000000000000000001000000,
	0b0000000000000000000000000000000000000000000000000000000010000000,
	0b0000000000000000000000000000000000000000000000000000000100000000,
	0b0000000000000000000000000000000000000000000000000000001000000000,
	0b0000000000000000000000000000000000000000000000000000010000000000,
	0b0000000000000000000000000000000000000000000000000000100000000000,
	0b0000000000000000000000000000000000000000000000000001000000000000,
	0b0000000000000000000000000000000000000000000000000010000000000000,
	0b0000000000000000000000000000000000000000000000000100000000000000,
	0b0000000000000000000000000000000000000000000000001000000000000000,
	0b0000000000000000000000000000000000000000000000010000000000000000,
	0b0000000000000000000000000000000000000000000000100000000000000000,
	0b0000000000000000000000000000000000000000000001000000000000000000,
	0b0000000000000000000000000000000000000000000010000000000000000000,
	0b0000000000000000000000000000000000000000000100000000000000000000,
	0b0000000000000000000000000000000000000000001000000000000000000000,
	0b0000000000000000000000000000000000000000010000000000000000000000,
	0b0000000000000000000000000000000000000000100000000000000000000000,
	0b0000000000000000000000000000000000000001000000000000000000000000,
	0b0000000000000000000000000000000000000010000000000000000000000000,
	0b0000000000000000000000000000000000000100000000000000000000000000,
	0b0000000000000000000000000000000000001000000000000000000000000000,
	0b0000000000000000000000000000000000010000000000000000000000000000,
	0b0000000000000000000000000000000000100000000000000000000000000000,
	0b0000000000000000000000000000000001000000000000000000000000000000,
	0b0000000000000000000000000000000010000000000000000000000000000000,
	0b0000000000000000000000000000000100000000000000000000000000000000,
	0b0000000000000000000000000000001000000000000000000000000000000000,
	0b0000000000000000000000000000010000000000000000000000000000000000,
	0b0000000000000000000000000000100000000000000000000000000000000000,
	0b0000000000000000000000000001000000000000000000000000000000000000,
	0b0000000000000000000000000010000000000000000000000000000000000000,
	0b0000000000000000000000000100000000000000000000000000000000000000,
	0b0000000000000000000000001000000000000000000000000000000000000000,
	0b0000000000000000000000010000000000000000000000000000000000000000,
	0b0000000000000000000000100000000000000000000000000000000000000000,
	0b0000000000000000000001000000000000000000000000000000000000000000,
	0b0000000000000000000010000000000000000000000000000000000000000000,
	0b0000000000000000000100000000000000000000000000000000000000000000,
	0b0000000000000000001000000000000000000000000000000000000000000000,
	0b0000000000000000010000000000000000000000000000000000000000000000,
	0b0000000000000000100000000000000000000000000000000000000000000000,
	0b0000000000000001000000000000000000000000000000000000000000000000,
	0b0000000000000010000000000000000000000000000000000000000000000000,
	0b0000000000000100000000000000000000000000000000000000000000000000,
	0b0000000000001000000000000000000000000000000000000000000000000000,
	0b0000000000010000000000000000000000000000000000000000000000000000,
	0b0000000000100000000000000000000000000000000000000000000000000000,
	0b0000000001000000000000000000000000000000000000000000000000000000,
	0b0000000010000000000000000000000000000000000000000000000000000000,
	0b0000000100000000000000000000000000000000000000000000000000000000,
	0b0000001000000000000000000000000000000000000000000000000000000000,
	0b0000010000000000000000000000000000000000000000000000000000000000,
	0b0000100000000000000000000000000000000000000000000000000000000000,
	0b0001000000000000000000000000000000000000000000000000000000000000,
	0b0010000000000000000000000000000000000000000000000000000000000000,
	0b0100000000000000000000000000000000000000000000000000000000000000,
	0b1000000000000000000000000000000000000000000000000000000000000000,
}

func ConvertIndexesIntoBitboardMask(indexes Indexes) bitboard.Mask {
	mask := bitboard.Empty
	for _, index := range indexes {
		mask |= IndexMask[index]
	}
	return mask
}

func SprintMask(mask bitboard.Mask) string {
	letters := "····a···b···c···d···e···f···g···h····\n"
	separator := "··+---+---+---+---+---+---+---+---+··\n"

	var result strings.Builder
	result.Grow(len(letters)*2 + len(separator)*9 + 8*(9*4+1))

	result.WriteString(letters)
	result.WriteString(separator)

	for y := Rank(7); y >= 0; y-- {
		result.WriteString(fmt.Sprintf("%d·|", y+1))
		for x := File(0); x < 8; x++ {
			symbol := " "
			if IndexMask[NewIndex(x, y)]&mask != 0 {
				symbol = "●"
			}
			result.WriteString(fmt.Sprintf(" %s |", symbol))
		}
		result.WriteString(fmt.Sprintf("·%d\n", y+1))
		result.WriteString(separator)
	}

	result.WriteString(letters)
	return result.String()
}
