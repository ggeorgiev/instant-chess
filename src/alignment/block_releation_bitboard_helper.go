package alignment

import (
	"github.com/ggeorgiev/instant-chess/src/attack"
	"github.com/ggeorgiev/instant-chess/src/bitboard"
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
				mask |= overlap(attack.RiseBitboardMasks[i], attack.RiseBitboardMasks[j], square.IndexMask[j]) |
					overlap(attack.FallBitboardMasks[i], attack.FallBitboardMasks[j], square.IndexMask[j]) |
					overlap(attack.LeftBitboardMasks[i], attack.LeftBitboardMasks[j], square.IndexMask[j]) |
					overlap(attack.RightBitboardMasks[i], attack.RightBitboardMasks[j], square.IndexMask[j]) |
					overlap(attack.RiseLeftBitboardMasks[i], attack.RiseLeftBitboardMasks[j], square.IndexMask[j]) |
					overlap(attack.RiseRightBitboardMasks[i], attack.RiseRightBitboardMasks[j], square.IndexMask[j]) |
					overlap(attack.FallLeftBitboardMasks[i], attack.FallLeftBitboardMasks[j], square.IndexMask[j]) |
					overlap(attack.FallRightBitboardMasks[i], attack.FallRightBitboardMasks[j], square.IndexMask[j])
			}
			masks = append(masks, mask)
		}
		masksList = append(masksList, masks)
	}

	return masksList
}
