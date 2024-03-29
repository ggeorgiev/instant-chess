package attack

import (
	"github.com/ggeorgiev/instant-chess/src/bitboard"
	"github.com/ggeorgiev/instant-chess/src/square"
)

func kingBitboardMasksInternalHelper() bitboard.Masks {
	var masks bitboard.Masks

	indexesList := FromKing

	for _, indexes := range indexesList {
		masks = append(masks, square.ConvertIndexesIntoBitboardMask(indexes))
	}

	return masks
}
