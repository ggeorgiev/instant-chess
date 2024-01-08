package peacealignment

import (
	"github.com/ggeorgiev/instant-chess/src/bitboard"
	"github.com/ggeorgiev/instant-chess/src/peaceattacks"
	"github.com/ggeorgiev/instant-chess/src/square"
)

func overlap(mask1 bitboard.Mask, mask2 bitboard.Mask, attackedSquareMask bitboard.Mask) bitboard.Mask {
	common := mask1 & (mask2 | attackedSquareMask)
	if common == bitboard.Empty {
		return bitboard.Empty
	}
	return mask1 & ^common
}

func blockRelationBitboardMasksInternalHelper() bitboard.MasksList {
	var masksList bitboard.MasksList

	for i := square.ZeroIndex; i <= square.LastIndex; i++ {
		var masks bitboard.Masks
		for j := square.ZeroIndex; j <= square.LastIndex; j++ {
			mask := bitboard.Empty
			if i != j {
				mask |= overlap(peaceattacks.RiseBitboardMasks[i], peaceattacks.RiseBitboardMasks[j], square.IndexMask[j]) |
					overlap(peaceattacks.FallBitboardMasks[i], peaceattacks.FallBitboardMasks[j], square.IndexMask[j]) |
					overlap(peaceattacks.LeftBitboardMasks[i], peaceattacks.LeftBitboardMasks[j], square.IndexMask[j]) |
					overlap(peaceattacks.RightBitboardMasks[i], peaceattacks.RightBitboardMasks[j], square.IndexMask[j]) |
					overlap(peaceattacks.RiseLeftBitboardMasks[i], peaceattacks.RiseLeftBitboardMasks[j], square.IndexMask[j]) |
					overlap(peaceattacks.RiseRightBitboardMasks[i], peaceattacks.RiseRightBitboardMasks[j], square.IndexMask[j]) |
					overlap(peaceattacks.FallLeftBitboardMasks[i], peaceattacks.FallLeftBitboardMasks[j], square.IndexMask[j]) |
					overlap(peaceattacks.FallRightBitboardMasks[i], peaceattacks.FallRightBitboardMasks[j], square.IndexMask[j])
			}
			masks = append(masks, mask)
		}
		masksList = append(masksList, masks)
	}

	return masksList
}
