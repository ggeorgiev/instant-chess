package peaceattacks

import (
	"github.com/ggeorgiev/instant-chess/src/bitboard"
	"github.com/ggeorgiev/instant-chess/src/peacemoves"
	"github.com/ggeorgiev/instant-chess/src/square"
)

func knightBitboardMasksInternalHelper() bitboard.Masks {
	var masks bitboard.Masks

	indexesList := peacemoves.KnightMoves

	for _, indexes := range indexesList {
		masks = append(masks, square.ConvertIndexesIntoBitboardMask(indexes))
	}

	return masks
}
