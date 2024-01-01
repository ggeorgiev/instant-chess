package peaceattacks

import (
	"github.com/ggeorgiev/instant-chess/src/bitboard"
	"github.com/ggeorgiev/instant-chess/src/square"
)

func whitePawnSquareIndexesInternalHelper() []square.Indexes {
	var squaresList []square.Indexes

	for s := square.ZeroIndex; s <= square.LastIndex; s++ {
		var squares square.Indexes
		x := s.X()
		y := s.Y()

		if y != 7 && y > 0 {
			if x > 0 {
				squares = append(squares, square.NewIndex(x-1, y+1))
			}
			if x < 7 {
				squares = append(squares, square.NewIndex(x+1, y+1))
			}
		}
		squaresList = append(squaresList, squares)
	}
	return squaresList
}

func blackPawnSquareIndexesInternalHelper() []square.Indexes {
	var squaresList []square.Indexes

	for s := square.ZeroIndex; s <= square.LastIndex; s++ {
		var squares square.Indexes
		x := s.X()
		y := s.Y()

		if y != 0 && y < 7 {
			if x > 0 {
				squares = append(squares, square.NewIndex(x-1, y-1))
			}
			if x < 7 {
				squares = append(squares, square.NewIndex(x+1, y-1))
			}
		}
		squaresList = append(squaresList, squares)
	}
	return squaresList
}

func whitePawnBitboardMasksInternalHelper() bitboard.Masks {
	var masks bitboard.Masks

	indexesList := FromWhitePawn

	for _, indexes := range indexesList {
		masks = append(masks, square.ConvertIndexesIntoBitboardMask(indexes))
	}

	return masks
}

func blackPawnBitboardMasksInternalHelper() bitboard.Masks {
	var masks bitboard.Masks

	indexesList := FromBlackPawn

	for _, indexes := range indexesList {
		masks = append(masks, square.ConvertIndexesIntoBitboardMask(indexes))
	}

	return masks
}
