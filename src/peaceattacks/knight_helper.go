package peaceattacks

import (
	"github.com/ggeorgiev/instant-chess/src/bitboard"
	"github.com/ggeorgiev/instant-chess/src/square"
)

func knightBitboardMasksInternalHelper() bitboard.Masks {
	var masks bitboard.Masks

	indexesList := FromKnight

	for _, indexes := range indexesList {
		masks = append(masks, square.ConvertIndexesIntoBitboardMask(indexes))
	}

	return masks
}
